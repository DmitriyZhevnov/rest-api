package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/DmitriyZhevnov/rest-api/internal/config"
	"github.com/DmitriyZhevnov/rest-api/internal/handler"
	"github.com/DmitriyZhevnov/rest-api/internal/repository"
	"github.com/DmitriyZhevnov/rest-api/internal/service"

	"github.com/DmitriyZhevnov/rest-api/pkg/client/mongodb"
	"github.com/DmitriyZhevnov/rest-api/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/rest-api/pkg/hash"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	maxAttemptsForConnectPostgres = 5
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	logger.Info("register user handler")

	cfgMongo := cfg.Storage.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Database)
	if err != nil {
		panic(err)
	}

	postgresClient, err := postgresql.NewClient(context.Background(), maxAttemptsForConnectPostgres, cfg.Storage.Postgresql)
	if err != nil {
		panic(err)
	}

	storage := repository.NewRepository(mongoDBClient, cfgMongo.Collection, postgresClient, logger)

	service := service.NewService(hasher, storage, logger)

	handler := handler.NewHandler(service, logger)
	handler.Register(router)

	startServer(router, cfg)
}

func startServer(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("server is listening port :%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
