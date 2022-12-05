package day03

import (
	"byjacob/adventureofcode2022/utils"
	"fmt"
	"strings"
)

func calculateRunePriority(item string) int {
	item1Rune := []rune(item)[0]
	var priority int
	if item1Rune >= 97 {
		priority = int(item1Rune) - 96
	} else {
		priority = int(item1Rune) - 38
	}
	return priority
}

func part1(day int) {
	content := utils.GetDataPerDay(day, false)
	sumPrioritiy := 0
	for _, rucksack := range strings.Split(content, "\n") {
		compartment1 := strings.Split(rucksack[:len(rucksack)/2], "")
		compartment2 := strings.Split(rucksack[len(rucksack)/2:], "")
		compartment1 = utils.RemoveDuplicate(compartment1)
		compartment2 = utils.RemoveDuplicate(compartment2)
		for _, item1 := range compartment1 {
			if utils.StringInSlice(item1, compartment2) {
				sumPrioritiy += calculateRunePriority(item1)
			}
		}
	}
	fmt.Println(sumPrioritiy)
}

func part2(day int) {
	content := utils.GetDataPerDay(day, false)
	rucksacks := strings.Split(content, "\n")
	sumPrioritiy := 0
	for i := 0; i < len(rucksacks); i = i + 3 {
		ruskack1 := utils.RemoveDuplicate(strings.Split(rucksacks[i], ""))
		ruskack2 := utils.RemoveDuplicate(strings.Split(rucksacks[i+1], ""))
		ruskack3 := utils.RemoveDuplicate(strings.Split(rucksacks[i+2], ""))
		for _, item1 := range ruskack1 {
			if utils.StringInSlice(item1, ruskack2) && utils.StringInSlice(item1, ruskack3) {
				sumPrioritiy += calculateRunePriority(item1)
			}
		}
	}
	fmt.Println(sumPrioritiy)
}

func Execute() {
	day := 3

	part1(day)
	part2(day)

}
