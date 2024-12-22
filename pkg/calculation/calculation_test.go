package calculation

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		expectErr  bool
	}{
		{"2+2*2", 6, false},
		{"(2+2)*2", 8, false},
		{"10/2", 5, false},
		{"2+2/0", 0, true}, // Деление на ноль
		{"2+2*", 0, true},  // Невалидное выражение
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		if test.expectErr && err == nil {
			t.Errorf("Expected error for %s", test.expression)
		}
		if !test.expectErr && err != nil {
			t.Errorf("Unexpected error for %s: %v", test.expression, err)
		}
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}
