package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jjuarez/gss-api/internal/config"
	"github.com/jjuarez/gss-api/internal/utils"

	"github.com/gin-gonic/gin"
)

const (
	gssAPITargetEnvKey       = "GSSAPI_TARGET"
	gssAPITargetDefaultValue = "Defalt Target"
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
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()
	server := &http.Server{
		Addr:    serverAddress.String(),
		Handler: router,
	}

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

	server.ListenAndServe()
}
