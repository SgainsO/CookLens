package scrape

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cookingVerbs []string = []string{}
var tools []string = []string{}
var timeMarkers []string = []string{}

func AssignWordLists() {
	cookingVerbs, _ = loadList("dictionaries/verbs.txt")
	tools, _ = loadList("dictionaries/tools.txt")
	timeMarkers, _ = loadList("dictionaries/time.txt")

	fmt.Println("Printing Loaded Tools")
}

func checkWordInArray(word string, array []string) bool {
	for _, item := range array {
		if item == word {
			return true
		}
	}
	return false
}

func IsInstruction(instructionParagraph string, savedInstructions *map[string]int) bool {
	var total int8 = 0
	var confIns int8 = 0
	var repitions int8 = 0

	fmt.Printf("Whole Paragraph: %s\n", instructionParagraph)

	newSlice := customSplit(instructionParagraph, []byte{';', '.'})
	for _, ins := range newSlice {
		fmt.Printf("Instruction: %s", ins)
		if _, exists := (*savedInstructions)[ins]; exists {
			repitions++
		} else if checkSentence(ins) {
			confIns++
			(*savedInstructions)[ins] = 1
			fmt.Printf(": %d\n", 1)
		} else {
			fmt.Printf(": %d\n", 0)
		}
		total++
	}

	strug := float32(confIns) / float32(total)
	strugReps := float32(repitions) / float32(total)
	fmt.Printf("Instruction Confidence: %.2f\n", strug)
	//fmt.Println(strug)
	// Try not to include Repitions
	if strugReps >= float32(0.33) {
		return false
	} else if strug >= float32(0.5) {
		return true
	} else {
		return false
	}
}

func loadList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func refineWord(word string) string {
	word = strings.ToLower(word)
	refined := strings.TrimRight(word, ",.!?;:")
	return refined
}

func checkSentence(ins string) bool {
	targetsMet := 0
	ins_arr := strings.Split(ins, " ")
	if ins_arr[0] == "add" {
		return true //Chances are far more likely than not that it is a recipe
	}
	fmt.Print("\nWords Found: ")
	for _, word := range ins_arr {
		word = refineWord(word)

		if checkWordInArray(word, cookingVerbs) {
			fmt.Printf("%s ", word)
			//		fmt.Println(word)
			targetsMet += 2
		}
		if checkWordInArray(word, tools) {
			fmt.Printf("%s ", word)
			//		fmt.Println(word)
			targetsMet += 1
		}
		if checkWordInArray(word, timeMarkers) {
			//		fmt.Println(word)
			fmt.Printf("%s ", word)
			targetsMet += 1
		}
	}

	return targetsMet >= 3
}

func customSplit(input string, delimters []byte) []string {
	newSlice := []string{}
	trackString := ""
	for _, char := range input {
		if checkInDelimiter(char, delimters) {
			newSlice = append(newSlice, trackString)
			trackString = ""
		} else {
			trackString += string(char)
		}
	}
	return newSlice
}

func checkInDelimiter(c rune, delim []byte) bool {
	for _, char := range delim {
		if char == byte(c) {
			return true
		}
	}
	return false
}
