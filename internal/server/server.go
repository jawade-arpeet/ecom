package server

import (
	"ecom/internal/client"
	"ecom/internal/config"
	"ecom/internal/handler"
	"ecom/internal/middleware"
	"ecom/internal/repo"
	"ecom/internal/router"
	"ecom/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(
	serverCfg *config.ServerConfig,
	hashingCfg *config.HashingConfig,
	jwtCfg *config.JWTConfig,
	client *client.Client,
) *Server {
	repo := repo.NewRepo(client)
	service := service.NewService(repo, hashingCfg, jwtCfg)
	handler := handler.NewHandler(service)
	middleware := middleware.NewMiddleware()
	router := router.NewRouter(handler, middleware)

	return &Server{
		config: serverCfg,
		router: router,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)

	zap.L().Info("starting server",
		zap.String("address", addr),
	)

	return s.router.Run(addr)
}
