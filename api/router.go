package api

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/zu1k/proxypool/app"
	"github.com/zu1k/proxypool/provider"
)

var router *gin.Engine

func setupRouter() {
	router = gin.Default()

	router.StaticFile("/clash/config", "example/clash-config.yaml")
	router.GET("/clash/proxies", func(c *gin.Context) {
		proxies := app.GetProxies()
		clash := provider.Clash{Proxies: proxies}
		c.String(200, clash.Provide())
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
