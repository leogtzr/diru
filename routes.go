package main

func initializeRoutes() {
	router.GET("/api/urls", viewURLs)
	router.GET("/api/stats", viewStats)
	router.GET("/u/:url", urlStats(), redirectShortURL)
	router.GET("/", showIndexPage)
	router.GET("/alv", showIndexPage)
	router.POST("/u/shorturl", shorturl)
	router.POST("/u/changelink", changeLink)
	router.GET("/stats", showStatsPage())
}
