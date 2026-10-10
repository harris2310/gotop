package buffer

type Buffer struct {
	Width, Height int
	cells         [][]rune
}

type TermBufferer interface {
	get() []rune
	resize() any
}

func (buf *Buffer) Get(line int) []rune {
	return buf.cells[line]
}

func (buf *Buffer) SetRow(line int, data []rune) {
	buf.cells[line] = data
}

func (buf *Buffer) SetCol(col int, data []rune, rowLen int) {
	for i := range rowLen {
		buf.cells[i][col] = data[i]
	}
}

func (buf *Buffer) Set(x, y int, r rune) {
	if y < 0 || y >= len(buf.cells) {
		return
	}
	if x < 0 || x >= len(buf.cells[y]) {
		return
	}

	buf.cells[y][x] = r
}

func NewBuffer(w, h int) *Buffer {
	cells := make([][]rune, h)
	for y := range cells {
		cells[y] = make([]rune, w)
		for x := range cells[y] {
			cells[y][x] = ' '
		}
	}

	return &Buffer{
		Width:  w,
		Height: h,
		cells:  cells,
	}
}
