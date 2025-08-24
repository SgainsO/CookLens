package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

var cookingVerbs []string = []string{}
var tools []string = []string{}
var timeMarkers []string = []string{}

func AssignWordLists() {
	cookingVerbs, _ = loadList("dictionaries/verbs.txt")
	tools, _ = loadList("dictionaries/tools.txt")
	timeMarkers, _ = loadList("dictionaries/time.txt")
}

func checkWordInArray(word string, array []string) bool {
	for _, item := range array {
		if item == word {
			return true
		}
	}
	return false
}

func IsInstruction(instructionParagraph string) bool{
	var total int8= 0
	var confIns int8 = 0
	newSlice := customSplit(instructionParagraph, []byte{';', '.'})
	for _, ins := range newSlice {
		if checkSentence(ins) {
			confIns++
		}
		total++
	}
	strug := float32(confIns) / float32(total)
	fmt.Println(strug)
	if strug > float32(0.5) {
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

func checkSentence(ins string) bool {
	targetsMet := 0
	ins_arr := strings.Split(ins, " ")
	for _, word := range ins_arr {
		word = strings.ToLower(word)
		if checkWordInArray(word, cookingVerbs){
			fmt.Println(word)
			targetsMet += 2
		}
		if checkWordInArray(word, tools){
			fmt.Println(word)
			targetsMet += 1
		}
		if checkWordInArray(word, timeMarkers){
			fmt.Println(word)
			targetsMet += 1
		}
	}
	if targetsMet >= 3 {
		return true
	}
	return false
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

func checkInDelimiter(c rune, delim []byte) bool{
	for _, char := range delim {
		if char == byte(c) {
			return true
		}
	}
	return false
}
