package cmp_test

import (
	"github.com/orange-juzipi/go-bytool/core/cmp"
	"testing"
)

func TestTernary(t *testing.T) {
	t.Log(cmp.Ternary(true, "yes", "no"))
	t.Log(cmp.Ternary(false, "yes", 123))
}
