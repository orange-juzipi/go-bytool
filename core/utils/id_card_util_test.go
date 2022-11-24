package utils_test

import (
	"fmt"
	"github.com/OrangePeel-2019/go-bytool/core/utils"
	"testing"
)

func TestGetIdCardByAge(t *testing.T) {
	idCard := "321083197811252119"
	err, i := utils.GetIdCardByAge(idCard)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)

}
