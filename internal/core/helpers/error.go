package helpers

import "fmt"

type InActiveUserError struct {
	err string
}

func (i InActiveUserError) Error() string {
	if i.err == "" {
		i.err = "user is inactive"
	}
	fmt.Println("erroroooooo")
	return i.err
}
