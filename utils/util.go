package utils

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
)

func RemoveDuplicateStrings(strList []string) []string {
	stringMap := make(map[string]bool)
	for _, str := range strList {
		stringMap[str] = true
	}
	newStrList := []string{}
	for key := range stringMap {
		if key != "" {
			newStrList = append(newStrList, key)
		}
	}
	return newStrList
}
func GenerateRandomString(length int) string {
	u1 := uuid.NewV4().String()
	uuidString := strings.Split(u1, "-")[0]
	var seed int64
	for _, c := range uuidString {
		seed += int64(c)
	}
	rand.NewSource(seed)
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
func RandomNum(max, min int) int {
	u1 := uuid.NewV4().String()
	uuidString := strings.Split(u1, "-")[0]
	var seed int64
	for _, c := range uuidString {
		seed += int64(c)
	}
	rand.NewSource(seed)
	// 生成1到100之间的随机整数
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
func GetStringLines(data string) []string {
	if strings.Index(data, "\r\n") != -1 {
		var res []string
		split := strings.Split(data, "\r\n")
		for _, v := range split {
			res = append(res, v+"\r\n")
		}
		return res
	} else {
		var res []string
		split := strings.Split(data, "\n")
		for _, v := range split {
			res = append(res, v+"\n")
		}
		return res
	}
}
