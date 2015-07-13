package main

import (
	"fmt"
	"strings"

	"books/programming-in-go/ch01/stacker/stack"
)

func main() {
	var heystack stack.Stack
	heystack.Push("hay")
	heystack.Push(-15)
	heystack.Push([]string{"pin", "clip", "needle"})
	heystack.Push(81.52)

	for {
		item, err := heystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	var aStack stack.Stack
	aStack.Push("Aarvark")
	aStack.Push(5)
	aStack.Push(19)
	x, err := aStack.Top()
	fmt.Println(x)
	aStack.Push(-6e-4)
	aStack.Push("Baker")
	aStack.Push(-3)
	aStack.Push("Cake")
	aStack.Push("Dancer")
	x, err = aStack.Top()
	fmt.Println(x)
	aStack.Push(11.7)
	fmt.Println("stack is empty", aStack.IsEmpty())
	fmt.Printf("Len() == %d  Cap == %d\n", aStack.Len(), aStack.Cap())
	difference := aStack.Cap() - aStack.Len()
	for i := 0; i < difference; i++ {
		aStack.Push(strings.Repeat("*", difference-i))
	}
	fmt.Printf("Len() == %d  Cap == %d\n", aStack.Len(), aStack.Cap())
	for aStack.Len() > 0 {
		x, _ = aStack.Pop()
		fmt.Printf("%T %v\n", x, x)
	}
	fmt.Println("stack is empty", aStack.IsEmpty())
	x, err = aStack.Pop()
	fmt.Println(x, err)
	x, err = aStack.Top()
	fmt.Println(x, err)
}
