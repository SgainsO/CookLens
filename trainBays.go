package main

import (
	"fmt"
	"github.com/jbrukh/bayesian"
	"os"
	"bufio"
	"strings"
)

const (
	Good bayesian.Class = "Ing"
	Bad bayesian.Class = "notIng"
)

 var ValueMaps = map[int]string{0:"Ing", 1:"notIng"}

func trainBayes() {
	notIngri := generateStringSlices("notIngredient.txt")
	Ingri := generateStringSlices("Ingredient.txt")

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
		if value == "Ing"{ings += 1} else if value == "notIng"{notIngs += 1}

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
