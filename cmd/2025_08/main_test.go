package main

import (
	"testing"
)

func Test_calcDist(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		p1   *point
		p2   *point
		want float64
	}{
		{
			name: "1d movement",
			p1: &point{
				x: 0,
				y: 0,
				z: 0,
			},
			p2: &point{
				x: 1,
				y: 0,
				z: 0,
			},
			want: 1,
		},
		{
			name: "2d movement",
			p1: &point{
				x: 0,
				y: 0,
				z: 0,
			},
			p2: &point{
				x: 3,
				y: 4,
				z: 0,
			},
			want: 5,
		},
		{
			name: "3d movement",
			p1: &point{
				x: 0,
				y: 0,
				z: 0,
			},
			p2: &point{
				x: 3,
				y: 4,
				z: 12,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcDist(tt.p1, tt.p2)
			if got != tt.want {
				t.Errorf("calcDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slicedelete(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	a[0] = append(a[0], a[1]...)

	a = append(a[:1], a[1+1:]...)

	t.Error(a)
}
