package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/meastblue/teckwatch/config"
	"github.com/meastblue/teckwatch/generated"
	"github.com/meastblue/teckwatch/resolver"
)

func Run() {
	c, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	sc := c.GetServerConf()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", sc.Host, sc.Port)
	log.Fatal(http.ListenAndServe(":"+sc.Port, nil))
}
