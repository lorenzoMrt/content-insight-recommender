package server

import (
	"time"

	kitlog "github.com/go-kit/log"
)

type ServerConfig struct {
	httpAddr        string
	shutdownTimeout time.Duration
}
type Server struct {
	config ServerConfig
	logger kitlog.Logger
}
