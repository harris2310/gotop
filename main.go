package main

import (
	"fmt"
	"golang.org/x/term"
	"gotop/internal"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
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

func readMem(ch chan []int) (mem int, total int) {
	file, err := os.OpenFile("/proc/meminfo", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal((err))
	}
	buff := make([]byte, 4096)
	defer file.Close()
	_, errR := file.Read(buff)
	if errR != nil {
		log.Fatal(err)
	}
	memTot := strings.Index(string(buff), "MemTotal:")
	memAvail := strings.Index(string(buff), "MemAvailable")
	re := regexp.MustCompile(`\d+`)
	numTot, errTot := strconv.Atoi(re.FindString(string(buff[memTot:])))
	if errTot != nil {
		log.Fatal(errTot)
	}
	numAvail, errAv := strconv.Atoi(re.FindString(string(buff[memAvail:])))
	if errAv != nil {

		log.Fatal(errAv)
	}
	newSlice := make([]int, 2)
	newSlice[0] = numAvail
	newSlice[1] = numTot
	ch <- newSlice
	return
}

type buffer struct {
	width, height int
	cells         [][]rune
}

type TermBufferer interface {
	resize() interface{}
}

func NewBuffer(w, h int) *buffer {
	cells := make([][]rune, h)
	for y := range cells {
		cells[y] = make([]rune, w)
		for x := range cells[y] {
			cells[y][x] = ' '
		}
	}

	return &buffer{
		width:  w,
		height: h,
		cells:  cells,
	}
}

func main() {
	// oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	// if err != nil {
	// panic(err)
	// }
	// defer term.Restore(int(os.Stdin.Fd()), oldState)

	internal.HideCursor()
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatal("Couldn't get size of terminal")
	}
	fmt.Print("\033[H\033[2J")
	termBuff := NewBuffer(width, height)
	time.Sleep(1 * time.Second)
	for {
		var wg sync.WaitGroup
		width, height, err := term.GetSize(0)
		if err != nil {
			log.Fatal("Couldn't compute terminal size")
		}
		buff := make([]byte, 128)
		ch := make(chan int)
		mem := make(chan []int)
		wg.Add(2)
		wg.Go(func() {
			readTemp(1, buff, ch)
			wg.Done()
		})
		wg.Go(func() {
			memAvail, memTot := readMem(mem)
			fmt.Println(memAvail, memTot)
			wg.Done()
		})
		//	len := <-ch
		mems := <-mem
		fmt.Println(mems[0])
		wg.Wait()
		time.Sleep(1 * time.Second)
	}
}
