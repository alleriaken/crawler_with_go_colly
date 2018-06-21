package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"os"
	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()

	SQL_HOST := os.Getenv("SQL_HOST")
	SQL_USERNAME := os.Getenv("SQL_USERNAME")
	SQL_PASSWORD := os.Getenv("SQL_PASSWORD")
	SQL_PORT := os.Getenv("SQL_PORT")
	SQL_DBNAME := os.Getenv("SQL_DBNAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%s/%s", SQL_USERNAME, SQL_PASSWORD, SQL_HOST, SQL_PORT, SQL_DBNAME))

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")

	defer db.Close()
}
