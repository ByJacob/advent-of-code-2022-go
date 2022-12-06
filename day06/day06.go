package day06

import (
	"byjacob/adventureofcode2022/utils"
	"fmt"
	"strings"
)

func findMessage(uniqNumber int, content string) int {
	searchNumber := 0
	for i := 0; i < len(content)-1-uniqNumber; i++ {
		testPart := content[i : i+uniqNumber]
		if len(utils.RemoveDuplicate(strings.Split(testPart, ""))) == len(testPart) {
			searchNumber = i + uniqNumber
			break
		}
	}
	return searchNumber
}

func part1(day int) {
	content := utils.GetDataPerDay(day, false)
	fmt.Println(findMessage(4, content))
}

func part2(day int) {
	content := utils.GetDataPerDay(day, false)
	fmt.Println(findMessage(14, content))
}

func Execute() {
	day := 6

	part1(day)
	part2(day)

}
