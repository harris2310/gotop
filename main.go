package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func readTemp(core int, buff []byte) (temp int) {
	file, err := os.OpenFile("/sys/class/thermal/thermal_zone0/hwmon"+strconv.Itoa(core)+"/temp1_input", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	n, err := file.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	fmt.Println("hello sir")
	sum := 1
	for sum < 10 {
		fmt.Println("------")
		sum += 1
	}
	for {

		buff := make([]byte, 128)

		var len int = readTemp(1, buff)
		fmt.Println(string(buff[:len-4]) + "C")
		time.Sleep(1 * time.Second)

	}
	os.Exit(1)
}
