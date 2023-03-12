package main

import (
	"fmt"
	"math"
)

func DrobExample(a, b int) {
	fmt.Println(drop(a, b, []int{}))
}

func drop(a, b int, chain []int) []int {
	in := int(math.Floor(float64(a / b)))
	chain = append(chain, in)

	ostatok := a % b
	if ostatok == 0 {
		return chain
	}

	return drop(b, ostatok, chain)
}
