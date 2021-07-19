package wphash

import (
	"crypto/md5"
	"math/rand"
	"strings"
	"time"
)

var itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func subStr(str string, start int, len int) string {
	return str[start:(start + len)]
}

func encode64(input string, count int) string {
	output := ""
	i := 0
	for {
		var value = rune(input[i])
		i++
		output += string(itoa64[value&0x3f])
		if i < count {
			value |= rune(input[i]) << 8
		}
		output += string(itoa64[(value>>6)&0x3f])
		if i >= count {
			break
		}
		i++
		if i < count {
			value |= rune(input[i]) << 16
		}
		output += string(itoa64[(value>>12)&0x3f])
		if i >= count {
			break
		}
		i++
		output += string(itoa64[(value>>18)&0x3f])
		if i >= count {
			break
		}
	}

	return output
}

func cryptPrivate(password string, setting string) string {
	output := "*0"
	if subStr(setting, 0, 2) == output {
		output += "*1"
	}

	if subStr(setting, 0, 3) != "$P$" {
		return output
	}

	var countLog = strings.Index(itoa64, string(setting[3]))
	if countLog < 7 || countLog > 30 {
		return output
	}

	var count = 1 << countLog

	var salt = subStr(setting, 4, 8)
	if len(salt) != 8 {
		return output
	}

	var hash = md5.Sum([]byte(salt + "" + password))
	for {
		hash = md5.Sum(append(hash[0:], []byte(password)...))
		count--
		if count <= 0 {
			break
		}
	}

	output = subStr(setting, 0, 12)
	output += encode64(string(hash[0:]), 16)
	return output
}

func gensalt(input string) string {
	output := "$P$"
	output += string(itoa64[13])
	output += encode64(input, 6)
	return output
}

func genRandomStr(length int) string {
	seed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	total := len(seed)
	var str []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		char := seed[rand.Intn(total)]
		str = append(str, char)
	}
	return string(str)
}

// Verify user submmited password and db password hash
func CheckWordPressPasswordHash(password string, hash string) bool {
	repeat := cryptPrivate(password, hash)
	return hash == repeat
}

// Hash password
func HashPassword(password string) string {
	salt := gensalt(genRandomStr(6))
	return cryptPrivate(password, salt)
}
