// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/jhabshoosh/etzer/db"
// 	"github.com/jhabshoosh/etzer/graph"
// 	"github.com/jhabshoosh/etzer/graph/generated"
// )

// const defaultPort = "8080"

// type HelloHandler struct{}
// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	res, err := db.HelloWorld("bolt://neo4j://localhost:7687", "neo4j", "test")
// 	if err != nil {
// 		fmt.Fprintf(w, err.Error())
// 	}
// 	fmt.Fprintf(w, res)
// }

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}


// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)
// 	hello := HelloHandler{}
// 	http.Handle("/foo", &hello)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }
