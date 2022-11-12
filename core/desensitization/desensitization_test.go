package desensitization_test

import (
	"fmt"
	"github.com/OrangePeel-2019/go-bytool/core/desensitization"
	"testing"
)

func TestMobilePhone(t *testing.T) {
	phone := "15066665666"
	mobilePhone := desensitization.MobilePhone(phone)

	fmt.Println(mobilePhone)
}
func TestPassword(t *testing.T) {
	password := "orange@peel"
	s := desensitization.Password(password)

	fmt.Println(s)
}

func TestEmail(t *testing.T) {
	email := "3xxxxxx@qq.com"
	s := desensitization.Email(email)

	fmt.Println(s)
}

func TestIdCardNum(t *testing.T) {
	idCardNum := "110562255999999999"
	s := desensitization.IdCardNum(idCardNum, 2, 14)

	fmt.Println(s)
}

func TestName(t *testing.T) {
	name := "橘子皮"
	s := desensitization.Name(name)

	fmt.Println(s)
}
