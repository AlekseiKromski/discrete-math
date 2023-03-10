package main

import "log"

func ShefferOperation(a, b []int) []int {
	result, err := Operation(a, b, ConjunctionCall())
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}

	return ReverseOperation(result)
}
