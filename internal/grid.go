package internal

import (
	"fmt"
	"strings"
)

const colString string = "\u23D0"

const degreeSign string = "\u00B0"

func printTemp(buff []byte, len int) {
	fmt.Println("\r" + string(buff[:len-4]) + degreeSign)
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

func RenderGrid(width int, height int, len int, buff []byte, isInitial bool) {
	sum := 1
	last := height - 1
	midway := int(height / 2)
	oneBelowMidway := midway - 1
	oneAboveMidway := midway + 1
	fmt.Printf("\033[1;1H")
	for sum < height {
		switch sum {
		case 1, last, oneBelowMidway, oneAboveMidway:
			printBoxHorizontal(width)
		case midway:
			if isInitial {
				printBoxHorizontal(width)
			} else {
				printTemp(buff, len)
			}
		default:
			printBoxVertical(width)

		}
		sum += 1
	}
}
