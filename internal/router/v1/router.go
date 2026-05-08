package v1

import (
	"ecom/internal/handler"

	"github.com/gin-gonic/gin"
)

func MountV1Routes(router *gin.Engine, handler *handler.Handler) {
	v1 := router.Group("/api/v1")

	mountHealthRoutes(v1, handler.Health)
}
