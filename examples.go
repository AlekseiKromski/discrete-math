package main

import (
	"fmt"
	"log"
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

func ShefferExample() {
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

func PirsExample() {
	result := Definition{
		Truth: PirsOperation(matrix[0], matrix[1]),
		Name:  "C",
	}

	printTable(
		"C - is a disjunction operation",
		result,
		A,
		B,
	)
}

func Practice3() {

	//Left side
	result, err := Operation(matrix[0], matrix[1], EqualCall())
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}
	leftSide := Definition{
		Truth: result,
		Name:  "C",
	}

	printTable("C - is a left side", leftSide, A, B)

	//Right side
	implication1, err := Operation(matrix[0], matrix[1], ImplicationCall())
	if err != nil {
		log.Fatal("Error in implication operation")
	}
	implication2, err := Operation(matrix[1], matrix[0], ImplicationCall())
	if err != nil {
		log.Fatal("Error in implication operation")
	}

	result, err = Operation(implication1, implication2, ConjunctionCall())
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}
	rightSide := Definition{
		Truth: result,
		Name:  "C",
	}

	printTable("C - is a right side", rightSide, Definition{Truth: implication1, Name: "A"}, Definition{Truth: implication2, Name: "B"})

	//compare two results
	if len(leftSide.Truth) != len(rightSide.Truth) {
		panic("LEFT SIDE AND RIGHT SIDE IS NOT THE EQUAL")
	}

	same := true
	for i := 0; i < len(leftSide.Truth); i++ {
		if leftSide.Truth[i] != rightSide.Truth[i] {
			same = false
		}
	}

	if !same {
		log.Fatal("LEFT SIDE AND RIGHT SIDE ARE NOT EQUAL")
	}

	fmt.Println("LEFT SIDE AND RIGHT SIDE -> ARE EQUAL")
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
