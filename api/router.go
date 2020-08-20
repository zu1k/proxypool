package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zu1k/proxypool/config"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/zu1k/proxypool/internal/cache"
	"github.com/zu1k/proxypool/pkg/provider"
)

var router *gin.Engine
var domain = "proxy.tgbot.co"

func setupRouter() {
	domain = config.SourceConfig.Domain
	fmt.Println("Domain:", domain)

	router = gin.Default()
	router.LoadHTMLGlob("assets/html/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"domain":               domain,
			"all_proxies_count":    cache.AllProxiesCount,
			"useful_proxies_count": cache.UsefullProxiesCount,
		})
	})

	router.GET("/clash", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clash.html", gin.H{
			"domain": domain,
		})
	})

	router.GET("/surge", func(c *gin.Context) {
		c.HTML(http.StatusOK, "surge.html", gin.H{
			"domain": domain,
		})
	})

	router.GET("/clash/config", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clash-config.yaml", gin.H{
			"domain": domain,
		})
	})

	router.GET("/surge/config", func(c *gin.Context) {
		c.HTML(http.StatusOK, "surge.conf", gin.H{
			"domain": domain,
		})
	})

	router.GET("/clash/proxies", func(c *gin.Context) {
		proxyTypes := c.DefaultQuery("type", "")
		text := ""
		if proxyTypes == "" {
			text = cache.GetString("clashproxies")
			if text == "" {
				proxies := cache.GetProxies("proxies")
				clash := provider.Clash{Proxies: proxies}
				text = clash.Provide()
				cache.SetString("clashproxies", text)
			}
		} else if proxyTypes == "all" {
			proxies := cache.GetProxies("allproxies")
			clash := provider.Clash{Proxies: proxies, Types: proxyTypes}
			text = clash.Provide()
		} else {
			proxies := cache.GetProxies("proxies")
			clash := provider.Clash{Proxies: proxies, Types: proxyTypes}
			text = clash.Provide()
		}
		c.String(200, text)
	})
	router.GET("/surge/proxies", func(c *gin.Context) {
		text := cache.GetString("surgeproxies")
		if text == "" {
			proxies := cache.GetProxies("proxies")
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
