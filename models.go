package main

type Definition struct {
	Truth []int
	Name  string
}

func ConjunctionCall() SpecFun {
	return func(a, b int) bool {
		return Conjunction(a, b)
	}
}

func DisjunctionCall() SpecFun {
	return func(a, b int) bool {
		return Disjunction(a, b)
	}
}

func EqualCall() SpecFun {
	return func(a, b int) bool {
		return Equals(a, b)
	}
}

func ImplicationCall() SpecFun {
	return func(a, b int) bool {
		return Implication(a, b)
	}
}
