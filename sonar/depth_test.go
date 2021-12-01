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

func equals(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func TestSlidingWindow(t *testing.T) {
	cases := []struct {
		in_measurements []int
		in_size         int
		want            []int
	}{
		{[]int{10, 20, 30, 40}, 1, []int{10, 20, 30, 40}},
		{[]int{10, 20, 30, 40}, 2, []int{30, 50, 70}},
		{[]int{10, 20, 30, 40}, 3, []int{60, 90}},
		{[]int{10, 20, 30, 40}, 4, []int{100}},
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 3, []int{607, 618, 618, 617, 647, 716, 769, 792}},
	}

	for _, c := range cases {
		got := SlidingWindow(c.in_measurements, c.in_size)
		if !equals(got, c.want) {
			t.Errorf("SlidingWindow(%v, %d) == %v, want %v", c.in_measurements, c.in_size, got, c.want)
		}
	}
}
