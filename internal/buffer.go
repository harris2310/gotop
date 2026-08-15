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

func (buf *buffer) Set(line int, data []rune) {

}

func NewBuffer(w, h int) *buffer {
	cells := make([][]rune, h)
	for y := range cells {
		cells[y] = make([]rune, w)
		for x := range cells[y] {
			cells[y][x] = 'e'
		}
	}

	return &buffer{
		width:  w,
		height: h,
		cells:  cells,
	}
}
