package main

import (
	"os"
	"strconv"

	"sllpklls/admin-service/db"
	"sllpklls/admin-service/handler"
	"sllpklls/admin-service/repository/repo_impl"
	"sllpklls/admin-service/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// lấy env, nếu không có thì dùng default
	host := getEnv("DB_HOST", "localhost")
	portStr := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "admin")
	pass := getEnv("DB_PASSWORD", "admin123")
	dbName := getEnv("DB_NAME", "mydb")

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
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
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
