package main

import "github.com/spf13/viper"

func initializeRoutes(config *viper.Viper) {
	router.Use(setUserStatus())

	router.GET("/api/users", viewUsers)
	router.GET("/api/urls", viewURLs)
	router.GET("/api/stats", viewStats)

	router.GET("/u/:url", urlStats(), redirectShortURL)
	router.GET("/", ensureNotLoggedIn(), showIndexPage)
	router.POST("/u/shorturl", shorturl)
	router.POST("/u/changelink", changeLink)

	// stats URLs
	router.GET("/stats", showStatsPage(config))
}
