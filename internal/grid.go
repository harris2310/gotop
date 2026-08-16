package internal

import (
	"fmt"
	"strconv"
	"strings"
)

const memoryPieRadius int8 = 3

const colString string = "|"

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
	fmt.Println("\r" + strconv.FormatFloat(float64(mems[0])/1000000, 'f', 2, 64) + "(avail)/" + strconv.FormatFloat(float64(mems[1])/1000000, 'f', 2, 64) + "(in-use)")
	return
}

func purpleize(target string) string {
	return fmt.Sprintf("\033[95m%s\033[0m", target)
}

func HideCursor() {
	fmt.Printf("\033[?25l")
}

func printBoxHorizontal(runeBuffer *buffer, width int, height int) {
	boxSize := int(width/3 - 2)
	boxString := strings.Repeat("─", boxSize)
	boxTopLine := boxString + " " + boxString + " " + boxString
	runeBuffer.SetRow(height, []rune(boxTopLine))
}

func printBoxVertical(runeBuffer *buffer, colNum int, rowLen int) {
	boxSize := int(rowLen/3 - 2)
	lineString := purpleize(colString) + strings.Repeat(" ", boxSize) + purpleize(colString) + strings.Repeat(" ", boxSize) + purpleize(colString)
	runeBuffer.SetCol(colNum, []rune(lineString), rowLen)
}

func RenderGrid(runeBuffer *buffer, width int, height int) {
	horizontal := '─'
	vertical := '│'
	corner := '┼'

	// Horizontal lines
	rows := []int{
		0,
		height / 2,
		height - 1,
	}

	// Vertical lines
	cols := []int{
		0,
		width / 4,
		(2 * width / 4),
		(3 * width / 4),
		width,
	}

	for _, y := range rows {
		for x := range width {
			runeBuffer.Set(x, y, horizontal)
		}
	}

	for _, x := range cols {
		for y := range height {
			runeBuffer.Set(x, y, vertical)
		}
	}

	// Intersections
	for _, y := range rows {
		for _, x := range cols {
			runeBuffer.Set(x, y, corner)
		}
	}
}
