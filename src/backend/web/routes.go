package web

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(corsCfg))

	r.Any("/graphql", graphQL)
	r.Any("/playground", graphQLPlayground)

	if gin.IsDebugging() {
		// Start webpack-dev-server
		cmd := exec.Command("yarn", "webpack-dev-server")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		var err = cmd.Start()
		if err != nil {
			log.Panic(err)
		}

		defer func() {
			if cmd.Process != nil {
				cmd.Process.Kill()
			}
		}()

		// Proxy to webpack-dev-server
		var reverseProxy = httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   "localhost:3001",
		})

		r.Use(func(c *gin.Context) {
			reverseProxy.ServeHTTP(c.Writer, c.Request)
		})
	} else {
		// Serve built files
		r.Static("/dist", "dist")

		r.GET("/", func(c *gin.Context) {
			c.File("dist/index.html")
		})
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
