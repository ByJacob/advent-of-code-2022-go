package day04

import (
	"byjacob/adventureofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func generateElfRange(strRange string) []string {
	borders := strings.Split(strRange, "-")
	start, _ := strconv.Atoi(borders[0])
	end, _ := strconv.Atoi(borders[1])
	var res []string
	for i := start; i <= end; i++ {
		res = append(res, strconv.Itoa(i))
	}
	return res
}

func part1(day int) {
	content := utils.GetDataPerDay(day, false)
	result := 0
	for _, line := range strings.Split(content, "\n") {
		elfs := strings.Split(line, ",")
		elf1 := generateElfRange(elfs[0])
		elf2 := generateElfRange(elfs[1])
		repeatCount := 0
		for _, section := range elf1 {
			if utils.StringInSlice(section, elf2) {
				repeatCount += 1
			}
		}
		if repeatCount == len(elf1) || repeatCount == len(elf2) {
			result += 1
		}
	}
	fmt.Println(result)
}

func part2(day int) {
	content := utils.GetDataPerDay(day, false)
	result := 0
	for _, line := range strings.Split(content, "\n") {
		elfs := strings.Split(line, ",")
		elf1 := generateElfRange(elfs[0])
		elf2 := generateElfRange(elfs[1])
		repeatCount := 0
		for _, section := range elf1 {
			if utils.StringInSlice(section, elf2) {
				repeatCount += 1
			}
		}
		if repeatCount > 0 {
			result += 1
		}
	}
	fmt.Println(result)
}

func Execute() {
	day := 4

	part1(day)
	part2(day)

}
