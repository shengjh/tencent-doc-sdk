package util

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	b := make([]byte, (length+1)/2) // 因为hex编码会使长度翻倍
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // 对于随机数生成失败，直接panic是合理的
	}
	return hex.EncodeToString(b)[:length]
}

// GenerateRandomNumber 生成指定范围内的随机数
func GenerateRandomNumber(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
