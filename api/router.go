package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zu1k/proxypool/app/cache"
	"github.com/zu1k/proxypool/provider"
)

var router *gin.Engine

func setupRouter() {
	router = gin.Default()

	router.StaticFile("/clash", "example/clash.html")
	router.StaticFile("/clash/config", "example/clash-config.yaml")
	router.GET("/clash/proxies", func(c *gin.Context) {
		text := cache.GetString("clashproxies")
		if text == "" {
			proxies := cache.GetProxies()
			clash := provider.Clash{Proxies: proxies}
			text = clash.Provide()
			cache.SetString("clashproxies", text)
		}
		c.String(200, text)
	})
}

func Run() {
	setupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
