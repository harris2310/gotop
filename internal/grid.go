package internal

import (
	"fmt"
	"strings"
)

func printTemp(buff []byte, len int) {
	fmt.Println("\r" + string(buff[:len-4]) + "C")
}

func printBoxHorizontal(width int) {
	fmt.Printf("%d", width)
	boxSize := int(width/3 - 2)
	boxString := strings.Repeat("─", boxSize)
	boxTopLine := boxString + " " + boxString + " " + boxString
	fmt.Println("\r" + boxTopLine)
}

func printBoxVertical(height, width int) {
	fmt.Printf("%d", height)
	boxSize := int(width/3 - 2)
	lineString := "|" + strings.Repeat(" ", boxSize) + "|" + " " + "|" + strings.Repeat(" ", boxSize) + "|"
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
			printBoxVertical(height, width)

		}
		sum += 1
	}
}
