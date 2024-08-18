package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// stringStream := make(chan string)

	// go func() {
	// 	stringStream <- "Channel use"
	// }()

	// fmt.Println(<-stringStream)

	// //Channels are always blocking, use if ok style for checks

	// intStream := make(chan int)
	// close(intStream)

	// integer, ok := <-intStream
	// fmt.Printf("(%v): %v", ok, integer)

	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	someIntStream := make(chan int, 4)
	go func() {
		defer close(someIntStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			someIntStream <- i
		}
	}()
	for integer := range someIntStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
