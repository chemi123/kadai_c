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
	}{}
	for _, tc := range testCases {
		day := collectMap(tc.schedules, tc.maxDay)
		if day != tc.wantDay {
			t.Fatalf("")
		}
	}
}
