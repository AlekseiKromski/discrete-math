package main

import "log"

func ShefferOperation(a, b []int) []int {
	result, err := ConjunctionOperation(matrix[0], matrix[1])
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}

	return ReverseOperation(result)
}
