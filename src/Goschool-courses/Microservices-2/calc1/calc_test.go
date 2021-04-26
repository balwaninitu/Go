package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}

	result = Add(1, 0)

	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1)
	}

	result = Add(2, -2)

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}

}

func TestSubtract(t *testing.T) {
	result := Subtract(1, 1)

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}

	result = Subtract(10, 5)

	if result != 5 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 5)
	}

	result = Subtract(-5, -5)

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(1, 4)

	if result != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 4)
	}

	result = Multiply(5, 5)

	if result != 25 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 25)
	}

	result = Multiply(100, 0)

	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(100, 5)

	if result != 20 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, float64(20))
	}

	result = Divide(100, 1000)

	if result != float64(0.1) {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, float64(0.1))
	}
}

func TestAddTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{1, 0, 1},
		{2, -2, 0},
	}
	for _, tt := range tests {
		//testname := fmt.Sprintf("%d %d", tt.a, tt.b)
		t.Run("Add Table Suite", func(t *testing.T) {
			ans := Add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
