// main
package main

import (
	"html/template"
	"log"
	"net"
	"time"

	"encoding/gob"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	var err error

	envConfig, err = readConfig("config.env", ".", map[string]interface{}{
		"port": "8080",
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	serverPort = envConfig.GetString("port")

	// Initialize DB:
	urlDAO = factoryURLDao()
	statsDAO = factoryStatsDao()

	gob.Register(&UserInMemory{})
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()
	router.SetFuncMap(template.FuncMap{
		"currentYear": func() int {
			return time.Now().Year()
		},
	})
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	initializeRoutes()

	if err := router.Run(net.JoinHostPort("", serverPort)); err != nil {
		log.Fatal(err)
	}
}
