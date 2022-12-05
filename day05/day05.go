package day05

import (
	"byjacob/adventureofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseStockGraph(strStock string) map[int][]string {
	stackPart := strings.Split(strStock, "\n")
	stacks := make(map[int][]string)
	for i := len(stackPart) - 2; i >= 0; i-- {
		line := stackPart[i]
		for j := 0; j < len(line); j += 4 {
			stackNumber := j / 4
			if line[j+1] != ' ' {
				stacks[stackNumber] = append(stacks[stackNumber], string(line[j+1]))
			}
		}
	}
	return stacks
}

func parseMoves(strMoves string) [][]int {
	var r [][]int
	for _, line := range strings.Split(strMoves, "\n") {
		parts := strings.Split(line, " ")
		move, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		r = append(r, []int{move, from - 1, to - 1})
	}
	return r
}

func part1(day int) {
	content := utils.GetDataPerDay(day, false)
	parts := strings.Split(content, "\n\n")
	stacks := parseStockGraph(parts[0])
	moves := parseMoves(parts[1])
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			last_idx_from := len(stacks[move[1]]) - 1
			crate := stacks[move[1]][last_idx_from]
			stacks[move[1]] = stacks[move[1]][:last_idx_from]
			stacks[move[2]] = append(stacks[move[2]], crate)
		}
	}
	result := ""
	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		result += stack[len(stack)-1]
	}
	fmt.Println(result)
}

func part2(day int) {
	content := utils.GetDataPerDay(day, false)
	parts := strings.Split(content, "\n\n")
	stacks := parseStockGraph(parts[0])
	moves := parseMoves(parts[1])
	for _, move := range moves {
		last_idx_from := len(stacks[move[1]])
		movePart := stacks[move[1]][last_idx_from-move[0] : last_idx_from]
		stacks[move[1]] = stacks[move[1]][:last_idx_from-move[0]]
		stacks[move[2]] = append(stacks[move[2]], movePart...)
	}
	result := ""
	for i := 0; i < len(stacks); i++ {
		stack := stacks[i]
		result += stack[len(stack)-1]
	}
	fmt.Println(result)
}

func Execute() {
	day := 5

	part1(day)
	part2(day)

}
