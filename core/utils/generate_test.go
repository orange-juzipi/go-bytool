package utils

import "testing"

const password = "123456"

func TestGeneratePassword(t *testing.T) {
	generatePassword := GeneratePassword(password, SaltPassword)
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
