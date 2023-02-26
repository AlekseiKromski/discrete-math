package main

import "log"

func PirsOperation(a, b []int) []int {
	result, err := Operation(a, b, DisjunctionCall())
	if err != nil {
		log.Fatal("Error in conjunction operation")
	}

	return ReverseOperation(result)
}
