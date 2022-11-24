package maps_test

import (
	"fmt"
	"github.com/OrangePeel-2019/go-bytool/core/maps"
	"testing"
)

func TestSortMap(t *testing.T) {
	m := make(map[string]interface{})
	m["c"] = 3
	m["a"] = 1
	m["b"] = 2

	m2 := maps.SortMap(m)
	t.Log(m2)
	fmt.Printf("%T----%v\n", m2, m2)
}
