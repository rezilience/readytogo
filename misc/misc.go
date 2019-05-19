package main

import (
	"fmt"
)

func main() {

	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			done <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	fmt.Println("Hello Go!")
}

func printSlice(s []byte) {
	fmt.Printf("len=%d cap=%d; %v\n", len(s), cap(s), s)
}
