package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"securevault/encryption"
	"securevault/auth"
	"strings"
)

func main() {
	fmt.Println("Welcome to SecureVault - Zero-Knowledge Encrypted Storage")

	var key string
	var validKey bool

	// Kullanıcıdan şifreyi almak ve doğrulamak
	for !validKey {
		promptKey := &survey.Input{
			Message: "Enter your encryption key (exactly 16 characters):",
		}
		survey.AskOne(promptKey, &key)

		key = strings.TrimSpace(key)
		if len(key) != 16 {
			fmt.Println("Error: Encryption key must be exactly 16 characters. Please try again.")
		} else {
			validKey = true
		}
	}

	// Kullanıcıdan şifrelenecek veriyi almak
	var data string
	promptData := &survey.Input{
		Message: "Enter the data you want to encrypt:",
	}
	survey.AskOne(promptData, &data)

	// Şifreleme işlemi
	encrypted, err := encryption.EncryptData([]byte(key), data)
	if err != nil {
		fmt.Println("Encryption failed:", err)
		return
	}
	fmt.Println("\nEncrypted Data:", encrypted)

	// Şifre çözme işlemi
	var decryptedData string
	decryptedData, err = encryption.DecryptData([]byte(key), encrypted)
	if err != nil {
		fmt.Println("Decryption failed:", err)
		return
	}

	fmt.Println("\nDecrypted Data:", decryptedData)

	// MFA işlemi
	otp := auth.GenerateOTP()
	fmt.Println("\nGenerated OTP:", otp)

	var userOTP string
	fmt.Print("Enter OTP: ")
	fmt.Scanln(&userOTP)

	if auth.VerifyOTP(userOTP, otp) {
		fmt.Println("OTP Verified. Access granted.")
	} else {
		fmt.Println("Invalid OTP. Access denied.")
	}
}
