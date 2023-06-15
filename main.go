package main

import (
	"fmt"
	"strconv"
	"strings"
)

var o [3]string = [3]string{" + ", " - ", ""}

func main() {
	const target int = 200
	var solutions []string
	possibleOperators := []int{0, 1, 2}
	size := 9
	operatorsCombinations := generateCombinationsWithRepetitions(possibleOperators, size)
	//fmt.Println(operatorsCombinations)

	for _, i := range operatorsCombinations {
		expr := fmt.Sprintf("9%s8%s7%s6%s5%s4%s3%s2%s1%s0", o[i[0]], o[i[1]], o[i[2]], o[i[3]], o[i[4]], o[i[5]], o[i[6]], o[i[7]], o[i[8]])
		//fmt.Printf("expr: %s || eval: %v\n", expr, eval(expr))
		if target == eval(expr) {
			solutions = append(solutions, expr)
		}
	}

	fmt.Println("valid expressions (equal 200):")
	for _, s := range solutions {
		fmt.Printf("%s\n", s)
	}
}

func generateCombinationsWithRepetitions(items []int, length int) [][]int {
	if length == 0 {
		return [][]int{{}}
	} else {
		combinations := [][]int{}
		for _, item := range items {
			subCombinations := generateCombinationsWithRepetitions(items, length-1)
			for _, subCombination := range subCombinations {
				combinations = append(combinations, append([]int{item}, subCombination...))
			}
		}
		return combinations
	}
}

func eval(expression string) int {
	tokens := strings.Split(expression, " ")
	result, _ := strconv.Atoi(tokens[0])

	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		operand, _ := strconv.Atoi(tokens[i+1])

		if operator == "+" {
			result += operand
		} else if operator == "-" {
			result -= operand
		}
	}

	return result
}
