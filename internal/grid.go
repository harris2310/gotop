package internal

import (
	"fmt"
)

func InitializeGrid(width int, height int) {
	sum := 4
	secondToLast := height - 1
	for sum < height {
		switch sum {
		case 1:
			fmt.Println("-----------")
		case secondToLast:
			fmt.Println("-----------")
		default:
			fmt.Println("\\||//")
		}
		sum += 1
	}
}

func RenderGrid(width int, height int, temp int) {
	sum := 4
	secondToLast := height - 1
	midway := int(height / 2)
	for sum < height {
		switch sum {
		case 1:
			fmt.Println("\r-----------")
		case secondToLast:
			fmt.Println("\r-----------")
		case midway:
			fmt.Println("\r" + string(temp))
		default:
			fmt.Println("\r||//")
		}
		sum += 1
	}
}
