package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	// "fmt"
	// "context"
	// "flag"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	// "strings"
	// "text/template"
	// "reflect"

	"github.com/gocolly/colly"
)

type VegFood struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Href  string `json:"href"`
}

func GetData(vegFoodsSlice []VegFood) []byte {
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("recipeforvegans.com"),
	)

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

	vegFoodJson, _ := json.Marshal(vegFoodsSlice)

	return vegFoodJson

}

func main() {

	var vegFoodsSlice []VegFood

	GetData(vegFoodsSlice)

	fmt.Println(reflect.TypeOf(GetData(vegFoodsSlice)))
}
