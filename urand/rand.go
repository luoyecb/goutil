// 随机数生成工具
package urand

import (
	// "crypto/rand"
	// "math/big"
	"bytes"
	mathrand "math/rand"
	"time"
)

var (
	letters = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	mathrand.Seed(time.Now().UnixNano())
}

func RandRange(min, max int64) int64 {
	return mathrand.Int63n(max-min+1) + min
}

func RandString(min, max int) string {
	var buf bytes.Buffer

	for i, j := int64(0), RandRange(int64(min), int64(max)); i < j; i++ {
		index := mathrand.Intn(26)
		buf.WriteByte(letters[index])
	}

	return buf.String()
}

/*
func RandRange2(min, max int64) int64 {
	maxInt := big.NewInt(max)
	v, _ := rand.Int(rand.Reader, maxInt)
	for v.Int64() < min {
		v, _ = rand.Int(rand.Reader, maxInt)
	}
	return v.Int64()
}
*/
