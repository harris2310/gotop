package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func readTemp(core int, buff []byte, wg *sync.WaitGroup, ch chan int) (temp int) {
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
	fmt.Println("hello sir")
	sum := 1
	for sum < 10 {
		fmt.Println("------")
		sum += 1
	}
	for {
		var wg sync.WaitGroup
		buff := make([]byte, 128)
		len := make(chan int)
		wg.Add(1)
		wg.Go(func() {
			defer wg.Done()
			readTemp(1, buff, &wg, len)
		})

		wg.Wait()
		fmt.Println(string(buff[:<-len-4]) + "C")
		time.Sleep(1 * time.Second)
	}
	os.Exit(1)
}
