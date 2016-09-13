package errors

import (
	"encoding/json"
	"fmt"
	"os"
)

// BadRequestError represents a bad request error with its error message
type BadRequestError struct {
	Error string `json:"error"`
}

// WrongUsage writes a message saying to the user that he's doing a wrong usage
// of the current command (wrong number of args, ...)
func WrongUsage() {
	fmt.Println("Wrong command usage, please use --help")
	os.Exit(-1)
}

// Error displays an unknown error to the user
func Error(err error) {
	fmt.Printf("An error occured: %v", err)
	os.Exit(-1)
}

// PrintBadRequestError parses this error sent by Admiral and print it
func PrintBadRequestError(data []byte) {
	var badReqErr BadRequestError
	err := json.Unmarshal(data, &badReqErr)
	if err != nil {
		panic(err)
	}

	fmt.Println(badReqErr.Error)
	os.Exit(-1)
}

// Unauthorized prints a message to the user to let him knows that he has
// insufficient rights to access to needed resource
func Unauthorized() {
	fmt.Println("You do not have the required rights to access this resource.")
	os.Exit(-1)
}
