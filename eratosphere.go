package main

import "fmt"

func EratosphereExample(n int) {
	result := eratosphere(n)
	fmt.Println(result)
}

func eratosphere(n int) []int {
	var a []int
	for i := 0; i < n+1; i++ {
		a = append(a, i)
	}

	a[1] = 0

	for i := 2; i <= n; i++ {
		if a[i] != 0 {
			j := i + i

			for {
				if j > n {
					break
				}
				a[j] = 0

				j += i
			}
		}
	}

	var cleared []int

	for _, value := range a {
		if value != 0 {
			cleared = append(cleared, value)
		}
	}
	return cleared
}
