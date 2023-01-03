package api

import (
	"cosmscan-go/config"

	"go.uber.org/zap"
)

type Server struct {
	log *zap.SugaredLogger
	cfg *config.ServerConfig
}

func NewServer(cfg *config.ServerConfig) *Server {
	return &Server{
		log: zap.S().Named("api"),
		cfg: cfg,
	}
}
