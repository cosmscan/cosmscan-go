package api

import (
	"cosmscan-go/api/handlers"
	"cosmscan-go/config"
	"cosmscan-go/db"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	log *zap.SugaredLogger
	cfg *config.ServerConfig
	db  db.DB
}

func NewServer(cfg *config.ServerConfig) *Server {
	return &Server{
		log: zap.S().Named("api"),
		cfg: cfg,
		db:  MustInitDB(cfg),
	}
}

// WithHandlers attach all the handlers to the gin.Engine
func (s *Server) WithHandlers(e *gin.Engine) {
	e.GET("/block/:height", handlers.GetBlockByHeight)
}

func (s *Server) Serve() error {
	e := gin.New()
	e.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	e.Use(ginzap.RecoveryWithZap(zap.L(), true))
	e.Use(MiddlewareDatabaseContext(s.db))
	s.WithHandlers(e)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.cfg.Http.Host, s.cfg.Http.Port),
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.log.Info("starting api server")
	return server.ListenAndServe()
}
