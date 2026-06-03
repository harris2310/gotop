package internal

import (
	"fmt"
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

func RenderGrid(width int, height int, len int, buff []byte) {
	sum := 4
	secondToLast := height - 1
	midway := int(height / 2)
	fmt.Printf("\033[1;1H")
	for sum < height {
		switch sum {
		case secondToLast:
			fmt.Println("-----------")
		case midway:
			fmt.Println("\r" + string(buff[:len-4]) + "C")
		default:
			fmt.Println("||//")
		}
		sum += 1
	}
}
