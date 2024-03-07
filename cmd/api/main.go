package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "host=db port=5434 user=postgres password=123456789 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "API em Go"})
	})
	g.Run(":3000")

	// Read the SQL file
	sqlFile := "createtables.sql"
	file, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatal(err)
	}

	// Convert file bytes to string
	queries := string(file)

	fmt.Println("File reading done")

	// Execute the SQL queries
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables created successfully")
}
