package main

import (
	"fmt"
	"gotop/internal/grid"
	"gotop/internal/iteration"
	"time"
)

func main() {

	grid.HideCursor()
	fmt.Print("\033[H\033[2J")
	for {
		iteration.RunIteration()
		time.Sleep(1 * time.Second)
	}
}
