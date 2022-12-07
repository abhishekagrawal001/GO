package main

import "fmt"

func fib(n int) {
	a, b := 0, 1
	c := 0
	// fmt.Printf("%v %v ", a,b)
	for i := 1; i <= n; i++ {
		fmt.Printf("%v ", a)
		c = a + b;
		a = b;
		b = c;
	}
}
func main() {
	fib(10)
}