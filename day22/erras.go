package main

import (
	"errors"
	"fmt"
)

type AppErr struct {
	Code    int
	Message string
}

func (e *AppErr) Error() string {
	return e.Message
}

func Process(fail bool) error {
	if fail == true {
		return fmt.Errorf("Process: %w", &AppErr{Code: 404, Message: "NotFound"})
	}
	return nil
}

func main() {
	err := Process(true)
	var appErr *AppErr
	if errors.As(err, &appErr) {
		fmt.Println("Code:", appErr.Code, "Message:", appErr.Message)
	}
}
