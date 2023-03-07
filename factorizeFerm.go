package main

import (
	"fmt"
	"math"
)

func FactorizeFermExample(n int) {
	a, b := FactorizeFerm(n)
	fmt.Println(a, b)
}

func FactorizeFerm(n int) (int, int) {
	x := int(math.Pow(float64(n), 0.5)) + 1

	for {
		if !checkPerfectSquare(x*x - n) {
			x += 1
			continue
		}
		break
	}

	y := int(math.Pow(float64(x*x-n), 0.5))
	a := x - y
	b := x + y

	return a, b
}

func checkPerfectSquare(x int) bool {
	z := int(math.Pow(float64(x), 0.5))

	return math.Pow(
		float64(z),
		2,
	) == float64(x)
}
