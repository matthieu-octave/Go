package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func(int) int {
	fibo := []int{0, 1}
	return func(x int) int {
		if x < 1 {
			return 0
		}
		fibo = append(fibo, (fibo[x] + fibo[x-1]))
		return (fibo[x])

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
	//fmt.Println(f(0))

}
