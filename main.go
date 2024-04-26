package main

import (
	"database/sql"
	"fmt"
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/KKGo-Software-engineering/assessment-tax/src/handler"
)

func main() {
	db := initDatabase()
	err := db.Ping()
	if err != nil {
		log.Fatal("Connect to database error")
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	e.POST("tax/calculations", handler.CalculateTax)

	e.Logger.Fatal(e.Start(":8080"))
}

func initDatabase() *sql.DB {
	fmt.Printf("DB_URL = %v", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
