package scan

import (
	"bufio"
	"fmt"
	"github.com/rickmvi/go/pkg/type"
	"os"
	"strconv"
	"strings"
)

// Line prompts the user with a message, reads input from stdin, trims surrounding whitespace, and returns the input and any error.
func Line(message string) (input string, err error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// Formatted formats a message using the provided arguments and prompts the user for input, returning the input, and any error.
func Formatted(message string, args ...any) (input string, err error) {
	msg := fmt.Sprintf(message, args...)
	return Line(msg)
}

// Int prompts the user with a message, reads input, and converts it to an integer of type T, returning an error if invalid.
func Int[T _type.Integer](message string) (input T, err error) {
	strInput, err := Line(message)

	if err != nil {
		return *new(T), err
	}

	val, err := strconv.ParseInt(strInput, 10, 64)

	if err != nil {
		return *new(T), err
	}

	return T(val), nil

}

// IntUntil repeatedly prompts the user with the given message until a valid integer of type T is entered and returned.
func IntUntil[T _type.Integer](message string) T {
	for {
		input, err := Int[T](message)
		if err == nil {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

// Float prompts the user with a message, reads input as a string, and converts it to a floating-point value of type T.
// Returns the converted value and an error if the input cannot be parsed as a float.
func Float[T _type.Float](message string) (input T, err error) {
	strInput, err := Line(message)

	if err != nil {
		return *new(T), err
	}

	val, err := strconv.ParseFloat(strInput, 64)

	if err != nil {
		return *new(T), err
	}

	return T(val), nil
}

// FloatUntil prompts the user with a message until a valid floating-point value of type T is entered and returns the value.
func FloatUntil[T _type.Float](message string) T {
	for {
		input, err := Float[T](message)
		if err == nil {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}
