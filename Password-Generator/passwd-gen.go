package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

// Characters to use in the password
const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

func generatePassword(length int, useUppercase, useDigits, useSpecialChars bool) string {
	// Define the character set based on user preferences
	charSet := lowercaseLetters
	if useUppercase {
		charSet += uppercaseLetters
	}
	if useDigits {
		charSet += digits
	}
	if useSpecialChars {
		charSet += specialChars
	}

	// Generate the password
	password := strings.Builder{}
	for i := 0; i < length; i++ {
		// Get a random index from the character set
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		password.WriteByte(charSet[randomIndex.Int64()])
	}

	return password.String()
}

func main() {
	// Prompt the user for password criteria
	var length int
	var useUppercase, useDigits, useSpecialChars bool

	fmt.Print("Enter password length: ")
	fmt.Scanln(&length)

	fmt.Print("Include uppercase letters? (true/false): ")
	fmt.Scanln(&useUppercase)

	fmt.Print("Include digits? (true/false): ")
	fmt.Scanln(&useDigits)

	fmt.Print("Include special characters? (true/false): ")
	fmt.Scanln(&useSpecialChars)

	// Generate the password
	password := generatePassword(length, useUppercase, useDigits, useSpecialChars)

	// Display the generated password
	fmt.Printf("Generated Password: %s\n", password)

	// Wait for user input before closing
	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
