package web

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bottleneckco/showgrabber/src/backend/graph"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/gin-gonic/gin"
)

var (
	main = handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	pg = playground.Handler("GraphQL", "/graphql")
)

func graphQL(c *gin.Context) {
	main.ServeHTTP(c.Writer, c.Request)
}

func graphQLPlayground(c *gin.Context) {
	pg.ServeHTTP(c.Writer, c.Request)
}
