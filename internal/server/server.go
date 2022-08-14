package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/jhabshoosh/etzer-api/internal/config"
	"github.com/jhabshoosh/etzer-api/internal/db"
	"github.com/jhabshoosh/etzer-api/internal/graph"
	"github.com/jhabshoosh/etzer-api/internal/graph/generated"
	"github.com/jhabshoosh/etzer-api/internal/services/person"
)
type Server struct {
	Router 			*mux.Router
	PersonService 	person.PersonService
	Env				config.Env
}

func Init() *Server {

	env := config.GetEnv()

	neo4Conn, err := db.NewNeo4jConnection("bolt", env.DBHost, env.DBPort, env.DBUser, env.DBPassword)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	neo4jDB := &db.Neo4JDB{
		Connection: neo4Conn,
	}
	
	personService := &person.PersonService{
		Neo4jClient: *neo4jDB,
	}

	return &Server{
		PersonService: *personService,
		Env: env,
	}

}

// Run executes app
func (a *Server) Run() {
	addrStr := fmt.Sprintf("0.0.0.0:%s", strconv.Itoa(a.Env.Port))
	log.Println("API Listening", "api_url", addrStr)
	log.Fatal("Service failure", http.ListenAndServe(addrStr, a.Router))
}

// InitRoutes initializing all the routes
func (s *Server) InitRoutes() {
	s.Router = mux.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{PersonService: s.PersonService}}))
	s.Router.Handle("/playground", playground.Handler("GoNeo4jGql GraphQL playground", "/persons"))
	s.Router.Handle("/persons", srv)
}