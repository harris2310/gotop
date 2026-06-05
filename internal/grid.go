package internal

import (
	"fmt"
	"strings"
)

func InitializeGrid(width int, height int) {
	sum := 4
	secondToLast := height - 1
	for sum < height {
		switch sum {
		case 4:
			fmt.Println("-----------")
		case secondToLast:
			fmt.Println("-----------")
		default:
			fmt.Println("\\||//")
		}
		sum += 1
	}
}

func printTemp(buff []byte, len int) {
	fmt.Println("\r" + string(buff[:len-4]) + "C")
}

func printBoxTop(width int) {
	fmt.Printf("%d", width)
	boxSize := int(width/3 - 2)
	boxString := strings.Repeat("─", boxSize)
	boxTopLine := boxString + " " + boxString + " " + boxString
	fmt.Println("\r" + boxTopLine)
}

func RenderGrid(width int, height int, len int, buff []byte) {
	sum := 1
	last := height - 1
	midway := int(height / 2)
	oneBelowMidway := midway - 1
	oneAboveMidway := midway + 1
	fmt.Printf("\033[1;1H")
	for sum < height {
		switch sum {
		case 1:
			printBoxTop(width)
		case last:
			printBoxTop(width)
		case midway:
			printTemp(buff, len)
		case oneBelowMidway:
			printBoxTop(width)
		case oneAboveMidway:
			printBoxTop(width)

		default:
			fmt.Println("||//")
		}
		sum += 1
	}
}
