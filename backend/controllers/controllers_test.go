package controllers

import (
	"fmt"
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	fmt.Println(generateVerificationCode(100))
}
