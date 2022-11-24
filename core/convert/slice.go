/*
title：切片转换工具类
author: orange
version: 1.0
*/

package convert

import (
	"log"
	"reflect"
)

// SliceToMap 切片转map
func SliceToMap(slice interface{}) any {
	reflectSlice := reflect.ValueOf(slice)
	log.Println(reflectSlice.Kind())
	log.Println(reflectSlice.Type())
	log.Println(reflectSlice.Len())
	log.Println(reflectSlice.Index(0))

	t := reflect.TypeOf(slice)
	// 获取切片的类型
	sliceType := t.Elem()
	log.Println(sliceType)
	// 获取切片元素的值
	sliceValue := reflect.Indirect(reflectSlice)
	log.Println(sliceValue)

	refMap := make(map[interface{}]interface{}, reflectSlice.Len()+1)

	if reflectSlice.Kind() == reflect.Slice {
		for i := 0; i < reflectSlice.Len(); i++ {
			// refMap = append(refMap, reflectSlice.Index(i))
			// 往map里添加元素
			refMap[reflectSlice.Index(i)] = reflectSlice.Index(i)

			// refMap[reflectSlice.Index(i).Interface().(sliceType)] = reflectSlice.Index(i)
		}
	} else {
		refMap = nil
		return refMap
	}

	return refMap
}
