package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "1"
	wrongPassword := "12345678"

	// Generate encrypted password
	fmt.Print("Generate encrypted password...")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	encryptedPassword := string(hash)
	fmt.Println(encryptedPassword)

	// Validate with correct password
	fmt.Print("Validate with correct password...")
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		fmt.Println("password wrong")
	} else {
		fmt.Println("password ok")
	}

	// Validate with wrong password
	fmt.Print("Validate with wrong password...")
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(wrongPassword))
	if err != nil {
		fmt.Println("password wrong")
	} else {
		fmt.Println("password ok")
	}

	// Validate with ruby encrypted password
	fmt.Print("Validate with ruby encrypted password...")
	encryptedPassword = "$2a$10$lYfmAVROQyQTKLTGqqf64.BT8eN8Brc2DW9.n18Z6yl3lOOGrf9kC"
	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		fmt.Println("password wrong")
	} else {
		fmt.Println("password ok")
	}
}
