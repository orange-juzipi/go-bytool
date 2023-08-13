package convert_test

import (
	"github.com/orange-juzipi/go-bytool/core/convert"
	"testing"
)

func TestNumberToChinese(t *testing.T) {
	t.Log(convert.NumberToChinese(999))
	t.Log(convert.NumberToChinese(188732))
}
