/*
title：增强 Go 的 cmp 包
author: orange
version: 1.0
*/

package cmp

// Ternary 三元运算符
func Ternary[T, U any](cond bool, a T, b U) any {
	if cond {
		return a
	}
	return b
}
