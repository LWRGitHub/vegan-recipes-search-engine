package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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

type VegFoods struct {
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

	// var vegFoodsStr []byte

	// vegFoodsSlice := []map[string][]byte{}

	c.OnHTML("a.dj-thumb-link", func(e *colly.HTMLElement) {

		title := e.Attr("title")
		href := e.Attr("href")
		img := e.ChildAttr("img", "src")

		vegFoods := VegFoods{
			Title: title,
			Img:   img,
			Href:  href,
		}

		// fmt.Println(vegFoods)
		vegFoodJson, _ := json.Marshal(vegFoods)

		f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		vegFoodsStr := string(vegFoodJson) + ","

		if _, err = f.WriteString(vegFoodsStr); err != nil {
			panic(err)
		}
	})

	c.Visit("https://recipeforvegans.com/")

	makeDataJS()

}
