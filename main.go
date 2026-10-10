package main

import (
	"fmt"
	"gotop/internal"
	"gotop/internal/buffer"
	"gotop/internal/grid"
	"log"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

func main() {

	grid.HideCursor()
	fmt.Print("\033[H\033[2J")
	for {
		buff := make([]byte, 128)
		ch := make(chan int)
		mem := make(chan []int)
		var wg sync.WaitGroup
		wg.Add(2)
		wg.Go(func() {
			internal.ReadTemp(1, buff, ch)
			wg.Done()
		})
		wg.Go(func() {
			internal.ReadMem(mem)
			wg.Done()
		})
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			log.Fatal("Couldn't get size of terminal")
		}
		termBuff := buffer.NewBuffer(width, height)
		grid.RenderGrid(termBuff, width, height)
		temp := <-ch
		mems := <-mem
		_, _ = temp, mems
		grid.PrintTemp(buff, temp, termBuff)
		for i := range height {
			fmt.Printf(string(termBuff.Get(i)))
		}
		wg.Wait()
		time.Sleep(1 * time.Second)

	}
}
