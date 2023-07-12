// main
package main

import (
	"html/template"
	"time"

	"encoding/gob"
	"fmt"
	"log"
	"net"
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

	// ctx = context.TODO()

	// Initialize DB:
	urlDAO = factoryURLDao()
	statsDAO = factoryStatsDao()

	gob.Register(&UserInMemory{})
}

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.Static("/assets", "./assets")

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	//router.LoadHTMLGlob("templates/*")

	// Parse the templates directory
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"currentYear": func() int {
			return time.Now().Year()
		},
	}).ParseGlob("templates/*")
	if err != nil {
		log.Fatal(err)
	}

	// Set the template engine
	router.SetHTMLTemplate(tmpl)

	// Initialize the routes
	initializeRoutes(envConfig)

	// Start serving the applications
	if err := router.Run(net.JoinHostPort("", serverPort)); err != nil {
		log.Fatal(err)
	}
}
