package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/kivaari/dynamic-api-gateway/internal/config"
	"github.com/kivaari/dynamic-api-gateway/internal/gateway/middleware"
	"github.com/kivaari/dynamic-api-gateway/internal/logger"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	if cfg.Security.CORS.Enabled {
		r.Use(middleware.CORSMiddleware(cfg.Security.CORS.AllowedOrigins))
	}

	if cfg.Security.JWT.Enabled {
		r.Use(middleware.JWTAuthMiddleware(cfg.Security.JWT.Secret, cfg.Security.JWT.TokenHeader))
	}

	if cfg.Security.RateLimit.Enabled {
		limiter := middleware.NewRateLimiter(cfg.Security.RateLimit.RequestsPerSecond)
		r.Use(limiter.Limit())
	}

	// Пока статические маршруты — потом заменим на динамические из discovery
	setupStaticRoutes(r, cfg)

	return r
}

func setupStaticRoutes(r *gin.Engine, cfg *config.Config) {
	for _, svc := range cfg.Discovery.Static.Services {
		target := svc.Target
		logger.Log.Infof("Registering route: %s → %s", svc.Host, target)
		r.Any("/"+svc.Name+"/*path", proxyHandler(target))
	}
}

func proxyHandler(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ProxyRequest(c, targetURL)
	}
}
