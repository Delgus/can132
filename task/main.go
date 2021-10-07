package main

import "fmt"

func main() {
	fmt.Print("1")

	go func() {
		fmt.Print("2")
	}()

	fmt.Print("3")
}
