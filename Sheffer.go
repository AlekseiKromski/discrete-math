package main

import "log"

func ShefferOperation(a, b []int) []int {
	result, err := Operation(matrix[0], matrix[1], ConjunctionCall())
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}

	return ReverseOperation(result)
}

func ConjunctionCall() SpecFun {
	return func(a, b int) bool {
		return Conjunction(a, b)
	}
}
