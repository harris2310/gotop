package internal

import (
	"fmt"
	"strconv"
	"strings"
)

const colString string = "\u23D0"

const degreeSign string = "\u00B0"

const vertAlLeft string = "\u2e20"

const smirkFace string = "\U0001F60F"

func printTemp(buff []byte, len int) {
	temp, err := strconv.Atoi(string(buff[:len-4]))
	if err != nil {
		fmt.Printf("Couldn't convert to int")
	}
	switch {
	case temp > 60:
		fmt.Println("\r" + strconv.Itoa(temp) + " " + smirkFace)
	default:
		fmt.Println("\r" + strconv.Itoa(temp) + degreeSign)
	}
}

func printMems(mems []int) {
	fmt.Println("\r" + strconv.FormatFloat(float64(mems[0]/1000000), 'g', 2, 64) + " " + strconv.FormatFloat(float64(mems[1]/1000000), 'g', 2, 64))
	return
}

func purpleize(target string) string {
	return fmt.Sprintf("\033[95m%s\033[0m", target)
}

func HideCursor() {
	fmt.Printf("\033[?25l")
}

func printBoxHorizontal(width int) {
	boxSize := int(width/3 - 2)
	boxString := strings.Repeat("─", boxSize)
	boxTopLine := boxString + " " + boxString + " " + boxString
	fmt.Println(purpleize(boxTopLine))
}

func printBoxVertical(width int) {
	boxSize := int(width/3 - 2)
	lineString := purpleize(colString) + strings.Repeat(" ", boxSize) + purpleize(colString) + strings.Repeat(" ", boxSize) + purpleize(colString)
	fmt.Println("\r" + lineString)
}

func RenderGrid(width int, height int, len int, buff []byte, mems []int, isInitial bool) {
	sum := 1
	last := height - 1
	midway := int(height / 2)
	oneBelowMidway := midway - 1
	oneAboveMidway := midway + 1
	fmt.Printf("\033[1;1H")
	for sum < height {
		switch sum {
		case 1, last, oneAboveMidway:
			printBoxHorizontal(width)
		case midway:
			if isInitial {
				printBoxHorizontal(width)
			} else {
				printTemp(buff, len)
			}
		case oneBelowMidway:
			printMems(mems)
		default:
			printBoxVertical(width)

		}
		sum += 1
	}
}
