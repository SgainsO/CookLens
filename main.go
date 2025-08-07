package main

import ("fmt"
		"github.com/gocolly/colly"
		"github.com/jbrukh/bayesian")



func main() {
    fmt.Println("Input a Link you want to find the recipe of!")
    c:=colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    c.OnError(func(_ *colly.Response, err error) {
        fmt.Println("Something went wrong:", err)
    })

    c.OnHTML("li", func(e *colly.HTMLElement) {
        fmt.Println(e.Text)
    })
    bmodel, _ := bayesian.NewClassifierFromFile("model/model.mo")

    search("https://www.allrecipes.com/recipe/218057/chicken-enchilada-slow-cooker-soup/", c, bmodel)
}


func search(link string, col *colly.Collector, bModel *bayesian.Classifier) {

	col.Visit(link)

}
