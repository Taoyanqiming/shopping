package utils

import (
	"math/rand"
	"time"
)

// 随机生成 user + n位随机数字
func RandomString(n int) string {
	var letters = []byte("1234567890")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return "user" + string(result)
}
