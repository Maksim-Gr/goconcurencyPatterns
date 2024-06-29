package main

import "fmt"

func main() {
	var i int
	f := func() int {
		i++
		return i
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	select {
	case ch1 <- f():
	case ch2 <- f():
	default:

	}
	fmt.Println(i)
}
