package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello sir")
	sum := 1
	for sum < 10 {
		fmt.Println("------")
		sum += 1
	}
	os.Exit(1)
}
