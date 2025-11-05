package checker

import "fmt"

type UnreachableError struct {
	URL string
	Err error
}

func (e *UnreachableError) Error() string {
	return fmt.Sprintf("Unreachable URL %s: %v", e.URL, e.Err)
}

func (e *UnreachableError) Unwrap() error {
	return e.Err
}
