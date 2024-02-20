package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	password := generatePassword()
	fmt.Println(password)
}

func generatePassword() string {
	var password strings.Builder

	password.WriteByte(uppercaseLetters[rand.Intn(len(uppercaseLetters))])
	password.WriteByte(uppercaseLetters[rand.Intn(len(uppercaseLetters))])
	password.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])

	password.WriteString("-")

	password.WriteByte(digits[rand.Intn(len(digits))])
	password.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])
	password.WriteByte(digits[rand.Intn(len(digits))])

	password.WriteString("-")

	password.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])
	password.WriteByte(uppercaseLetters[rand.Intn(len(uppercaseLetters))])
	password.WriteByte(uppercaseLetters[rand.Intn(len(uppercaseLetters))])

	password.WriteString("-")

	password.WriteByte(digits[rand.Intn(len(digits))])
	password.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])
	password.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])

	return password.String()
}
