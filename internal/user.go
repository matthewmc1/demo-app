package internal

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

func createUsers() map[string]*User {
	users := make(map[string]*User)

	user1 := &User{Username: "johnDoe", Password: "secret123"}
	user2 := &User{Username: "janeSmith", Password: "myPassword"}

	users[user1.Username] = user1
	users[user2.Username] = user2

	return users
}

func GetUser(username string) (*User, error) {
	users := createUsers()

	retrievedUser := users[username]

	// Check if the user exists
	if retrievedUser != nil {
		fmt.Println("User found:", retrievedUser.Username)
		return retrievedUser, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
