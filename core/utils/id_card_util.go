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
	"github.com/OrangePeel-2019/go-bytool/pkg"
	"strconv"
	"time"
)

// GetIdCardByAge 获取身份证的年龄
func GetIdCardByAge(idCard string) (int, error) {
	if err := idCardLen(idCard); err != nil {
		return 0, err
	}

	year, _ := strconv.Atoi(idCard[6:10])
	month, _ := strconv.Atoi(idCard[10:12])
	day, _ := strconv.Atoi(idCard[12:14])
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	birthdayTime, err := dates.Parse(date.Format(time.DateTime))
	if err != nil {
		return 0, err
	}

	if birthdayTime.Before(time.Now()) {
		now := time.Now()
		age := now.Year() - birthdayTime.Year()
		if now.Month() < birthdayTime.Month() || now.Day() < birthdayTime.Day() {
			age -= 1
		}
		return age, nil
	}

	return 0, err
}

func idCardLen(idCard string) error {
	if len(idCard) != 18 {
		return errors.New("身份证号码只支持18位数")
	}
	return nil
}
func GetIdCardBirthday(idCard string) (string, error) {
	if err := idCardLen(idCard); err != nil {
		return "", err
	}
	return idCard[10:12] + "月" + idCard[12:14] + "日", nil
}

// GetIdCardProvince 获取省份
func GetIdCardProvince(idCard string) (string, error) {
	if err := idCardLen(idCard); err != nil {
		return "", err
	}
	provinceCode := idCard[0:2]
	code, _ := strconv.Atoi(provinceCode)
	province := pkg.GetProvince(uint(code))

	return province, nil
}
