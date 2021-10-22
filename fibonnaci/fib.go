package main

func Fibonacci() func() int {
	a, b := 1, 0
	f := func() int {
		a, b = b, a+b
		return a
	}
	return f
}
