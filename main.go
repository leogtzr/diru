// main
package main

import (
	"context"
	"html/template"
	"log"
	"net"
	"time"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/gin-contrib/sessions"
	redisSession "github.com/gin-contrib/sessions/redis"
)

var Ctx = context.Background()

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

	// REDIS_DSN=localhost:6379
	dsn := "redis:6379"
	redisClient = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	return rdb
}

const (
	// MaxIdleConnections ...
	MaxIdleConnections = 10
)

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

	dsn := "redis:6379"
	store, _ := redisSession.NewStore(MaxIdleConnections, "tcp", dsn, "", []byte(envConfig.GetString("SESSION_SECRET")))
	router.Use(sessions.Sessions("sid", store))

	/*
		if someError := redisClient.Set("bb", "b", -1).Err(); someError != nil {
			log.Fatal(someError)
		}*/

	initializeRoutes()

	if err := router.Run(net.JoinHostPort("", serverPort)); err != nil {
		log.Fatal(err)
	}
}
