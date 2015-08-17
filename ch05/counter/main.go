package main

import "fmt"

func main() {
	counterA := createCounter(2)   // counterA is of type chan int
	counterB := createCounter(102) // counterB is of type chan int

	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A→%d, B→%d) ", a, <-counterB)
	}

	fmt.Println()
}

func createCounter(start int) chan int {
	next := make(chan int)

	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)

	return next
}
