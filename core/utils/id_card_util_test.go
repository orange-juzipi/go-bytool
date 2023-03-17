package utils_test

import (
	"fmt"
	"github.com/OrangePeel-2019/go-bytool/core/utils"
	"testing"
)

var (
	idCard = "321083197811252119"
)

// 获取年龄
func TestGetIdCardByAge(t *testing.T) {
	i, err := utils.GetIdCardByAge(idCard)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(i, "岁")
}

// 获取生日
func TestGetIdCardBirthday(t *testing.T) {
	birthday, err := utils.GetIdCardBirthday(idCard)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(birthday)
}

// 获取省份
func TestGetIdCardProvince(t *testing.T) {
	birthday, err := utils.GetIdCardProvince(idCard)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(birthday)
}
