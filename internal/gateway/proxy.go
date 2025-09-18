package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/kivaari/dynamic-api-gateway/internal/logger"
)

func ProxyRequest(c *gin.Context, targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil {
		logger.Log.Errorf("Invalid target URL: %s", targetURL)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	c.Request.Host = u.Host
	proxy.ServeHTTP(c.Writer, c.Request)
}
