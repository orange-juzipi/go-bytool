package convert_test

import (
	"github.com/OrangePeel-2019/go-bytool/core/convert"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	// slice转换为map
	m := convert.SliceToMap([]string{"a", "b", "c"})
	t.Log(m)
}
