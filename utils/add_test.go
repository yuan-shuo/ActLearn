package utils

import "testing"

func TestMain_Add(t *testing.T) {
	// 测试用例
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"正数相加", 2, 3, 5},
		{"负数相加", -1, -1, -2},
		{"零加正数", 0, 7, 7},
		{"零加零", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
