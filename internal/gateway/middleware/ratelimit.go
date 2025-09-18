package middleware

import (
	"net/http"

	"github.com/didip/tollbooth/v5"
	"github.com/didip/tollbooth/v5/limiter"
	"github.com/gin-gonic/gin"
)

// RateLimiter — структура для управления лимитером
type RateLimiter struct {
	limiter *limiter.Limiter
}

// NewRateLimiter — создаёт новый лимитер с указанным количеством запросов в секунду
func NewRateLimiter(rps int) *RateLimiter {
	// Создаём лимитер: rps запросов в секунду
	limiter := tollbooth.NewLimiter(float64(rps), nil)
	limiter.SetMessage("Too many requests, please try again later.")
	limiter.SetStatusCode(http.StatusTooManyRequests)

	return &RateLimiter{
		limiter: limiter,
	}
}

// Limit — возвращает Gin middleware для ограничения запросов
func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(rl.limiter, c.Writer, c.Request)
		if httpError != nil {
			c.AbortWithStatusJSON(httpError.StatusCode, gin.H{
				"error": httpError.Message,
			})
			return
		}
		c.Next()
	}
}
