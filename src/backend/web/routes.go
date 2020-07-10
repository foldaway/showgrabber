package web

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{"http://localhost:1234"}
	r.Use(cors.New(corsCfg))

	r.GET("/graphql", graphQLPlayground)
	r.POST("/graphql", graphQL)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
