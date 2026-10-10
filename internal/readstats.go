package internal

import (
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func ReadTemp(core int, buff []byte, ch chan int) {
	var file *os.File
	var err error
	if runtime.GOOS == "windows" {

	} else {
		file, err = os.OpenFile("/sys/class/thermal/thermal_zone0/hwmon"+strconv.Itoa(core)+"/temp1_input", os.O_RDONLY, 0644)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	n, err := file.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	ch <- n
}

func ReadMem(ch chan []int) {
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
}
