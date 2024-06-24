package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
	"strings"
)

// GeneratePassword generate password
func GeneratePassword(str, salt string) string {
	m := md5.New()
	m.Write([]byte(str))
	mByte := m.Sum(nil)

	h := hmac.New(sha256.New, []byte(salt))
	h.Write(mByte)

	return hex.EncodeToString(h.Sum(nil))
}

// GenerateHashPassword generate hash password
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash HashPassword hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateRand 生成随机字符串
func GenerateRand(num uint) string {
	var builder strings.Builder
	for i := 0; i < int(num); i++ {
		builder.WriteByte(randStr[rand.IntN(len(randStr))])
	}

	return builder.String()
}
