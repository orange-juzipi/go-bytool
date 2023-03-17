package dates_test

import (
	"fmt"
	"github.com/OrangePeel-2019/go-bytool/core/dates"

	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	fmt.Printf("获取当前时间: %v\n", dates.Now())
	fmt.Printf("日期时间格式化: %v\n", dates.FormatDateTime(time.Now()))

	fmt.Printf("日期格式化: %v\n", dates.FormatDate(time.Now()))
	fmt.Printf("时间格式化: %v\n", dates.FormatTime(time.Now()))

	t2, err := dates.ParseTime(dates.FormatDateTime(time.Now()))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("字符串转时间戳: %v\n", t2)

	fmt.Println("---------------------------------------")
	s := time.Now().Format("2006年01月02 15:04:05")
	fmt.Printf("官方的时间格式化: %v\n", s)

	// 有时候，部分场景下不需要 "日" ，所以可以在 format 里不填写 "日"，我们是根据 format 格式化：2022年1月1 12:00:01
	replaceDateTime := dates.SplitTime("2021-12-31 12:00:01", "2022年1月01 HH:mm:ss")
	//replaceDateTime := dates.SplitTime("2021-12-31 12:00:01", "2022年1月01日 HH:mm:ss")
	//replaceDateTime := dates.SplitTime("2021-01-1 12:00:01", "yyyy.MM.dd HH:mm:ss")
	fmt.Printf("我们的自定义格式化后: %v\n", replaceDateTime)
	fmt.Println("---------------------------------------")
}
