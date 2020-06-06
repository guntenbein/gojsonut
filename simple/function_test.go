package simple

import "testing"

func TestTakePart(t *testing.T) {
	tests := []struct {
		percentage   float64
		total        int64
		expectedPart int64
	}{
		{100, 100, 100},
		{95.5, 100, 96},
		{-10.4, 100, 0},
		{150, 100, 100},
		{0.5, 100, 1},
		{30, 5, 2},
	}
	for _, test := range tests {
		if result := TakePart(test.percentage, test.total); result != test.expectedPart {
			t.Errorf("expected %d but have %d", test.expectedPart, result)
		}
	}
}
