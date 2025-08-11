package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jbrukh/bayesian"
)

type Memory struct {
	Items       [3]string
	Item_Type   [3]int8 //0  is not Ing, 1 is Ing
	Amt         int8
	Amt_Correct int8
}

func (m Memory) ReturnLeftovers(s []string) []string {
	for index, value := range m.Items {
		if m.Item_Type[index] == 1 {
			s = AddToIngSlice(value)
		} else {
			break
		}
	}
	return s
}
func (m Memory) ClearMemory() {
	m.Items = [3]string{}
	m.Item_Type = [3]int8{}
	m.Amt_Correct = 0
}

var memory Memory = Memory{Items: [3]string{},
	Item_Type: [3]int8{}, Amt: 0, Amt_Correct: 0}

var Ings []string = []string{}

func main() {
	fmt.Println("Input a Link you want to find the recipe of!")
	c := colly.NewCollector()
	leftovers := [3]string{"", "", ""}
	leftoverSet := false

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("li", func(e *colly.HTMLElement) {
		trimmedText := strings.TrimSpace(e.Text)
		if trimmedText != "" {
			fmt.Println(trimmedText)
			memory.AddToMemory(trimmedText)
			memory.Amt++
			if memory.Amt == 3 {
				leftoverSet, leftovers = onThreeInMemory(leftoverSet, leftovers)
				memory.ClearMemory()
			}
			if IsIngredient(trimmedText) {
				memory.Amt_Correct++
				fmt.Printf("%s registered!\n", trimmedText)
			}
		}
	})
	bmodel, _ := bayesian.NewClassifierFromFile("model/model.mo")

	search("https://www.allrecipes.com/recipe/218057/chicken-enchilada-slow-cooker-soup/", c, bmodel)
}

func PrintAllInSlice(s []string) {
	for _, value := range s {
		fmt.Println(value)
	}
}

func onThreeInMemory(leftoverSet bool, leftovers [3]string) (bool, [3]string) {
	if memory.Amt_Correct == 2 {
		leftoverSet = true // May have a leftover set
		leftovers = memory.Items
	}

	if leftoverSet && memory.Amt_Correct <= 2 {
		Ings = memory.ReturnLeftovers(Ings)
		leftoverSet = false // They were indeed leftovers
	} else {
		// Wasn't leftovers, most likely a false positive
		for _, value := range leftovers {
			Ings = AddToIngSlice(value)
		}
		leftoverSet = false
	}
	return leftoverSet, leftovers
}

func AddToIngSlice(ing string) []string {
	Ings = append(Ings, ing)
	return Ings
}

func (memory Memory) AddToMemory(ing string) {
	memory.Items[2] = memory.Items[1]
	memory.Items[1] = memory.Items[0]
	memory.Items[0] = ing
}

func search(link string, col *colly.Collector, bModel *bayesian.Classifier) {
	col.Visit(link)
	fmt.Println("Ended search")
	PrintAllInSlice(Ings)
}
