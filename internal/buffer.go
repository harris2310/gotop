package internal

type buffer struct {
	width, height int
	cells         [][]rune
}

type TermBufferer interface {
	get() []rune
	resize() any
}

func (buf *buffer) Get(line int) []rune {
	return buf.cells[line]
}

func (buf *buffer) SetRow(line int, data []rune) {
	buf.cells[line] = data
}

func (buf *buffer) SetCol(col int, data []rune, rowLen int) {
	for i := range rowLen {
		buf.cells[i][col] = data[i]
	}
}

func (b *buffer) Set(x, y int, r rune) {
	if y < 0 || y >= len(b.cells) {
		return
	}
	if x < 0 || x >= len(b.cells[y]) {
		return
	}

	b.cells[y][x] = r
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
