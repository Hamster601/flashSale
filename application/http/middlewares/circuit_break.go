package middleware

import (
	"net/http"

	"github.com/Hamster601/flashSale/application/infrastructures/utils"
	"github.com/gin-gonic/gin"
)

func NewCircuitBreakMiddleware(cb *utils.CircuitBreaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok := cb.Allow(func() bool {
			c.Next()
			if c.Writer.Status() >= http.StatusInternalServerError {
				return false
			}
			return true
		})
		if !ok {
			c.AbortWithStatus(http.StatusServiceUnavailable)
		}
	}
}
