package utils

import (
	"math/rand"
	"time"
)

// GetRandomCode 获取随机数
// num：多少个随机数
func GetRandomCode(num uint) string {
	const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 该随机函数为 go1.20 推荐此写法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, num)
	for i := 0; i < int(num); i++ {
		result[i] = str[r.Intn(len(str))]
	}
	return string(result)
}
