package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showStatsPage(c *gin.Context) {
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

	urlsFull := urlsToFullStat(&urlStats)

	c.HTML(
		http.StatusOK,
		"stats.html",
		gin.H{
			"title": "URL Stats",
			"urls":  urlsFull,
		},
	)
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

	c.JSON(http.StatusOK, stats)
}
