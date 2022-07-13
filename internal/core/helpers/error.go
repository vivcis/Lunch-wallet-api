package helpers

type InActiveUserError struct {
	err string
}

func (i InActiveUserError) Error() string {
	if i.err == "" {
		i.err = "user is inactive"
	}
	return i.err
}
