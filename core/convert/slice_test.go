package convert_test

import (
	"github.com/orange-juzipi/go-bytool/core/convert"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	// slice转换为map
	m := convert.SliceToMap([]string{"a", "b", "c"})
	t.Log(m)
}
