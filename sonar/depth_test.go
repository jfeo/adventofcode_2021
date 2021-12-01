package sonar

import (
	"testing"
)

func TestCountDepthIncreases(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
		{[]int{100, 200}, 1},
		{[]int{100, 99}, 0},
		{[]int{100}, 0},
		{[]int{100, 100, 100, 101}, 1},
	}

	for _, c := range cases {
		got := CountDepthIncreases(c.in)
		if got != c.want {
			t.Errorf("CountDepthIncreases(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
