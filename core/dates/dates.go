/*
title：日期时间工具类
author: orange
version: 1.0
*/

package dates

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Now 获取当前时间，格式：yyyy-MM-dd HH:mm:ss
func Now() string {
	t := time.Now()
	return t.Format(time.DateTime)
}

// Format 格式化日期，格式：自定义，符合日期即可
func Format(t time.Time, format string) string {
	return t.Format(format)
}

// FormatDate 格式化日期，格式：yyyy-MM-dd
func FormatDate(t time.Time, format ...string) string {
	return t.Format(time.DateOnly)
}

// FormatTime 格式化时间，格式：HH:mm:ss
func FormatTime(t time.Time) string {
	return t.Format(time.TimeOnly)
}

// FormatDateTime 时间格式化，格式：yyyy-MM-dd HH:mm:ss
func FormatDateTime(t time.Time) string {
	return t.Format(time.DateTime)
}

// Parse 字符串转换为Time
func Parse(value string) (time.Time, error) {
	return time.ParseInLocation(time.DateTime, value, time.Local)
}

// ParseTime 字符串转换为时间戳
func ParseTime(value string) (int64, error) {
	location, err := Parse(value)
	if err != nil {
		return 0, err
	}
	return location.UnixMilli(), nil
}

// ParseDate 字符串转换为纳秒
func ParseDate(value string) (int64, error) {
	location, err := Parse(value)
	if err != nil {
		return 0, err
	}
	return location.UnixNano(), nil
}

// SplitTime 分割时间字符串，返回日期和时间
// 参数：仅限修改日期分隔符，如果不修改，则使用默认分隔符，时间分隔符默认为：":"
// value：时间字符串
// format：自定义格式：yyyy-MM-dd HH:mm:ss | 2022年1月1日 12:00:01
// 例：
// 有时候，部分场景下不需要 "日" ，所以可以在 format 里不填写 "日"，我们是根据 format 格式化：2022年1月1 12:00:01
func SplitTime(value, format string) string {
	// 正则获取自定义格式化分隔符
	reg, _ := regexp.MatchString("([0-9\a-zA-Z]{4}([\\.\\-/|年月\\s]{1,3}[0-9\a-zA-Z]{1,2}){2}日?(\\s?\\d{2}:\\d{2}(:\\d{2})?)?)|(\\d{1,2}\\s?(分钟|小时|天)前)", format)
	if !reg {
		return value
	}

	// 得到分隔符
	separatorValue := MatchDate(value)
	separatorFormat := MatchDate(format)

	// 替换分隔符
	var replaceDateTime string
	// 日期处理
	s := strings.Split(value, separatorValue)
	log.Println("s:", s[0], "----", s[1], s[2][:1], "--", s[2][2:])
	// 日期不足两位，前面补0
	i, _ := strconv.Atoi(s[1])
	i2, _ := strconv.Atoi(s[2][:1])
	if i < 10 && len(s[1]) == 1 {
		s[1] = "0" + s[1]
	}
	if i2 < 10 && len(s[2][:1]) == 1 {
		s[2] = "0" + s[2][:1]
	} else if i2 < 10 && len(s[2][:1]) == 2 {
		s[2] = "0" + s[2][:2]
	}

	// 时间处理
	s2 := strings.Split(value, ":")
	// log.Println("s2:", s2[0][len(s2[0])-2:], "----", s2[1], s2[2])
	// 时间不足两位，前面补0
	if len(s2[0][len(s2[0])-2:]) == 1 {
		s2[0] = "0" + s2[0][len(s2[0])-2:]
	}
	if len(s2[1]) == 1 {
		s2[1] = "0" + s2[1]
	}
	if len(s2[2]) == 1 {
		s2[2] = "0" + s2[2]
	}

	if separatorFormat == "年" || separatorFormat == "月" || separatorFormat == "日" {
		replaceDateTime = s[0] + "年" + s[1] + "月" + s[2]
		if strings.Index(format, "日") != -1 {
			replaceDateTime += "日"
		}
		replaceDateTime += " " + s2[0][len(s2[0])-2:] + ":" + s2[1] + ":" + s2[2]
	} else {
		replaceDateTime = s[0] + separatorFormat + s[1] + separatorFormat + s[2] + " " + s2[0][len(s2[0])-2:] + ":" + s2[1] + ":" + s2[2]
	}

	return replaceDateTime
}

// MatchDate 匹配字符串的日期格式，返回日期格式
func MatchDate(format string) string {
	if len(format) == 0 {
		return ""
	}
	// 声明日期格式匹配
	var datePattern string

	// 声明常见的日期格式，只需到年份即可
	var datePatterns = []string{"-", "/", ".", "年", "月", "日"}
	for i := 0; i < len(datePatterns); i++ {
		if strings.Contains(format, datePatterns[i]) {
			datePattern = datePatterns[i]
			break
		}
	}
	return datePattern
}
