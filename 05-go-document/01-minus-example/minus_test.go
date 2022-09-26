package sum

import "testing"

func TestMinus(t *testing.T) {
	t.Run("minus 2 and 1", func(t *testing.T) {
		want := 1

		got := Minus(2, 1)

		if got != want {
			t.Errorf("minus(1, 2) = %d; want %d", got, want)
		}
	})

}
