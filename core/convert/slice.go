/*
title：切片转换工具类
author: orange
version: 1.0
*/

package convert

import (
	"reflect"
)

// SliceToMap 切片转map
func SliceToMap(slice any) any {
	reflectSlice := reflect.ValueOf(slice)

	refMap := make(map[any]any, reflectSlice.Len()+1)
	if reflectSlice.Kind() == reflect.Slice {
		for i := 0; i < reflectSlice.Len(); i++ {
			// 往map里添加元素
			refMap[reflectSlice.Index(i)] = reflectSlice.Index(i)
		}
	} else {
		refMap = nil
		return refMap
	}

	return refMap
}
