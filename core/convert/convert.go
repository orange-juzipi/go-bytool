/*
title：转换工具类
author: orange
version: 1.0
*/

package convert

import (
	"fmt"
	"math"
	"strings"
)

// NumberToChinese 把数字货币转换为中文货币
func NumberToChinese(number int) string {
	if number == 0 {
		return "零"
	}
	// return NumberChineseFormatter.format(n.doubleValue(), true, true);
	builder := strings.Builder{}
	if number < 0 {
		builder.WriteString("负")
	}
	yuan := math.Round(float64(number * 100))
	fmt.Println(yuan)
	return ""
}
