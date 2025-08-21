package main

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
	"github.com/jbrukh/bayesian"
)

type Memory struct {
	Amt         int8
	Items       [3]string
	Item_Type   [3]int8 //0  is not Ing, 1 is Ing
}

func (m Memory) Amt_Correct(typeNum int8) int8 {
	newAmt := int8(0)
	for _, value := range m.Item_Type{
		if value == typeNum {
			newAmt++
		}
	}
	return newAmt
}

func (m Memory) ReturnLeftovers(s []string, toRet int8) []string {
	for index, value := range m.Items {
		if m.Item_Type[index] == toRet {
			s = AddToSlice(value, s)
		} else {
			break
		}
	}
	return s
}
func (m *Memory) ClearMemory() {
	m.Items = [3]string{}
	m.Amt = 0
	fmt.Println("Memory cleared")
}

var memory Memory = Memory{Items: [3]string{},
	Item_Type: [3]int8{}, Amt: 0}

var Ings []string = []string{}
var Recipe []string = []string{}

func main() {

	trainBayes()
	LoadPositives()
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
			if IsIngredient(trimmedText) {
				memory.AddToMemory(trimmedText, 1)
				fmt.Printf("%s registered!\n", trimmedText)
			}else {
				fmt.Println("XXX: ", trimmedText)
				memory.AddToMemory(trimmedText, 0)
			}

			if memory.Amt == 3 {
				if memory.Amt_Correct(1) == 3{
					fmt.Println("This is running")
					for _, item := range memory.Items {
						Ings = AddToSlice(item, Ings)
					}
				}
				if memory.Amt_Correct(2) == 3{
					for _, item := range memory.Items {
						Recipe = AddToSlice(item, Recipe)
					}
				}
				leftoverSet, leftovers = handleLeftovers(leftoverSet, leftovers)
				memory.ClearMemory()
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

func handleLeftovers(leftoverSet bool, leftovers [3]string) (bool, [3]string) {

	if memory.Amt_Correct(1) == 2  || memory.Amt_Correct(2) == 2 {
		leftoverSet = true // May have a leftover set
		leftovers = memory.Items
	}

	if leftoverSet || memory.Amt_Correct(1) <= 1 {
		Ings = memory.ReturnLeftovers(Ings, 1)
		leftoverSet = false // They were indeed leftovers
	}else if memory.Amt_Correct(2) <= 1 {
		Recipe = memory.ReturnLeftovers(Recipe, 2)
		leftoverSet = false // They were indeed leftovers
	}else { //This will only run when the array has three positives
		if memory.Amt_Correct(1) > memory.Amt_Correct(2) {
			for _, value := range leftovers {
				Ings = AddToSlice(value, Ings)
			} //Adds to which ever the false negative actually belonged to
		}else {
			for _, value := range leftovers {
				Recipe = AddToSlice(value, Recipe)
			}
		}

		fmt.Println("Seeing if this actually runs")

		leftoverSet = false
	}
	return leftoverSet, leftovers
}

func AddToSlice(ing string, s []string) []string {
	newSlice := append(s, ing)
	return newSlice
}

func (memory *Memory) AddToMemory(ing string, corState int8) {
	memory.Items[2] = memory.Items[1]
	memory.Items[1] = memory.Items[0]
	memory.Items[0] = ing


	memory.Item_Type[2] = memory.Item_Type[1]
	memory.Item_Type[1] = memory.Item_Type[0]
	memory.Item_Type[0] = corState
	memory.Amt++
}

func search(link string, col *colly.Collector, bModel *bayesian.Classifier) {
	col.Visit(link)
	fmt.Println("Ended search")
	PrintAllInSlice(Ings)
	fmt.Println("-----------------")
	PrintAllInSlice(Recipe)
}
