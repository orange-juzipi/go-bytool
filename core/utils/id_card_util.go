/*
title：身份证工具
author: orange
version: 1.0
remark: 暂时只支持18位身份证号码
*/

package utils

import (
	"errors"
	"github.com/OrangePeel-2019/go-bytool/core/dates"
	"time"
)

// GetIdCardByAge 获取身份证的年龄
func GetIdCardByAge(idCard string) (error, int) {
	if err := idCardLen(idCard); err != nil {
		return err, 0
	}

	idCardYear := idCard[6:10]
	idCardMonth := idCard[10:12]
	idCardDay := idCard[12:14]

	birthdayTime, err := dates.ParseDate(idCardYear + "-" + idCardMonth + "-" + idCardDay)
	if err != nil {
		return err, 0
	}

	if birthdayTime.Before(time.Now()) {
		now := time.Now()
		age := now.Year() - birthdayTime.Year()
		if now.Month() < birthdayTime.Month() || now.Day() < birthdayTime.Day() {
			age -= 1
		}
		return nil, age
	}

	return nil, 0
}

func idCardLen(idCard string) error {
	if len(idCard) != 18 {
		return errors.New("身份证号码只支持18位数")
	}
	return nil
}
func getBirthday(idCard string) string {
	return idCard[6:14]
}
