package api

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/zu1k/proxypool/config"
	binhtml "github.com/zu1k/proxypool/internal/bindata/html"
	C "github.com/zu1k/proxypool/internal/cache"
	"github.com/zu1k/proxypool/pkg/provider"
)

const version = "v0.3.10"

var router *gin.Engine

func setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	store := persistence.NewInMemoryStore(time.Minute)
	router.Use(gin.Recovery(), cache.SiteCache(store, time.Minute))
	temp, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(temp)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "assets/html/index.html", gin.H{
			"domain":               config.Config.Domain,
			"getters_count":        C.GettersCount,
			"all_proxies_count":    C.AllProxiesCount,
			"ss_proxies_count":     C.SSProxiesCount,
			"ssr_proxies_count":    C.SSRProxiesCount,
			"vmess_proxies_count":  C.VmessProxiesCount,
			"trojan_proxies_count": C.TrojanProxiesCount,
			"useful_proxies_count": C.UsefullProxiesCount,
			"last_crawl_time":      C.LastCrawlTime,
			"version":              version,
		})
	})

	router.GET("/clash", func(c *gin.Context) {
		c.HTML(http.StatusOK, "assets/html/clash.html", gin.H{
			"domain": config.Config.Domain,
		})
	})

	router.GET("/surge", func(c *gin.Context) {
		c.HTML(http.StatusOK, "assets/html/surge.html", gin.H{
			"domain": config.Config.Domain,
		})
	})

	router.GET("/clash/config", func(c *gin.Context) {
		c.HTML(http.StatusOK, "assets/html/clash-config.yaml", gin.H{
			"domain": config.Config.Domain,
		})
	})

	router.GET("/surge/config", func(c *gin.Context) {
		c.HTML(http.StatusOK, "assets/html/surge.conf", gin.H{
			"domain": config.Config.Domain,
		})
	})

	router.GET("/clash/proxies", func(c *gin.Context) {
		proxyTypes := c.DefaultQuery("type", "")
		proxyCountry := c.DefaultQuery("c", "")
		proxyNotCountry := c.DefaultQuery("nc", "")
		text := ""
		if proxyTypes == "" && proxyCountry == "" && proxyNotCountry == "" {
			text = C.GetString("clashproxies")
			if text == "" {
				proxies := C.GetProxies("proxies")
				clash := provider.Clash{
					provider.Base{
						Proxies: &proxies,
					},
				}
				text = clash.Provide()
				C.SetString("clashproxies", text)
			}
		} else if proxyTypes == "all" {
			proxies := C.GetProxies("allproxies")
			clash := provider.Clash{
				provider.Base{
					Proxies:    &proxies,
					Types:      proxyTypes,
					Country:    proxyCountry,
					NotCountry: proxyNotCountry,
				},
			}
			text = clash.Provide()
		} else {
			proxies := C.GetProxies("proxies")
			clash := provider.Clash{
				provider.Base{
					Proxies:    &proxies,
					Types:      proxyTypes,
					Country:    proxyCountry,
					NotCountry: proxyNotCountry,
				},
			}
			text = clash.Provide()
		}
		c.String(200, text)
	})

	router.GET("/surge/proxies", func(c *gin.Context) {
		proxyTypes := c.DefaultQuery("type", "")
		proxyCountry := c.DefaultQuery("c", "")
		proxyNotCountry := c.DefaultQuery("nc", "")
		text := ""
		if proxyTypes == "" && proxyCountry == "" && proxyNotCountry == "" {
			text = C.GetString("surgeproxies")
			if text == "" {
				proxies := C.GetProxies("proxies")
				surge := provider.Surge{
					provider.Base{
						Proxies: &proxies,
					},
				}
				text = surge.Provide()
				C.SetString("surgeproxies", text)
			}
		} else if proxyTypes == "all" {
			proxies := C.GetProxies("allproxies")
			surge := provider.Surge{
				provider.Base{
					Proxies:    &proxies,
					Types:      proxyTypes,
					Country:    proxyCountry,
					NotCountry: proxyNotCountry,
				},
			}
			text = surge.Provide()
		} else {
			proxies := C.GetProxies("proxies")
			surge := provider.Surge{
				provider.Base{
					Proxies:    &proxies,
					Types:      proxyTypes,
					Country:    proxyCountry,
					NotCountry: proxyNotCountry,
				},
			}
			text = surge.Provide()
		}
		c.String(200, text)
	})

	router.GET("/ss/sub", func(c *gin.Context) {
		proxies := C.GetProxies("proxies")
		ssSub := provider.SSSub{
			provider.Base{
				Proxies: &proxies,
				Types:   "ss",
			},
		}
		c.String(200, ssSub.Provide())
	})

	router.GET("/ssr/sub", func(c *gin.Context) {
		proxies := C.GetProxies("proxies")
		ssrSub := provider.SSRSub{
			provider.Base{
				Proxies: &proxies,
				Types:   "ssr",
			},
		}
		c.String(200, ssrSub.Provide())
	})

	router.GET("/vmess/sub", func(c *gin.Context) {
		proxies := C.GetProxies("proxies")
		vmessSub := provider.VmessSub{
			provider.Base{
				Proxies: &proxies,
				Types:   "vmess",
			},
		}
		c.String(200, vmessSub.Provide())
	})

	router.GET("/sip002/sub", func(c *gin.Context) {
		proxies := C.GetProxies("proxies")
		sip002Sub := provider.SIP002Sub{
			provider.Base{
				Proxies: &proxies,
				Types:   "ss",
			},
		}
		c.String(200, sip002Sub.Provide())
	})

	router.GET("/link/:id", func(c *gin.Context) {
		idx := c.Param("id")
		proxies := C.GetProxies("allproxies")
		id, err := strconv.Atoi(idx)
		if err != nil {
			c.String(500, err.Error())
		}
		if id >= proxies.Len() || id < 0 {
			c.String(500, "id out of range")
		}
		c.String(200, proxies[id].Link())
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

func loadTemplate() (t *template.Template, err error) {
	_ = binhtml.RestoreAssets("", "assets/html")
	t = template.New("")
	for _, fileName := range binhtml.AssetNames() {
		data := binhtml.MustAsset(fileName)
		t, err = t.New(fileName).Parse(string(data))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
