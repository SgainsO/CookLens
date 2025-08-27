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


func (m Memory) Amt_Correct(typeNum int8) int {
	newAmt := 0
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

func (m *Memory) DeepCopyMemory(toCopy Memory) {
	for i := 0; i < 3; i++ {
		m.Items[i] = toCopy.Items[i]
		m.Item_Type[i] = toCopy.Item_Type[i]
	}
	m.Amt = toCopy.Amt
}

func (m *Memory) ClearMemory() {
	m.Items = [3]string{}
	m.Amt = 0
//	fmt.Println("Memory cleared")
}

var memory Memory = Memory{Items: [3]string{},
	Item_Type: [3]int8{}, Amt: 0}

var Leftovers Memory = Memory{Items: [3]string{"", "", ""},
	Item_Type: [3]int8{0, 0, 0}}

var Ings []string = []string{}
var Recipe []string = []string{}
func main() {
	AssignWordLists()
	LoadPositives()
	fmt.Println("Input a Link you want to find the recipe of!")
	c := colly.NewCollector()

	leftoverPossible := false

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
			} else if IsInstruction(trimmedText) {
				memory.AddToMemory(trimmedText, 2)
				fmt.Printf("%s INSTRUCTION \n", trimmedText)
			} else {
				memory.AddToMemory(trimmedText, 0)
			//	fmt.Printf("%s other\n", trimmedText)
			}


			//Always will run regardless of corectness
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
				fmt.Printf("Before entering\n")
				leftoverPossible = handleLeftovers(leftoverPossible)
				memory.ClearMemory()
			}
		}
	})
	bmodel, _ := bayesian.NewClassifierFromFile("model/model.mo")

	search("https://www.recipetineats.com/mexican-corn-salad/", c, bmodel)
}

func PrintAllInSlice(s []string) {
	for _, value := range s {
		fmt.Println(value)
	}
}

func putAllInArrayIntoSlice(list []string, array []string) {
	for _, value := range array {
		list = AddToSlice(value, list)
	}
}

func handleLeftovers(leftoverSet bool) (bool) {

	//This patch runs first, will always be wrong
	fmt.Printf("%d %d %d correct", leftoverSet, memory.Amt_Correct(2), Leftovers.Amt_Correct(2))
	if leftoverSet {
		for kind := int8(1); kind <= 2; kind++ {
			for i := 0; i < memory.Amt_Correct(kind); i++ {
				if Leftovers.Items[i] != "" {
					if kind == 1 {
						Ings = AddToSlice(Leftovers.Items[i], Ings)
					} else {
						Recipe = AddToSlice(Leftovers.Items[i], Recipe)
					}
				}
			}

			if memory.Amt_Correct(kind) == 0 && Leftovers.Amt_Correct(kind) > 1 {
				for i := int(Leftovers.Amt) - 1; i >= int(0); i-- {
					if Leftovers.Items[i] != "" && Leftovers.Item_Type[i] == int8(kind) {
						if kind == 1 {
							Ings = AddToSlice(Leftovers.Items[i], Ings)
						} else {
							Recipe = AddToSlice(Leftovers.Items[i], Recipe)
						}
					}
				}
			}
		}
	}

	leftoverSet = false


	if memory.Amt_Correct(1) >= 1 && memory.Amt_Correct(1) < 3 ||
	memory.Amt_Correct(2) >= 1 && memory.Amt_Correct(2) < 3  {
		leftoverSet = true // May have a leftover set
		Leftovers.DeepCopyMemory(memory)
	}

	return leftoverSet
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
