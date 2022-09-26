package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return 3 when 1 and 2", func(t *testing.T) {

		got := sum(1, 2)

		if got != 3 {
			t.Errorf("sum(1, 2) = %d; want 3", got)
		}
	})

	t.Run("should return 1 when 1 and 0", func(t *testing.T) {

		got := sum(1, 0)

		if got != 1 {
			t.Errorf("sum(1, 0) = %d; want 1", got)
		}
	})

	t.Run("should return -2 when -1 and -1", func(t *testing.T) {

		got := sum(-1, -1)

		if got != -2 {
			t.Errorf("sum(-1, -1) = %d; want -2", got)
		}
	})
}
