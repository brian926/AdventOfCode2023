package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}

func findFirstLast(line string) (string, string, error) {
	var f, l string
	var foundF = false
	var min, max int

	input := containsWords(line)

	for i, char := range input {
		if unicode.IsDigit(char) {
			if !foundF {
				f = string(char)
				l = string(char)
				foundF = true
			} else if foundF && i > max {
				l = string(char)
				max = i
			} else if foundF && i < min {
				f = string(char)
				min = i
			}
		}
	}
	fmt.Printf("First: %s last: %s word: %s origin: %s", f, l, input, line)
	fmt.Println()
	return f, l, nil
}

func containsWords(input string) string {
	myMap := make(map[string]string)
	myMap["one"] = "one1one"
	myMap["two"] = "two2two"
	myMap["three"] = "three3three"
	myMap["four"] = "four4four"
	myMap["five"] = "five5five"
	myMap["six"] = "six6six"
	myMap["seven"] = "seven7seven"
	myMap["eight"] = "eight8eight"
	myMap["nine"] = "nine9nine"

	words := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, word := range words {
		input = strings.Replace(input, word, myMap[word], -1)
	}
	return input
}

func main() {
	filePath := "inputs.txt"

	lines, err := readFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, line := range lines {
		first, last, err := findFirstLast(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		temp := first + last
		tempI, err := strconv.Atoi(temp)
		if err != nil {
			fmt.Println(err)
			return
		}

		total += tempI
	}

	fmt.Println(total)
}
