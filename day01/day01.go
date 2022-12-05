package day01

import (
	"byjacob/adventureofcode2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part1(day int) {
	content := utils.GetDataPerDay(day)
	content = strings.Replace(content, "\n", ",", -1)
	content = strings.Replace(content, ",,", "\n", -1)
	max_calories := 0

	for _, elf_bag := range strings.Split(content, "\n") {
		sum_calories := 0
		for _, calorie := range strings.Split(elf_bag, ",") {
			intVar, _ := strconv.Atoi(calorie)
			sum_calories += intVar
		}

		if sum_calories > max_calories {
			max_calories = sum_calories
		}
	}

	fmt.Println(max_calories)
}

func part2(day int) {
	content := utils.GetDataPerDay(day)
	content = strings.Replace(content, "\n", ",", -1)
	content = strings.Replace(content, ",,", "\n", -1)
	var elf_calories []int

	for _, elf_bag := range strings.Split(content, "\n") {
		sum_calories := 0
		for _, calorie := range strings.Split(elf_bag, ",") {
			intVar, _ := strconv.Atoi(calorie)
			sum_calories += intVar
		}
		elf_calories = append(elf_calories, sum_calories)
	}

	sort.Ints(elf_calories)
	bestElfCalories := elf_calories[len(elf_calories)-3:]
	sumBestElfValories := 0
	for _, calorie := range bestElfCalories {
		sumBestElfValories += calorie
	}
	fmt.Println(sumBestElfValories)
}

func Execute() {
	day := 1

	part1(day)
	part2(day)

}
