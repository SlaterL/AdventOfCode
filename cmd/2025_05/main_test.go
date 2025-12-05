package main

import (
	"testing"
)

func Test_consolidateRanges(t *testing.T) {
	tests := []struct {
		name   string
		ranges [][]int
		want   [][]int
	}{
		{
			name:   "in",
			ranges: [][]int{{2, 4}, {1, 5}},
			want:   [][]int{{1, 5}},
		},
		{
			name:   "out",
			ranges: [][]int{{1, 5}, {2, 4}},
			want:   [][]int{{1, 5}},
		},
		{
			name:   "mid",
			ranges: [][]int{{1, 3}, {6, 8}, {2, 7}},
			want:   [][]int{{1, 8}},
		},
		{
			name:   "mid2",
			ranges: [][]int{{1, 3}, {2, 7}, {6, 8}},
			want:   [][]int{{1, 8}},
		},
		{
			name:   "mid3",
			ranges: [][]int{{2, 7}, {1, 3}, {6, 8}},
			want:   [][]int{{1, 8}},
		},
		{
			name:   "idk",
			ranges: [][]int{{1, 3}, {6, 8}},
			want:   [][]int{{1, 3}, {6, 8}},
		},
		{
			name:   "idk2",
			ranges: [][]int{{1, 3}, {1, 4}},
			want:   [][]int{{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := consolidateRanges(tt.ranges)
			if len(got) != len(tt.want) {
				t.Errorf("consolidateRanges(), want: %v got: %v", tt.want, got)
			}
			for i := range len(got) {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Error()
				}
			}
		})
	}
}
