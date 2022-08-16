package server

import (
	"fmt"
	"log"
	"net/http"
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
	Router        *mux.Router
	PersonService person.PersonService
}

func Init() *Server {

	ogm := db.InitNeo4JOGM()

	personService := &person.PersonService{
		Ogm: *ogm,
	}

	return &Server{
		PersonService: *personService,
	}

}

// Run executes app
func (a *Server) Run() {
	env := config.GetEnv()
	addrStr := fmt.Sprintf("0.0.0.0:%s", strconv.Itoa(env.Port))
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
