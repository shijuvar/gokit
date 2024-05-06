package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errEmptyFname = errors.New("first name is empty")
	errEmptyLname = errors.New("last name is empty")
	errAgeMinor   = errors.New("age is less than 18")
)

func getDetails(fname string, lname string, age int) (string, error) {
	if err := validate(fname, lname, age); err != nil {
		return "", err
	}
	return fmt.Sprintf("First Name: %s, Last Name: %s, Age:%d", fname, lname, age), nil
}

func validate(fname string, lname string, age int) error {
	if len(strings.TrimSpace(fname)) == 0 {
		return errEmptyFname
	}
	if len(strings.TrimSpace(lname)) == 0 {
		return errEmptyLname
	}
	if age < 18 {
		return errAgeMinor
	}
	return nil
}

func main() {
	if str, err := getDetails("shiju", "var", 17); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str)
	}
	if str, err := getDetails("shiju", "var", 45); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str)
	}

	// UnWrap error values
	if _, err := getDetailsWrapErr("shiju", "var", 16); err != nil {
		// Is reports whether any error in err's chain matches target.
		if errors.Is(err, errAgeMinor) {
			fmt.Println("Age is less than 18")
		}
		if err := errors.Unwrap(err); err != nil {
			fmt.Println("UnWrap Error:", err)
		}
	}
}

func getDetailsWrapErr(fname string, lname string, age int) (string, error) {
	if err := validate(fname, lname, age); err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}
	return fmt.Sprintf("First Name: %s, Last Name: %s, Age:%d", fname, lname, age), nil
}
