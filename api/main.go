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
	"github.com/noel-vega/finances/api/internal/email"
	"github.com/noel-vega/finances/api/internal/logging"
	middleware "github.com/noel-vega/finances/api/internal/middleware"
	"github.com/noel-vega/finances/api/internal/user"
)

func main() {
	cfg, errs := config.New()
	if len(errs) > 0 {
		for _, err := range errs {
			slog.Error("invalid env variable", "error", err.Error())
		}
		os.Exit(1)
	}

	slog.SetDefault(logging.New(os.Stdout, cfg.Environment))

	db, err := sqlx.Connect("pgx", cfg.DatabaseConnectionString)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	userService := user.NewService(user.NewRepository(db))
	emailService := email.NewService(cfg.ResendKey)
	authService, err := auth.NewService(userService, emailService, cfg.Domain, cfg.JWTSecret)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	authHandler := auth.NewHandler(authService, cfg.Environment)

	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "unable to connect to db",
			})
			return
		}
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
