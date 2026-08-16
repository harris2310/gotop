package main

import (
	"fmt"
	"gotop/internal"
	"log"
	"sync"
	"time"

	"golang.org/x/term"
)

func main() {

	internal.HideCursor()
	fmt.Print("\033[H\033[2J")
	for {
		var wg sync.WaitGroup
		width, height, err := term.GetSize(0)
		if err != nil {
			log.Fatal("Couldn't get size of terminal")
		}
		termBuff := internal.NewBuffer(width, height)
		if err != nil {
			log.Fatal("Couldn't compute terminal size")
		}
		internal.RenderGrid(termBuff, width, height)
		buff := make([]byte, 128)
		ch := make(chan int)
		mem := make(chan []int)
		wg.Add(2)
		wg.Go(func() {
			internal.ReadTemp(1, buff, ch)
			wg.Done()
		})
		wg.Go(func() {
			memAvail, memTot := internal.ReadMem(mem)
			_, _ = memAvail, memTot
			wg.Done()
		})
		temp := <-ch
		mems := <-mem
		_, _ = temp, mems
		for i := range height {
			fmt.Printf(string(termBuff.Get(i)))
		}
		wg.Wait()
		time.Sleep(1 * time.Second)

	}
}
