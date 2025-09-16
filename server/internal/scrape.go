package scrape

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type ItemHolder struct {
	item   string
	number int
}

type Memory struct {
	Amt       int8
	Items     [3]ItemHolder
	Item_Type [3]int8 //0  is not Ing, 1 is Ing
}

func (m Memory) Amt_Correct(typeNum int8) int {
	newAmt := 0
	for _, value := range m.Item_Type {
		if value == typeNum {
			newAmt++
		}
	}
	return newAmt
}

func (m Memory) ReturnLeftovers(s []string, toRet int8) []string {
	for index, value := range m.Items {
		if m.Item_Type[index] == toRet {
			s = AddToSlice(value.item, s)
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
	m.Items = [3]ItemHolder{}
	m.Amt = 0
	// fmt.Println("Memory cleared")
}

// ParseRecipeFromURL extracts ingredients and recipe steps from a URL
// This function is safe for use in coroutines as it uses local variables
func ParseRecipeFromURL(url string, ings *[]ItemHolder, recipe *[]ItemHolder) error {
	AssignWordLists()
	LoadPositives()

	c := colly.NewCollector()
	itemNumber := 0
	leftoverPossible := false

	// Local memory variables for coroutine safety
	memory := Memory{Items: [3]ItemHolder{}, Item_Type: [3]int8{}, Amt: 0}
	leftovers := Memory{Items: [3]ItemHolder{}, Item_Type: [3]int8{0, 0, 0}}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("li", func(e *colly.HTMLElement) {
		itemNumber++
		trimmedText := strings.TrimSpace(e.Text)
		if trimmedText != "" {
			if IsIngredient(trimmedText) {
				memory.AddToMemory(trimmedText, 1, itemNumber)
				fmt.Printf("%s registered!\n", trimmedText)
			} else if IsInstruction(trimmedText) {
				memory.AddToMemory(trimmedText, 2, itemNumber)
				fmt.Printf("%s INSTRUCTION \n", trimmedText)
			} else {
				memory.AddToMemory(trimmedText, 0, itemNumber)
			}

			if memory.Amt == 3 {
				if memory.Amt_Correct(1) == 3 {
					fmt.Println("This is running")
					for _, item := range memory.Items {
						*ings = AddItemToSlice(item.item, item.number, *ings)
					}
				}
				if memory.Amt_Correct(2) == 3 {
					for _, item := range memory.Items {
						*recipe = AddItemToSlice(item.item, item.number, *recipe)
					}
				}
				fmt.Printf("Before entering\n")
				leftoverPossible = handleLeftoversLocal(leftoverPossible, itemNumber, &memory, &leftovers, ings, recipe)
				memory.ClearMemory()
			}
		}
	})

	c.Visit(url)

	// Sort by order number to preserve original webpage order
	sortItemHoldersByOrder(*ings)
	sortItemHoldersByOrder(*recipe)

	return nil
}

func Scrape(url string) ([]string, []string, bool) {
	var ings []ItemHolder = []ItemHolder{}
	var recipe []ItemHolder = []ItemHolder{}

	fmt.Println("Input a Link you want to find the recipe of!")
	err := ParseRecipeFromURL(url, &ings, &recipe)
	if err != nil {
		fmt.Printf("Error parsing recipe: %v\n", err)
		return []string{}, []string{}, false
	}
	var ret_ings []string = []string{}
	var ret_rec []string = []string{}
	for ind := 0; ind < len(ings); ind++ {
		ret_ings = append(ret_ings, ings[ind].item)
	}

	for ind := 0; ind < len(recipe); ind++ {
		ret_rec = append(ret_rec, recipe[ind].item)
	}

	fmt.Println("Ended search")
	PrintAllItemHolders(ings)
	fmt.Println("-----------------")
	PrintAllItemHolders(recipe)

	return ret_ings, ret_rec, true
}

func PrintAllInSlice(s []string) {
	for _, value := range s {
		fmt.Println(value)
	}
}

func PrintAllItemHolders(s []ItemHolder) {
	for _, value := range s {
		fmt.Printf("[%d] %s\n", value.number, value.item)
	}
}

func putAllInArrayIntoSlice(list []string, array []string) {
	for _, value := range array {
		list = AddToSlice(value, list)
	}
}

func handleLeftoversLocal(leftoverSet bool, currentItemNumber int, memory *Memory, leftovers *Memory, ings *[]ItemHolder, recipe *[]ItemHolder) bool {

	//This patch runs first, will always be wrong
	fmt.Printf("%d %d %d correct", leftoverSet, memory.Amt_Correct(2), leftovers.Amt_Correct(2))
	if leftoverSet {
		for kind := int8(1); kind <= 2; kind++ {
			for i := 0; i < memory.Amt_Correct(kind); i++ {
				if leftovers.Items[i].item != "" {
					if kind == 1 {
						*ings = AddItemToSlice(leftovers.Items[i].item, leftovers.Items[i].number, *ings)
					} else {
						*recipe = AddItemToSlice(leftovers.Items[i].item, leftovers.Items[i].number, *recipe)
					}
				}
			}

			if memory.Amt_Correct(kind) == 0 && leftovers.Amt_Correct(kind) > 1 {
				for i := int(leftovers.Amt) - 1; i >= int(0); i-- {
					if leftovers.Items[i].item != "" && leftovers.Item_Type[i] == int8(kind) {
						if kind == 1 {
							*ings = AddItemToSlice(leftovers.Items[i].item, leftovers.Items[i].number, *ings)
						} else {
							*recipe = AddItemToSlice(leftovers.Items[i].item, leftovers.Items[i].number, *recipe)
						}
					}
				}
			}
		}
	}

	leftoverSet = false

	if memory.Amt_Correct(1) >= 1 && memory.Amt_Correct(1) < 3 ||
		memory.Amt_Correct(2) >= 1 && memory.Amt_Correct(2) < 3 {
		leftoverSet = true // May have a leftover set
		leftovers.DeepCopyMemory(*memory)
	}

	return leftoverSet
}

func AddToSlice(ing string, s []string) []string {
	newSlice := append(s, ing)
	return newSlice
}

func AddItemToSlice(item string, orderNum int, s []ItemHolder) []ItemHolder {
	newItem := ItemHolder{item: item, number: orderNum}
	newSlice := append(s, newItem)
	return newSlice
}

func (memory *Memory) AddToMemory(ing string, corState int8, orderNum int) {
	memory.Items[2] = memory.Items[1]
	memory.Items[1] = memory.Items[0]
	memory.Items[0] = ItemHolder{item: ing, number: orderNum}

	memory.Item_Type[2] = memory.Item_Type[1]
	memory.Item_Type[1] = memory.Item_Type[0]
	memory.Item_Type[0] = corState
	memory.Amt++
}

func sortItemHoldersByOrder(items []ItemHolder) {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j].number > items[j+1].number {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}
