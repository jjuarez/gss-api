package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jjuarez/gss-api/internal/config"
	"github.com/jjuarez/gss-api/internal/utils"

	"github.com/gin-gonic/gin"
)

const (
	gssAPITargetEnvKey       = "GSSAPI_TARGET"
	gssAPITargetDefaultValue = "Defalt Target"
	configurationErrorCode   = 1
)

var (
	// Version the version of the GSS API release
	Version string
	// GitCommit information about the CVS
	GitCommit string
)

func main() {
	configuredTarget := utils.GetEnvWithDefault(gssAPITargetEnvKey, gssAPITargetDefaultValue)
	serverAddress, err := config.New()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(configurationErrorCode)
	}

	log.Printf("GSS API Version: %s", Version)
	log.Printf("  server address: %s", serverAddress.String())
	log.Printf("  configured target: %s", configuredTarget)

	router := gin.Default()
	server := &http.Server{
		Addr:    serverAddress.String(),
		Handler: router,
	}

	// Routes
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/target", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": configuredTarget,
		})
	})

	router.GET("/slow", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"message": "This was superslow",
		})
	})

	// Main
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shundown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout 5 seconds.")
	}
	log.Println("Server exiting")
}
