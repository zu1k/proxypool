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

	router.StaticFile("/", "assets/index.html")
	router.StaticFile("/clash", "assets/clash.html")
	router.StaticFile("/surge", "assets/surge.html")
	router.StaticFile("/clash/config", "assets/clash-config.yaml")
	router.StaticFile("/surge/config", "assets/surge.conf")
	router.GET("/clash/proxies", func(c *gin.Context) {
		proxyTypes := c.DefaultQuery("type", "")
		text := ""
		if proxyTypes == "all" || proxyTypes == "" {
			text = cache.GetString("clashproxies")
			if text == "" {
				proxies := cache.GetProxies()
				clash := provider.Clash{Proxies: proxies}
				text = clash.Provide()
				cache.SetString("clashproxies", text)
			}
		} else {
			proxies := cache.GetProxies()
			clash := provider.Clash{Proxies: proxies, Types: proxyTypes}
			text = clash.Provide()
		}
		c.String(200, text)
	})
	router.GET("/surge/proxies", func(c *gin.Context) {
		text := cache.GetString("surgeproxies")
		if text == "" {
			proxies := cache.GetProxies()
			surge := provider.Surge{Proxies: proxies}
			text = surge.Provide()
			cache.SetString("surgeproxies", text)
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
