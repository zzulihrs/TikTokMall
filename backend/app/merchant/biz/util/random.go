package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 定义一些常用的邮箱域名
var emailDomains = []string{
	"gmail.com",
	"yahoo.com",
	"hotmail.com",
	"outlook.com",
	"163.com",
	"qq.com",
	"sohu.com",
	"sina.com",
}

// 生成随机字符串
func GenerateRandomString(length int) string {
	// 定义可用的字符集
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}

// 生成随机邮箱
func GenerateRandomEmail(usernameLength int) string {
	// 生成随机用户名
	username := GenerateRandomString(usernameLength)
	// 随机选择一个邮箱域名
	domain := emailDomains[rand.Intn(len(emailDomains))]
	return fmt.Sprintf("%s@%s", username, domain)
}

// 生成随机数字
func GenerateRandomNumber(mi, mx int) int {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seededRand.Intn(mx-mi+1) + mi
}

func GenerateRandomInt32(mi, mx int32) int32 {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seededRand.Int31n(mx-mi+1) + mi
}

func GenerateRandomInt64(mi, mx int64) int64 {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seededRand.Int63n(mx-mi+1) + mi
}

func GenerateRandomFloat32(mi, mx float32) float32 {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return seededRand.Float32()*(mx-mi) + mi
}
