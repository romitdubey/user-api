package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	_ "github.com/lib/pq"

	"github.com/romitdubey1/user-api/config"
	"github.com/romitdubey1/user-api/internal/handler"
	"github.com/romitdubey1/user-api/internal/logger"
	"github.com/romitdubey1/user-api/internal/middleware"
	"github.com/romitdubey1/user-api/internal/repository"
	"github.com/romitdubey1/user-api/internal/routes"
	"github.com/romitdubey1/user-api/internal/service"
)

func main() {
	// load env
	_ = godotenv.Load()

	// load config
	cfg := config.Load()

	// zap logger
	logr := logger.New(cfg.Env)
	defer logr.Sync()

	// db connection
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		logr.Fatal("failed to connect database", zap.Error(err))
	}

	// fiber app
	app := fiber.New()

	// middleware
	app.Use(middleware.RequestLogger(logr))

	// dependencies
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService()
	h := handler.NewUserHandler(repo, svc, logr)

	// routes
	routes.Register(app, h)

	logr.Info("server started", zap.String("port", cfg.AppPort))
	logr.Fatal("server stopped", zap.Error(app.Listen(":"+cfg.AppPort)))
}
