package main

import (
	"fmt"
	"math"
)

func IsPrimeExample(n int) {
	result := isPrime(n)
	if result {
		fmt.Println("простое")
		return
	}

	fmt.Println("составное")
}

func isPrime(n int) bool {
	prime := true
	for i := 2.0; i <= math.Sqrt(float64(n)); i++ {
		if n%int(i) == 0 {
			prime = false
			break
		}
	}

	return prime
}
