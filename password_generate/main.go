package main

import ( // libraries used
	"fmt"
	"math/rand"
	"time"
	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*?" // all characters for password generation
const length = 12 // sets min 

func generatePassword() (string) { // generates a password and returns a string
	// time.Now().UnixNano() creates a random number based on the current time in nano seconds
	source := rand.NewSource(time.Now().UnixNano())
	// rand.NewSource
	r := rand.New(source) // // creates a new random number generator using the custom source

	result := make([]byte, length) // allocate a byte slice of the length of password
	for i := range result { // loop until 12 character is reached
		result[i] = charset[r.Intn(len(charset))] // picks a random character and assigns it to given index
	}
	return string(result) // convert the byte slice back into a string
}

func hashPassword(password string) (string, error) { // takes in the password and returns a string and error message
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // hashes the password using bcrypt
	return string(bytes), err
}

func verifyPassword(password, hash string) (bool, error) { // takes password and hash strings
	// checks if hashed password corresponds to plain text
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password));
	if err != nil {
		return false, err
	}
	return true, nil
}

func main() {
	password := generatePassword() // calls generatePassword function which returns a string
	fmt.Println("Generated password:", password) // prints password to console

	hashedPassword, err := hashPassword(password) // gets the hashed password and error message if there was one
	if err != nil { // checks if error message returned contained a message
		fmt.Println("Error hashing password:", err)
		return
	}

	fmt.Println("Hashed password:", hashedPassword) // prints the hashed password to console
	success, err := verifyPassword(password, hashedPassword) // verifys the password by returning two variables, a boolean and error message 

	if !success || err != nil {  // if not successful
		fmt.Println("PASSWORD INCORRECT")
	}
	fmt.Println("PASSWORD CORRECT")
}