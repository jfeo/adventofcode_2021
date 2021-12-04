package submarine

import (
	"testing"
)

func TestSimplePilotMany(t *testing.T) {
	cases := []struct {
		in   []Command
		want Submarine
	}{
		{
			in:   []Command{{Type: "forward", X: 1}},
			want: Submarine{Depth: 0, Horizontal: 1},
		},
		{
			in:   []Command{{Type: "down", X: 5}},
			want: Submarine{Depth: 5, Horizontal: 0},
		},
		{
			in:   []Command{{Type: "down", X: 5}, {Type: "up", X: 3}},
			want: Submarine{Depth: 2, Horizontal: 0},
		},
	}

	for _, c := range cases {
		submarine := Submarine{}
		submarine.SimplePilotMany(c.in)
		if submarine != c.want {
			t.Errorf("submarine.SimplePilotMany(%v) => %v, want %v", c.in, submarine, c.want)
		}
	}
}

func TestPilotMany(t *testing.T) {
	cases := []struct {
		in   []Command
		want Submarine
	}{
		{
			in:   []Command{{Type: "forward", X: 1}},
			want: Submarine{Depth: 0, Horizontal: 1, Aim: 0},
		},
		{
			in:   []Command{{Type: "down", X: 5}},
			want: Submarine{Depth: 0, Horizontal: 0, Aim: 5},
		},
		{
			in:   []Command{{Type: "down", X: 5}, {Type: "up", X: 3}},
			want: Submarine{Depth: 0, Horizontal: 0, Aim: 2},
		},
		{
			in:   []Command{{Type: "forward", X: 3}, {Type: "down", X: 2}, {Type: "forward", X: 3}},
			want: Submarine{Depth: 6, Aim: 2, Horizontal: 6},
		},
	}

	for _, c := range cases {
		submarine := Submarine{}
		submarine.PilotMany(c.in)
		if submarine != c.want {
			t.Errorf("submarine.PilotMany(%v) => %v, want %v", c.in, submarine, c.want)
		}
	}
}
