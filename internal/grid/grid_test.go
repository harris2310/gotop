package grid

import (
	"fmt"
	"gotop/internal/buffer"
	"strconv"
	"testing"
)

func TestRenderGrid(t *testing.T) {
	tests := [][]int{{50, 20}, {30, 100}, {40, 140}}
	test_buffer := buffer.NewBuffer(20, 30)
	for _, tt := range tests {
		t.Run("test case "+strconv.Itoa(tt[0]), func(t *testing.T) {
			fmt.Print(tt)
			_, err := RenderGrid(test_buffer, tt[0], tt[1])
			if err != nil {
				t.Error("Terminal too large")
			}
		})
	}
}
