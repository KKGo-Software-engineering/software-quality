package sum

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)

	if got != 3 {
		t.Errorf("sum(1, 2) = %d; want 3", got)
	}
}

func TestSumShouldReturn1WhenInput1and0(t *testing.T) {
	got := sum(1, 0)

	if got != 1 {
		t.Errorf("sum(1, 0) = %d; want 1", got)
	}
}

func TestSumShouldReturnMinus2WhenInputMinus1andMinus1(t *testing.T) {
	got := sum(-1, -1)

	if got != -2 {
		t.Errorf("sum(-1, -1) = %d; want -2", got)
	}
}
