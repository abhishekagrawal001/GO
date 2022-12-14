package main

import "fmt"

func fib(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		c := a + b
		a = b
		b = c
	}
}

func main() {

	var n int
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d", &n)

	fmt.Printf("Fibonacci series:")
	fib(n)
}
