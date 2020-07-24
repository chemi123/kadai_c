package main

import (
	"testing"
)

func Test_collectMap(t *testing.T) {
	testCases := []struct {
		name      string
		schedules [][]int
		maxDay    int
		wantDay   int
	}{
		{
			name:      "succeed to collect all the pieces of map",
			schedules: [][]int{[]int{1, 6}, []int{2, 3}, []int{1, 2}, []int{3, 4, 5}},
			maxDay:    6,
			wantDay:   3,
		},
		{
			name:      "fail to collect all the pieces of map",
			schedules: [][]int{[]int{2}, []int{5}},
			maxDay:    5,
			wantDay:   -1,
		},
		{
			name:      "one succesor",
			schedules: [][]int{[]int{2, 5, 10, 30}},
			maxDay:    30,
			wantDay:   0,
		},
	}
	for _, tc := range testCases {
		day := collectMap(tc.schedules, tc.maxDay)
		if day != tc.wantDay {
			t.Fatalf("test fail. want: %v, got: %v\n", tc.wantDay, day)
		}
	}
}

func Test_collectMap2(t *testing.T) {
	testCases := []struct {
		name      string
		schedules [][]int
		maxDay    int
		wantDay   int
	}{
		{
			name:      "succeed to collect all the pieces of map",
			schedules: [][]int{[]int{1, 6}, []int{2, 3}, []int{1, 2}, []int{3, 4, 5}},
			maxDay:    6,
			wantDay:   3,
		},
		{
			name:      "fail to collect all the pieces of map",
			schedules: [][]int{[]int{2}, []int{5}},
			maxDay:    5,
			wantDay:   -1,
		},
		{
			name:      "one succesor",
			schedules: [][]int{[]int{2, 5, 10, 30}},
			maxDay:    30,
			wantDay:   0,
		},
	}
	for _, tc := range testCases {
		day := collectMap2(tc.schedules, tc.maxDay)
		if day != tc.wantDay {
			t.Fatalf("test fail. want: %v, got: %v\n", tc.wantDay, day)
		}
	}

}
