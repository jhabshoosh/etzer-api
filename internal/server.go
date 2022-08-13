package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/jhabshoosh/etzer-api/internal/config"
	"github.com/jhabshoosh/etzer-api/internal/db"
	"github.com/jhabshoosh/etzer-api/internal/graph"
	"github.com/jhabshoosh/etzer-api/internal/graph/generated"
)


type Server struct {
	Router  *mux.Router
	Neo4JDB db.Neo4JDB
	Env		config.Env
}

func Init() *Server {

	env := config.GetEnv()

	neo4Conn, err := db.NewNeo4jConnection(env.Neo4JProtocol, env.Neo4JHost, env.Neo4JPort, env.Neo4JUser, env.Neo4JPassword)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	neo4jDB := &db.Neo4JDB{
		Connection: neo4Conn,
	}

	return &Server{
		Neo4JDB: *neo4jDB,
		Env: env,
	}

}

// Run executes app
func (a *Server) Run() {
	addrStr := fmt.Sprintf("0.0.0.0:%s", &a.Env.Port)
	log.Println("API Listening", "api_url", addrStr)
	log.Fatal("Service failure", http.ListenAndServe(addrStr, a.Router))
}

// InitRoutes initializing all the routes
func (a *Server) InitRoutes() {
	a.Router = mux.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Neo4JDB: a.Neo4JDB}}))
	a.Router.Handle("/playground", playground.Handler("GoNeo4jGql GraphQL playground", "/movies"))
	a.Router.Handle("/movies", srv)
}