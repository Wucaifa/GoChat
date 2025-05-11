package random

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// GetRandomInt 获取len位随机数
func GetRandomInt(len int) int {
	return rand.Intn(9*int(math.Pow(10, float64(len-1)))) + int(math.Pow(10, float64(len-1)))
}

// 拼接当前日期（格式 20240507）与一个随机整数，形成一个唯一字符串。
func GetNowAndLenRandomString(len int) string {
	return time.Now().Format("20060102") + strconv.Itoa(GetRandomInt(len))
}
