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

// 生成随机字符串作为邮箱用户名
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
