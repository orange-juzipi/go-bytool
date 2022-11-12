/*
title：脱敏
author: orange
version: 1.0
*/

package desensitization

import (
	"strings"
)

// MobilePhone 手机号
// 必须是数字，且长度为11位
func MobilePhone(phone string) string {
	if len(phone) <= 0 {
		return ""
	}

	return hide(phone, 3, 7)
}

// Password 密码
// 密码全部替换成*
func Password(password string) string {
	if len(password) <= 0 {
		return ""
	}

	return hide(password, 0, len(password))
}

// Email 邮箱
func Email(email string) string {
	if len(email) <= 0 {
		return ""
	}
	index := strings.Index(email, "@")
	if index <= 1 {
		return email
	}

	return hide(email, 1, index)
}

// IdCardNum 身份证号码
// from 从哪里开始
// end 到哪里结束
func IdCardNum(idCardNum string, from, end int) string {
	if len(idCardNum) <= 0 {
		return ""
	}

	return hide(idCardNum, from, end)
}

// Name 姓名
func Name(name string) string {
	if len(name) <= 0 {
		return ""
	}

	return hide(name, 1, len(name))
}

// hide 隐藏方法
// start – 开始位置（包含）
// end – 结束位置（不包含）
func hide(str string, start, end int) string {
	if end < start {
		return str
	}
	r := []rune(str)

	// 如果rune的长度 < 指定的end，那么在for中就循环len(rune)的次数
	// 否则循环次数就是是用end的次数
	// 例如
	//		12345678910
	// 		rune的长度为 7，end为 10
	//		那么在for中打印 * ，只需要7个
	//		1*******910
	// 不做这个判断，直接使用end的话，那么会打印10个 *
	rLen := end
	if len(r) < end {
		rLen = len(r)
	}

	builder := strings.Builder{}

	builder.WriteString(string(r[:start]))
	for i := start; i < rLen; i++ {
		builder.WriteString("*")
	}
	builder.WriteString(str[end:])

	return builder.String()
}
