package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	mu sync.RWMutex
)

var (
	router    *gin.Engine
	envConfig *viper.Viper
	urlDAO    *URLDao
	statsDAO  *StatsDAO

	serverPort string
)
