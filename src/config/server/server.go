package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/kristiansantos/learning/src/config/initializer"
	"github.com/kristiansantos/learning/src/core/api/handlers"
	"github.com/kristiansantos/learning/src/core/api/routes"
	"github.com/kristiansantos/learning/src/shared/database/elasticsearch"
	"github.com/kristiansantos/learning/src/shared/database/mongodb"
	"github.com/kristiansantos/learning/src/shared/middleware"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

type server struct {
	Addr       string
	Port       int
	httpServer http.Server
}

func New(addr string, port int) *server {
	return &server{
		Addr: addr,
		Port: port,
	}
}

func (s *server) Run(app initializer.Application, log logger.ILoggerProvider) error {
	log.Info("server.main.Run", fmt.Sprintf("Server running on port :%d", s.Port))
	log.Info("server.main.Run", fmt.Sprintf("Environment: %s", app.Environment))
	log.Info("server.main.Run", fmt.Sprintf("Version: %s", app.Version))

	//Mongo db conection
	ctx := context.TODO()

	mongoConnection := mongodb.New(ctx)
	if mongoConnection.Error != nil {
		panic(fmt.Sprintf("error connecting to mongodb: %v", mongoConnection.Error))
	} else {
		log.Info("server.main.Run", fmt.Sprintf("Mongodb start connection"))
	}

	elasticsearchConnectionError := elasticsearch.Connection(app, log)
	if elasticsearchConnectionError != nil {
		panic(fmt.Sprintf("error connecting to elasticsearch: %v", mongoConnection.Error))
	} else {
		log.Info("server.main.Run", fmt.Sprintf("Elasticsearch start connection"))
	}

	handlerDependencies := handlers.Dependencies{Logger: log}
	router := routes.NewRoutes(handlers.NewHandler(handlerDependencies))
	router.Setup()

	s.httpServer = http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.Addr, s.Port),
		Handler:      middleware.Recovery(router.Client),
		ReadTimeout:  app.ReadTimeout * 2,
		WriteTimeout: app.WriteTimeout * 2,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				return
			} else {
				fmt.Println(err)
			}

			process, err := os.FindProcess(os.Getpid())
			if err != nil {
				fmt.Println("couldn't find process to exit")
				os.Exit(1)
			}

			if err := process.Signal(os.Interrupt); err != nil {
				fmt.Println("couldn't find process to exit")
				os.Exit(1)
			}

		}
	}()

	return nil
}
