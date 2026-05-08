package router

import (
	"ecom/internal/handler"
	"ecom/internal/middleware"
	v1 "ecom/internal/router/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	handler *handler.Handler,
	middleware *middleware.Middleware,
) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.AddRequestID())

	v1.MountV1Routes(router, handler)

	return router
}