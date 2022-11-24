/*
title：切片转换工具类
author: orange
version: 1.0
*/

package maps

import (
	"reflect"
	"sort"
	"strconv"
)

// SortMap 提取map的key放到切片上排序，
func SortMap(m map[string]any) reflect.Value {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 反射获取 m
	refMap := reflect.ValueOf(m)
	// 反射创建map
	refNewMap := reflect.MakeMap(refMap.Type())

	mLen := refMap.Len()
	for i := 0; i < mLen; i++ {
		for _, v := range refMap.MapKeys() {
			if keys[i] == v.String() {
				// 设置map的值
				refNewMap.SetMapIndex(reflect.ValueOf(keys[i]), reflect.ValueOf(m[keys[i]]))
			}
		}
	}
	return refNewMap
}

// SortMapInt int的Key转换为string,再排序，
func SortMapInt(m map[int]any) reflect.Value {
	// int转换为any
	var anyMap = make(map[string]any)
	for k, v := range m {
		anyMap[strconv.Itoa(k)] = v
	}
	return SortMap(anyMap)
}

// SortMapFloat64 float64转换为string，再排序，
func SortMapFloat64(m map[float64]any) reflect.Value {
	// float64转换为any
	var anyMap = make(map[string]any)
	// float64转换为string
	for k, v := range m {
		anyMap[strconv.FormatFloat(k, 'f', -1, 64)] = v
	}
	return SortMap(anyMap)
}
