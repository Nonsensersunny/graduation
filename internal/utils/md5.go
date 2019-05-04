package utils

import (
	md52 "crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandSalt() string {
	return randString(6)
}

func MD5Encode(password, salt string) string {
	md5 := md52.New()
	md5.Write([]byte(password))
	md5.Write([]byte(salt))
	return hex.EncodeToString(md5.Sum(nil))
}

func GenAuthCode(width int) string {
	num := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}
	r := len(num)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", num[rand.Intn(r)])
	}
	return sb.String()
}