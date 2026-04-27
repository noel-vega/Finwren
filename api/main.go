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
	"github.com/noel-vega/finances/api/internal/config"
	"github.com/noel-vega/finances/api/internal/logging"
	middleware "github.com/noel-vega/finances/api/internal/middleware"
	"github.com/noel-vega/finances/api/internal/user"
)

func main() {
	slog.SetDefault(logging.New(os.Stdout, slog.LevelInfo))
	cfg, errs := config.New()
	if len(errs) > 0 {
		for _, err := range errs {
			slog.Error("invalid env variable", "error", err.Error())
		}
		os.Exit(1)
	}

	slog.SetDefault(logging.New(os.Stdout, cfg.LogLevel))

	db, err := sqlx.Connect("pgx", cfg.DatabaseConnectionString)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	userService := user.NewService(user.NewRepository(db))
	authService := auth.NewService(userService, cfg.Domain, cfg.JWTSecret)
	authHandler := auth.NewHandler(authService)

	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	authRoute := r.Group("/auth")
	authRoute.POST("/sign-in", authHandler.SignIn)
	authRoute.POST("/sign-up", authHandler.SignUp)

	// Start server on port 8080 (default)
	if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
