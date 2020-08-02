package main

import (
	"log"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/util"
	"github.com/bottleneckco/showgrabber/src/backend/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.DB.LogMode(gin.IsDebugging())

	err := util.SeedLanguages()
	if err != nil {
		log.Println(err)
	}

	web.StartServer()
}
