package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"github.com/joho/godotenv"
	"regexp"
	"examword_crawler/models"
	"net/http"
	"io/ioutil"
	"strings"
)


func main() {

	godotenv.Load()
	models.InitDB()
	c := colly.NewCollector(colly.CacheDir("./.examword_cache"))
	detailCollector := c.Clone()

	word_def_url := "https://www.examword.com/netservice/servicexmloutgate.aspx?at=lookupbyword&word=%s"

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

	c.OnHTML("span[class='listWordWord2']", func(element *colly.HTMLElement) {
		word := element.Text[0:len(element.Text)-1]
		def_url := fmt.Sprintf(word_def_url, word)
		fmt.Println(word, def_url)
		resp, _ := http.Get(def_url)
		body_bytes, _ := ioutil.ReadAll(resp.Body)
		meaning := string(body_bytes)
		i := strings.Index(meaning, ".")
		word_type := meaning[:i]
		fmt.Println(word_type, meaning)
		w := new(models.Word)
		w.NewWord(word, word_type, meaning, "", element.Attr("id"))
		fmt.Println(w)
		models.SaveWord(w)
		//fmt.Println(element.Attr("id"))
	})

	detailCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting detail", r.URL)
	})

	detailCollector.OnHTML("div[class]", func(element *colly.HTMLElement) {
		fmt.Println(element.Attr("href"))
	})

	c.Visit("https://www.examword.com/ielts-list/4000-general-word-1")

	fmt.Println(add(1,2))

	defer models.CloseDB()
}

func add(x,y int) int {
	return x + y
}