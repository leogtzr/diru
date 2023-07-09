package main

import (
	"fmt"
	"net/http"

	"net"

	"github.com/Showmax/go-fqdn"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func showStatsPage(config *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		urlStats, err := (*urlDAO).findAllByUser()
		if err != nil {
			c.HTML(
				http.StatusInternalServerError,
				"error5xx.html",
				gin.H{
					"title":             "Error",
					"error_description": err.Error(),
				},
			)

			return
		}

		fqdnHostName, err := fqdn.FqdnHostname()
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}

		domain := net.JoinHostPort(fqdnHostName, config.GetString("port"))

		urlsFull := urlsToFullStat(&urlStats)

		c.HTML(
			http.StatusOK,
			"stats.html",
			gin.H{
				"title":  "URL Stats",
				"domain": domain,
				"urls":   urlsFull,
			},
		)
	}
}

func urlStats() gin.HandlerFunc {
	return func(c *gin.Context) {

		shortURLParam := c.Param("url")
		if shortURLParam == "" {
			c.HTML(
				http.StatusInternalServerError,
				"error5xx.html",
				gin.H{
					"title":             "Error",
					"error_description": `error: missing url argument to redirect to`,
				},
			)

			return
		}

		headers := map[string][]string(c.Request.Header)
		// As of now, all the headers are being saved, we might want to consider to save only a few, such as:
		// Referrer, User-Agent, etc
		(*statsDAO).save(shortURLParam, &headers)
	}
}

func viewStats(c *gin.Context) {
	stats, err := (*statsDAO).findAll()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	fmt.Println(stats)

	c.JSON(http.StatusOK, stats)
}
