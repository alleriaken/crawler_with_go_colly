package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"os"
	"github.com/joho/godotenv"
	"regexp"
	"examword_crawler/models"
)


func main() {

	godotenv.Load()
	models.InitDB()
	c := colly.NewCollector(colly.CacheDir("./.examword_cache"))
	detailCollector := c.Clone()
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		matched, err := regexp.MatchString(`ielts-list/4000-general-word-\d\?`, e.Attr("href"))
		if err != nil {
			fmt.Print(e.Attr("href"))
			os.Exit(400)
		}
		if matched {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div[class='glface']", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})

	detailCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting detail", r.URL)
	})

	detailCollector.OnHTML("div[class]", func(element *colly.HTMLElement) {
		fmt.Println("aaaaaaa")
		fmt.Println(element.Attr("href"))
	})

	c.Visit("https://www.examword.com/ielts-list/4000-general-word-1")

	fmt.Println(add(1,2))

	defer models.CloseDB()
}

func add(x,y int) int {
	return x + y
}