package middleware

import (
	"net/http"

	"github.com/didip/tollbooth/v5"
	"github.com/didip/tollbooth/v5/limiter"
	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	limiter *limiter.Limiter
}

func NewRateLimiter(rps int) *RateLimiter {
	limiter := tollbooth.NewLimiter(float64(rps), nil)
	limiter.SetMessage("Too many requests, please try again later.")
	limiter.SetStatusCode(http.StatusTooManyRequests)

	return &RateLimiter{
		limiter: limiter,
	}
}

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
