package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (mw *Middleware) AddRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("request_id", uuid.New())
		ctx.Next()
	}
}
