package main

import (
	"log"
	"os"
	"strconv"

	"sllpklls/admin-service/db"
	"sllpklls/admin-service/handler"
	"sllpklls/admin-service/repository/repo_impl"
	"sllpklls/admin-service/router"

	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 5432
	}

	sql := &db.Sql{
		Host:     host,
		Port:     port,
		UserName: user,
		Password: pass,
		DbName:   dbName,
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	userRepo := repo_impl.NewUserRepo(sql)
	userHandler := handler.UserHandler{
		UserRepo: userRepo,
	}
	infraComponentHandler := handler.InfraComponentHandler{
		InfraComponentRepo: repo_impl.NewInfraComponentRepo(sql),
		UserRepo:           userRepo,
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}

			return next(c)
		}
	})
	api := router.API{
		Echo:                  e,
		UserHandler:           userHandler,
		InfraComponentHandler: infraComponentHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}

// helper: lấy env hoặc fallback sang default
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
