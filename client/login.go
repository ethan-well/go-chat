package main

import "fmt"

func login(userID int, password string) error {
	fmt.Printf("userID: %v, password: %v\n", userID, password)
	return nil
}
