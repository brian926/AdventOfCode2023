package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// func readFile(filePath string) ([]string, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return nil, fmt.Errorf("error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var lines []string

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		lines = append(lines, line)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return nil, fmt.Errorf("error reading file: %v", err)
// 	}

// 	return lines, nil
// }

// func main() {
// 	filePath := "test.txt"

// 	lines, err := readFile(filePath)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for _, line := range lines {
// 		fmt.Println(line)
// 	}
// }

type Pull struct {
	Red   int
	Green int
	Blue  int
}

type Games map[int][]Pull

func parseInput(filename string) (Games, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	games := make(Games)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		gameID, _ := strconv.Atoi(parts[0][5:])
		games[gameID] = make([]Pull, 0)

		results := strings.Split(parts[1], "; ")
		for _, result := range results {
			cubes := strings.Split(result, ", ")
			pull := Pull{
				Red:   parseCubeCount(cubes, "red"),
				Green: parseCubeCount(cubes, "green"),
				Blue:  parseCubeCount(cubes, "blue"),
			}
			games[gameID] = append(games[gameID], pull)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func parseCubeCount(cubes []string, color string) int {
	for _, cube := range cubes {
		countAndColor := strings.Split(cube, " ")
		if countAndColor[1] == color {
			count, _ := strconv.Atoi(countAndColor[0])
			return count
		}
	}
	return 0
}

func power(pull Pull) int {
	red := pull.Red
	green := pull.Green
	blue := pull.Blue
	var total = red * green * blue
	fmt.Println(red, green, blue, total)
	fmt.Println("Returning ", total)
	return total
}

func possibleGames(games Games, cubes Pull) []int {
	var result []int

	for id, pulls := range games {
		valid := true
		for _, pull := range pulls {
			if !isValidPull(cubes, pull) {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, id)
		}
	}

	return result
}

func isValidPull(cubes Pull, pull Pull) bool {
	return cubes.Red >= pull.Red && cubes.Green >= pull.Green && cubes.Blue >= pull.Blue
}

func minPulls(games Games) []Pull {
	var result []Pull
	fmt.Println(games)

	for _, pulls := range games {
		minPull := Pull{
			Red:   math.MinInt64,
			Green: math.MinInt64,
			Blue:  math.MinInt64,
		}

		for _, pull := range pulls {
			if pull.Red > minPull.Red {
				minPull.Red = pull.Red
			}
			if pull.Green > minPull.Green {
				minPull.Green = pull.Green
			}
			if pull.Blue > minPull.Blue {
				minPull.Blue = pull.Blue
			}
		}

		result = append(result, minPull)
	}
	fmt.Println(result)
	return result
}

func main() {
	filename := "inputs.txt"
	games, err := parseInput(filename)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	cubes := Pull{12, 13, 14}
	gameIDs := possibleGames(games, cubes)
	powers := make([]Pull, len(games))

	powers = minPulls(games) // Adjust index since game IDs start from 1

	temp := make([]int, len(games))
	for id, i := range powers {
		fmt.Println(i)
		temp[id] = power(i)
	}

	fmt.Printf("Sum of possible game IDs: %d\n", sumIntSlice(gameIDs))
	fmt.Printf("Sum of minimum powers: %d\n", sumIntSlice(temp))
}

func sumIntSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}
