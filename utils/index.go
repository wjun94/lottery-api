package utils

import (
	"math/rand"
	"strconv"
)

// randInt64 随机数生成
func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// GetID 获取8位数id
func GetID() string {
	return strconv.FormatInt(randInt64(10000000, 100000000), 10)
}
