/*
title：转换工具类
author: orange
version: 1.0
*/

package convert

import (
	"fmt"
	"strings"
)

// NumberToChinese 把数字货币转换为中文货币
func NumberToChinese(num int) string {
	chinese := ""
	chineseMap := []string{"圆整", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseNum := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	var listNum []int
	for ; num > 0; num = num / 10 {
		listNum = append(listNum, num%10)
	}
	n := len(listNum)

	//注意这里是倒序的
	for i := n - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	//注意替换顺序
	for {
		copyChinese := chinese
		copyChinese = strings.Replace(copyChinese, "零万", "万", 1)
		copyChinese = strings.Replace(copyChinese, "零亿", "亿", 1)
		copyChinese = strings.Replace(copyChinese, "零十", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零百", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零千", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零零", "零", 1)
		copyChinese = strings.Replace(copyChinese, "零圆", "圆", 1)

		if copyChinese == chinese {
			break
		} else {
			chinese = copyChinese
		}
	}

	return chinese
}
