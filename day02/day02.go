package day02

import (
	"byjacob/adventureofcode2022/utils"
	"fmt"
	"strings"
)

func getMapSymbols() map[string]string {
	mapSymbols := make(map[string]string)
	mapSymbols["A"] = "Rock"
	mapSymbols["X"] = "Rock"
	mapSymbols["B"] = "Paper"
	mapSymbols["Y"] = "Paper"
	mapSymbols["C"] = "Scissor"
	mapSymbols["Z"] = "Scissor"
	return mapSymbols
}

func getSymbolPoints() map[string]int {
	mapSymbols := make(map[string]int)
	mapSymbols["Rock"] = 1
	mapSymbols["Paper"] = 2
	mapSymbols["Scissor"] = 3
	return mapSymbols
}

func getRoundPoints(ch1 string, ch2 string) int {
	if ch1 == "Paper" && ch2 == "Scissor" {
		return 6
	} else if ch1 == "Scissor" && ch2 == "Rock" {
		return 6
	} else if ch1 == "Rock" && ch2 == "Paper" {
		return 6
	} else if ch1 == ch2 {
		return 3
	} else {
		return 0
	}
}

func getCh2(ch1 string, status string) string {

	mapWin := make(map[string]string)
	mapWin["Rock"] = "Paper"
	mapWin["Paper"] = "Scissor"
	mapWin["Scissor"] = "Rock"

	mapLoose := make(map[string]string)
	mapLoose["Rock"] = "Scissor"
	mapLoose["Paper"] = "Rock"
	mapLoose["Scissor"] = "Paper"

	if status == "Y" {
		return ch1
	} else if status == "X" {
		return mapLoose[ch1]
	} else {
		return mapWin[ch1]
	}
}

func part1(day int) {

	mapSymbols := getMapSymbols()
	mapPoints := getSymbolPoints()

	content := utils.GetDataPerDay(day)

	points := 0

	for _, round := range strings.Split(content, "\n") {
		chooses := strings.Split(round, " ")
		choose1 := mapSymbols[chooses[0]]
		choose2 := mapSymbols[chooses[1]]
		points += mapPoints[choose2]
		points += getRoundPoints(choose1, choose2)
	}

	fmt.Println(points)
}

func part2(day int) {
	mapSymbols := getMapSymbols()
	mapPoints := getSymbolPoints()

	content := utils.GetDataPerDay(day)

	points := 0

	for _, round := range strings.Split(content, "\n") {
		chooses := strings.Split(round, " ")
		choose1 := mapSymbols[chooses[0]]
		status := chooses[1]
		choose2 := getCh2(choose1, status)
		points += mapPoints[choose2]
		points += getRoundPoints(choose1, choose2)
	}

	fmt.Println(points)
}

func Execute() {
	day := 2

	part1(day)
	part2(day)

}
