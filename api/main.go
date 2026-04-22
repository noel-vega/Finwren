package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/noel-vega/finances/api/internal/auth"
)

func main() {
	r := gin.Default()
	config, err := NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config: %v\n", err)
		os.Exit(1)
	}

	db, err := sqlx.Connect("pgx", config.DatabaseConnectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	authHandler := auth.NewHandler()
	authRoute := r.Group("/auth")
	authRoute.POST("/sign-in", authHandler.SignIn)
	authRoute.POST("/sign-up", authHandler.SignUp)

	// Start server on port 8080 (default)
	if err := r.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
