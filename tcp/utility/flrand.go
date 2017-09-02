package utility

import (
	"math/rand"
	"time"
)

//获取随机数目
// @param min 最小数目
// @param max 最大数目
func GenerateRandNum(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := r.Intn(max)
	if result < min {
		return min
	}
	return result
}
