package grid

import (
	"fmt"
	"gotop/internal/buffer"
	"strconv"
	"testing"
)

func TestRenderGrid(t *testing.T) {
	tests := [][]int{{50, 20, 30}, {30, 100, 30}}
	test_buffer := buffer.NewBuffer(20, 30)
	for _, tt := range tests {
		t.Run("test case "+strconv.Itoa(tt[0]), func(t *testing.T) {
			fmt.Print(tt)
			t.Error("hi")
		})
		RenderGrid(test_buffer, 20, 30)
	}
}
