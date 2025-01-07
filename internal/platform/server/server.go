package server

import (
	kitlog "github.com/go-kit/log"
	"time"
)

type Config struct {
	httpAddr        string
	shutdownTimeout time.Duration
}
type Server struct {
	config Config
	logger kitlog.Logger
}
