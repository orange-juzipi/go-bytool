package dates_test

import (
	"fmt"
	"github.com/OrangePeel-2019/go-bytool/core/dates"

	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	fmt.Printf("dates.Now(): %v\n", dates.Now())
	fmt.Printf("dates.Format(time.Now(), \"2006-01-02 15:04:05\"): %v\n", dates.FormatDateTime(time.Now()))

	fmt.Printf("dates.FormatDate(time.Now()): %v\n", dates.FormatDate(time.Now()))
	fmt.Printf("dates.FormatTime(time.Now()): %v\n", dates.FormatTime(time.Now()))

	t2, err := dates.ParseTime(dates.FormatDateTime(time.Now()))
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("dates.ParseTime(): %v\n", t2)
	s := time.Now().Format("2006年01月02 15:04:05")
	fmt.Printf("s: %v\n", s)

	// replaceDateTime := dates.SplitTime("2021-12-31 12:00:01", "2022年1月01日 HH:mm:ss")
	replaceDateTime := dates.SplitTime("2021-01-1 12:00:01", "yyyy.MM.dd HH:mm:ss")
	fmt.Printf("replaceDateTime格式化后: %v\n", replaceDateTime)

	// 声明int64类型的时间戳
	var times = 1648005297371
	t3 := time.Unix(int64(times), 0)
	// fmt.Printf("t3: %v\n", t3)
	fmt.Printf("t3: %v\n", t3.Format("2006-01-02 15:04:05"))

}
