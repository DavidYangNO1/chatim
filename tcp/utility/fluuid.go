package utility

import (
	"fmt"
	"strconv"
	"time"
)

func GetUUIDStr() string {

	orderNo := strconv.FormatInt(time.Now().Unix(), 10) + fmt.Sprintf("%d", GenerateRandNum(100, 999))

	return orderNo
}

func GenUniqueByPrefix(prefix string) string {

	unique := prefix + strconv.FormatInt(time.Now().UnixNano(), 10) + fmt.Sprintf("%d", GenerateRandNum(1000, 9999))
	fmt.Printf("unique str is %s", unique)
	return unique
}
