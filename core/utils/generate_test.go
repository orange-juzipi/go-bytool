package utils

import "testing"

const (
	password     = "123456"
	saltPassword = "ag7i5hI4sF2HJ3ihs3iioIh09"
)

func TestGeneratePassword(t *testing.T) {
	generatePassword := GeneratePassword(password, saltPassword)
	t.Log(generatePassword)
}
func TestGenerateHashPassword(t *testing.T) {
	generateHashPassword, _ := GenerateHashPassword(password)
	t.Log("generate:", generateHashPassword)

	checkPasswordHash := CheckPasswordHash(password, generateHashPassword)
	t.Log("check:", checkPasswordHash)
}

func TestGenerateRand(t *testing.T) {
	rand := GenerateRand(4)
	t.Log(rand)
}
