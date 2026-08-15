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
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatal("Couldn't get size of terminal")
	}
	fmt.Print("\033[H\033[2J")
	termBuff := internal.NewBuffer(width, height)
	time.Sleep(1 * time.Second)
	for {
		var wg sync.WaitGroup
		// width, height, err := term.GetSize(0)
		if err != nil {
			log.Fatal("Couldn't compute terminal size")
		}
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
			fmt.Println(memAvail, memTot)
			wg.Done()
		})
		temp := <-ch
		mems := <-mem
		fmt.Println(mems, temp)
		for i := range height {
			fmt.Printf(string(termBuff.Get(i)))
		}
		wg.Wait()
		fmt.Println(" ")
		time.Sleep(1 * time.Second)

	}
}
