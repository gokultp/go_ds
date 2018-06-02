package arrays_test

import (
	"testing"

	"github.com/gokultp/ds/arrays"
)

func TestMap(t *testing.T) {
	a := arrays.Array{1, 2, 3, 4}
	b := a.Map(func(value, index int) interface{} {
		return value * 5
	})
	t.Error(b[2])

}
