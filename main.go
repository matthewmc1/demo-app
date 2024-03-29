package main

import (
	"demo-app/internal"
	"fmt"
)

func main() {

	token, err := internal.CreateToken("johnDoe")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("token %v", token)

	token, err = internal.CreateToken("test")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("token %v", token)

}
