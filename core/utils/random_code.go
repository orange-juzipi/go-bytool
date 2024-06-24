package utils

import (
	"math/rand/v2"
)

// GetRandomCode 获取随机数
// num：多少个随机数
func GetRandomCode(num uint) string {

	result := make([]byte, num)
	for i := 0; i < int(num); i++ {
		result[i] = randStr[rand.IntN(len(randStr))]
	}
	return string(result)
}
