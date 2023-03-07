package main

import "fmt"

func EphclidFunc(a, b int) {
	for {
		if a != 0 && b != 0 {
			if a > b {
				a = a % b
				continue
			}
			b = b % a
		} else {
			break
		}
	}

	fmt.Println(a + b)
}
