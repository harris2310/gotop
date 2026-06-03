package main

import (
	"fmt"
	"gotop/internal"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/term"
)

func readTemp(core int, buff []byte, ch chan int) (temp int) {
	file, err := os.OpenFile("/sys/class/thermal/thermal_zone0/hwmon"+strconv.Itoa(core)+"/temp1_input", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	n, err := file.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	ch <- n
	return
}

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatal("Couldn't get size of terminal")
	}
	internal.InitializeGrid(width, height)
	for {
		var wg sync.WaitGroup
		buff := make([]byte, 128)
		ch := make(chan int)
		// oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		// if err != nil {
		// 	panic(err)
		// }
		// defer term.Restore(int(os.Stdin.Fd()), oldState)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		wg.Go(func() {
			readTemp(1, buff, ch)
			wg.Done()
		})
		len := <-ch
		wg.Wait()
		internal.RenderGrid(width, height, len)
		fmt.Print("\r" + string(buff[:len-4]) + "C")
		time.Sleep(1 * time.Second)
	}
	os.Exit(1)
}
