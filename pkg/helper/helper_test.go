package helper

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	tests := []struct {
		min int
		max int
	}{
		{min: 0, max: 1},
		{min: 2, max: 4},
		{min: 4, max: 7},
		{min: 7, max: 10},
		{min: 11, max: 12},
		{min: 15, max: 18},
		{min: 1563, max: 2354},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RandInt(tt.min, tt.max)
			t.Logf("randInt(%v, %v) = %v", tt.min, tt.max, got)
			if got < tt.min || got > tt.max {
				t.Errorf("randInt(%v, %v) = %v, want %v <= '%v' <= %v", tt.min, tt.max, got, tt.min, got, tt.max)
			}
		})
		// time.Sleep(30 * time.Millisecond)
	}
}

func TestGetRootPath(t *testing.T) {

	// got := GetRootPath()
	// fmt.Println(got)
	// want := "'"
	// if got != want {
	// 	t.Errorf("GetRootPath() = %v, want %v", got, want)
	// }
}
