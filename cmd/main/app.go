package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/DmitriyZhevnov/rest-api/internal/config"
	"github.com/DmitriyZhevnov/rest-api/internal/user"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
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
