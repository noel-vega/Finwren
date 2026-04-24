package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/noel-vega/finances/api/internal/auth"
	"github.com/noel-vega/finances/api/internal/user"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	r := gin.Default()
	config, err := NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to connect to database: %v\n", err))
		os.Exit(1)
	}

	db, err := sqlx.Connect("pgx", config.DatabaseConnectionString)
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to connect to database: %v\n", err))
		os.Exit(1)
	}

	userService := user.NewService(user.NewRepository(db))
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	authRoute := r.Group("/auth")
	authRoute.POST("/sign-in", authHandler.SignIn)
	authRoute.POST("/sign-up", authHandler.SignUp)

	// Start server on port 8080 (default)
	if err := r.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
