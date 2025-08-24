package main

import (
	"fmt"
	"github.com/jbrukh/bayesian"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const (
	Good bayesian.Class = "Rep"
	Bad bayesian.Class = "NotRep"
)

var RecipeWords map[string]int = make(map[string]int)
var ValueMaps = map[int]string{0:"Rep", 1:"NotRep"}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64) // Try to parse as a float64
	return err == nil
}

func EvaluateSentence(Sentence string) bool{
	points := 0
	sentence := deleteBadCharacters(strings.ToLower(Sentence))
	senArray := strings.Split(sentence, " ")
	amountWord := len(senArray)

	if amountWord < 2 {
	 points += 2
	}else if  amountWord < 6 {
		points += 1			//Inbetween these answers, it just doesn't gain a point
	}else if amountWord > 10{
		return false
	}

	for _, word := range senArray {
		if RecipeWords[word] == 1 || isNumeric(word){
			points += 1
		}
	}

	if points > 3{
		return true
	}else{
		return false
	}
}

func deleteBadCharacters(sentence string) string {
	strings.NewReplacer("(", " ", ")"," ", "!", "", "#", "")
	return strings.ReplaceAll(sentence, ".", "")
}

func LoadPositives() {
	RecipeWords = fileIntoPositiveWords("dictionaries/amount.txt", RecipeWords)
	RecipeWords = fileIntoPositiveWords("dictionaries/ingre.txt", RecipeWords)
}

func fileIntoPositiveWords(path string, mapTo map[string]int) map[string]int{
	f1, _ := os.Open(path)

	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			mapTo[word] = 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return mapTo
}


func main_hide() {
	notIngri := generateStringSlices("notRecipe.txt")
	Ingri := generateStringSlices("recipe.txt")

	classifier := bayesian.NewClassifier(Good, Bad)
	classifier.Learn(notIngri, Bad)
	classifier.Learn(Ingri, Good)

	var correct int16 = 0
	var Incorrect int16 = 0
	var total int16 = 0

	var ings int8 = 0
	notIngs := 0

	testData := LoadTesting("test.txt")

	for key, value := range testData {
		total++
		_, likely, _ := classifier.LogScores(strings.Split(key, " "))
		if value == "Rep"{ings += 1} else if value == "notRep"{notIngs += 1}

		if value == ValueMaps[likely] {		//Since Classifier returns an array
			correct++
		} else {
			fmt.Println("Incorrect")
			fmt.Printf("Given Value: %s\n", key)
			fmt.Printf("Given Results: %s Model Output: %s\n", value, ValueMaps[likely])
			fmt.Printf("---------------\n")
			Incorrect++
		}
	}

	fmt.Printf("Amount of Ings: %d\n", ings)
	fmt.Printf("Amount of notIngs: %d\n", notIngs)
	fmt.Printf("Accuracy: %d/%d\n", correct, total)
	fmt.Printf("Incorrect: %d/%d\n", Incorrect, total)
	classifier.WriteToFile("model/model.mo")
}

func IsIngredientOLD(input string) bool {
	classifier, _ := bayesian.NewClassifierFromFile("model/model.mo")
	parts := strings.Split(input, " ")
	_, likely, _ := classifier.LogScores(parts)
	return likely == 0
}

func IsIngredient(input string) bool {
	return EvaluateSentence(input)
}

func generateStringSlices(fileName string) []string {
	var notIngri []string
	f1, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return []string{"ERROR"}
	}
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		notIngri = append(notIngri, scanner.Text())
	}

	defer f1.Close()
	return notIngri
}

func SeperateTest(line string) []string {
	parts := strings.Split(line, ":")
	for i:= 0;i<len(parts);i++{
		parts[i] = strings.TrimSpace(parts[i])
	}
	if len(parts) == 2 {
		return parts
	} else if len(parts) > 2 {
		toRet := make([]string, 2)
		strBuild := strings.Builder{}
		for i := 0; i < len(parts) -1; i++ {
			strBuild.WriteString(parts[i])
			if i < len(parts) -2 {
				strBuild.WriteString(":")
			}
		}
		toRet[0] = strBuild.String()
		toRet[1] = parts[len(parts)-1]
		return toRet
	}else {
		return []string{"error", ""}
	}
}


func LoadTesting(fileName string) map[string]string {
	var testData = make(map[string]string)
	f1, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return map[string]string{"error": "error"}
	}
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		split := SeperateTest(scanner.Text())
		testData[split[0]] = split[1]
	}

	defer f1.Close()
	return testData
}
