package main

import (
	"github.com/bottleneckco/showgrabber/src/backend/web"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	web.StartServer()
}
