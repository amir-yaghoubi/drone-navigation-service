package main

import (
	"context"
	"drone_navigation_service/cmd/rest/config"
	"drone_navigation_service/pkg/dns"
	handler "drone_navigation_service/pkg/handler/http"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"

	_ "drone_navigation_service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Drone Navigation Service API
// @version 1.0

// @contact.name Amirhossein Yaghoubi
// @contact.email amir.yaghoubi.dev@gmail.com

// @BasePath /api/
func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Invalid configuration. %v", err)
	}

	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	apiRouter := router.Group("/api")

	svc := dns.NewDNS(cfg.SectorID)
	handler := handler.NewHandler(apiRouter, svc)

	handler.Setup()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	shutdownSignal := make(chan os.Signal, 1)

	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting DNS Service secID:%d on port: %d\n", cfg.SectorID, cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("cannot listen %v", err)
		}
	}()

	<-shutdownSignal

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Bye :-)")
}
