package main

import "fmt"

func FactorizeExample(n int) {
	result := eratosphere(n)
	var factors []int
	for _, value := range result {
		if n%value == 0 {
			factors = append(factors, value)
		}
	}

	fmt.Println(factors)
}
