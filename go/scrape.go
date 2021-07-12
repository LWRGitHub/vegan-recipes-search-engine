package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	// "fmt"
	// "context"
	// "flag"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	// "strings"
	// "text/template"

	"github.com/gocolly/colly"
)

type VegFood struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Href  string `json:"href"`
}

func makeDataJS() {
	filename := "file.txt"
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	contents := string(fileContents)

	fmt.Println(contents)

	f, err := os.Create(strings.SplitN("datas", ".", 2)[0] + ".js")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(f, contents)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func main() {
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("recipeforvegans.com"),
	)

	var vegFoodsSlice []VegFood

	// vegFoodSlice := make([]VegFood, 5)

	c.OnHTML("a.dj-thumb-link", func(e *colly.HTMLElement) {

		title := e.Attr("title")
		href := e.Attr("href")
		img := e.ChildAttr("img", "src")

		vegFood := VegFood{
			Title: title,
			Img:   img,
			Href:  href,
		}
		vegFoodsSlice = append(vegFoodsSlice, vegFood)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(vegFoodsSlice)
		vegFoodJson, _ := json.Marshal(vegFoodsSlice)

		vegFoodsStr := string(vegFoodJson)
		fmt.Println(vegFoodsStr)

		if err := os.WriteFile("file.json", []byte(vegFoodsStr), 0666); err != nil {
			log.Fatal(err)
		}
	})

	c.Visit("https://recipeforvegans.com/")

}
