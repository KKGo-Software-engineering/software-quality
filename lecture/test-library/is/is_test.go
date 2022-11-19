package is_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func ParseBinary(b string) (bool, error) {
	return true, nil
}

func TestSomething(t *testing.T) {
	is := is.New(t)

	b, err := ParseBinary("0")

	is.NoErr(err)
	is.Equal(b, true)
	is.Equal([]string{"a", "b", "c"}, []string{"a", "b", "c"})

	got := "anuchito is gopher"
	is.True(strings.Contains(got, "anuchito"))
}
