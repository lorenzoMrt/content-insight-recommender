package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	models "github.com/lorenzoMrt/content-insight-recommender/internal/platform/storage/postgresql"
	"github.com/lorenzoMrt/content-insight-recommender/internal/recommender"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	urlDb := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", "admin", "development", "localhost", 5432, "recomendations")
	db, err = gorm.Open(postgres.Open(urlDb), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Content{})
	if err != nil {
		log.Fatal(err)
	}
}

func recommendContent(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	var contents []models.Content
	db.Find(&contents)

	var recommendations []models.Content
	for _, content := range contents {
		similarity := recommender.CosineSimilarity(user.Interests[0], content.Tags[0])
		if similarity > 0.5 { // Umbral de similitud
			recommendations = append(recommendations, content)
		}
	}

	c.JSON(200, recommendations)
}

func main() {
	r := gin.Default()
	r.GET("/recommend/:id", recommendContent)
	r.Run(":8080")
}
