package main

import (
	"fmt"
)

// Table of Truth
var matrix = [][]int{
	{
		1, 0, 0, 1,
	},
	{
		1, 0, 1, 0,
	},
}

var A = Definition{
	Truth: matrix[0],
	Name:  "A",
}

var B = Definition{
	Truth: matrix[1],
	Name:  "B",
}

func main() {
	result := Definition{
		Truth: ShefferOperation(matrix[0], matrix[1]),
		Name:  "C",
	}

	printTable(
		"C - is a sheffer operation",
		result,
		A,
		B,
	)
}

func printTable(comment string, result Definition, definitions ...Definition) {

	fmt.Println(comment)

	//Create a header of table
	var total string
	for _, def := range definitions {
		total += fmt.Sprintf("| %s |", def.Name)
	}

	total += fmt.Sprintf("| %s |", result.Name)

	fmt.Println(total)
	sub(len(total))

	//Create a body of table
	for i := 0; i < len(result.Truth); i++ {
		total = ""
		for _, def := range definitions {
			total += fmt.Sprintf("| %d |", def.Truth[i])
		}
		total += fmt.Sprintf("| %d |", result.Truth[i])

		fmt.Println(total)
		sub(len(total))
	}
}

func sub(count int) {
	var result string
	for i := 0; i < count; i++ {
		result += fmt.Sprint("-")
	}
	fmt.Println(result)
}
