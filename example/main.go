package main

import (
	"errors"
	"fmt"

	"github.com/4CHILL3S101/handleErr"
)

func main() {
	err := errors.New("database connection failed")

	if isErr, msg := handleErr.CheckErr(err, "Unable to load user data", true); isErr {
		fmt.Println(msg)
		return
	}

	println("Success!")
}
