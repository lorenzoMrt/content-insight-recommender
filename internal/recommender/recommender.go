package recommender

import (
	"math"
	"strings"
)

// Tokeniza y normaliza el texto
func tokenize(text string) map[string]float64 {
	words := strings.Fields(strings.ToLower(text))
	wordFreq := make(map[string]float64)

	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

// Calcula la similitud de coseno entre dos textos
func CosineSimilarity(text1, text2 string) float64 {
	vec1 := tokenize(text1)
	vec2 := tokenize(text2)

	var dotProduct, magA, magB float64

	for word, freq1 := range vec1 {
		if freq2, exists := vec2[word]; exists {
			dotProduct += freq1 * freq2
		}
		magA += freq1 * freq1
	}

	for _, freq2 := range vec2 {
		magB += freq2 * freq2
	}

	if magA == 0 || magB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(magA) * math.Sqrt(magB))
}
