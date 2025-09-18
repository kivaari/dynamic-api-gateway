package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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

	originalPath := c.Request.URL.Path

	parts := strings.Split(originalPath, "/")

	if len(parts) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route format"})
		return
	}

	newPath := "/" + strings.Join(parts[2:], "/")
	if newPath == "//" {
		newPath = "/"
	}

	c.Request.URL.Path = newPath

	proxy := httputil.NewSingleHostReverseProxy(u)
	c.Request.Host = u.Host

	proxy.ServeHTTP(c.Writer, c.Request)
}
