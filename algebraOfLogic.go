package main

import "fmt"

type SpecFun func(a, b int) bool

func Operation(a, b []int, operation SpecFun) ([]int, error) {
	if !checkTable(a, b) {
		return nil, fmt.Errorf("the tables have different lenght")
	}

	var result []int
	for i := 0; i < len(a); i++ {
		result = append(result, convertToInt(operation(a[i], b[i])))
	}

	return result, nil
}

func Conjunction(a, b int) bool { return convertToBool(a) && convertToBool(b) }

func ReverseOperation(data []int) []int {
	var tmp []int
	for i := 0; i < len(data); i++ {
		tmp = append(tmp, Reverse(data[i]))
	}

	return tmp
}

func Reverse(x int) int {
	return convertToInt(convertToBool(x) != convertToBool(1))
}

func convertToBool(x int) bool { return x != 0 }

func convertToInt(x bool) int {
	if x {
		return 1
	}
	return 0
}

func checkTable(a, b []int) bool { return len(a) == len(b) }
