package main

import "log"

func Practice4(a, b, c []int) []int {

	//First step -> disjunction between A & B
	firstStepResult, err := Operation(a, b, DisjunctionCall())
	if err != nil {
		log.Fatal("Error in disjunction operation")
	}

	printTable(
		"disjunction between A and B",
		Definition{
			Truth: firstStepResult,
			Name:  "|",
		},
		Definition{
			Truth: a,
			Name:  "A",
		},
		Definition{
			Truth: b,
			Name:  "B",
		},
	)

	//Second step -> make reverse C
	secondStepResult := ReverseOperation(c)

	printTable(
		"make reverse C",
		Definition{
			Truth: secondStepResult,
			Name:  "!",
		},
		Definition{
			Truth: c,
			Name:  "C",
		},
	)

	//Third step -> conjunction
	thirdStepResult, err := Operation(secondStepResult, b, ConjunctionCall())
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}

	printTable(
		"conjunction between last operation (reverse -> 2) and B",
		Definition{
			Truth: thirdStepResult,
			Name:  "^",
		},
		Definition{
			Truth: secondStepResult,
			Name:  "2",
		},
		Definition{
			Truth: b,
			Name:  "B",
		},
	)

	//Fourth step -> XOR -> anti-disjunction -> Pirs arrow
	fourthStepResult := PirsOperation(thirdStepResult, a)

	printTable(
		"You Pirs arrow (XOR) between last operation (3) and A",
		Definition{
			Truth: fourthStepResult,
			Name:  "XOR",
		},
		Definition{
			Truth: thirdStepResult,
			Name:  "3",
		},
		Definition{
			Truth: a,
			Name:  "A",
		},
	)

	//Fiv step -> implication
	fivStepResult, err := Operation(firstStepResult, fourthStepResult, ImplicationCall())
	if err != nil {
		log.Fatal("Error in implication operation")
	}

	return fivStepResult
}
