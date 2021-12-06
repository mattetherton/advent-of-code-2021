package day3

import (
	"testing"
)

var input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func Test_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: input,
			want:  198,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDay3(input)
			if got := d.Part1(); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
