package convert_test

import (
	"flag"
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	// slice转换为map
	// m := convert.SliceToMap([]string{"a", "b", "c"})
	// t.Log(m)
	//定义值类型的命令行参数变量，xxxVar对其进行初始化。可以为Int，String，Bool
	var postCode int
	flag.IntVar(&postCode, "Postcode", 1234, "Input your post code")
	flag.Parse()
	fmt.Println(postCode)
}
