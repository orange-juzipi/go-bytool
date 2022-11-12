package maps_test

import (
	"go-bytool/core/maps"
	"testing"
)

func TestA(t *testing.T) {
	// map
	m := make(map[string]interface{})
	m["c"] = 3
	m["a"] = 1
	m["b"] = 2
	// t.Log(m)
	m2 := maps.SortMap(m)
	t.Log(m2)
	// for k, v := range m {
	// 	log.Println(k, "----", v)
	// }

	// v := convs.SliceToMap(m)
	// t.Log("----------", v)
}
