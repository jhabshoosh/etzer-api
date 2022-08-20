package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/jhabshoosh/etzer-api/internal/config"
	"github.com/jhabshoosh/etzer-api/internal/db"
	"github.com/jhabshoosh/etzer-api/internal/graph"
	"github.com/jhabshoosh/etzer-api/internal/graph/generated"
	"github.com/jhabshoosh/etzer-api/internal/services/person"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {

	ogm := db.InitNeo4JOGM()

	ps := &person.PersonService{
		Ogm: *ogm,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	srv := createServer(e, ps)

	initTransports(srv)
	initRoutes(e, srv)

	env := config.GetEnv()
	addrStr := fmt.Sprintf("0.0.0.0:%s", strconv.Itoa(env.Port))
	e.Logger.Fatal(e.Start(addrStr))
}

func createServer(e *echo.Echo, ps *person.PersonService) *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{PersonService: ps}}))
	initTransports(srv)
	initRoutes(e, srv)
	return srv
}

func initTransports(srv *handler.Server) {
	// Configure WebSocket with CORS
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
}

func initRoutes(e *echo.Echo, srv *handler.Server) {
	e.POST("/persons", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/persons").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Etzer API!")
	})
}
