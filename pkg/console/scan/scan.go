package scan

import (
	"bufio"
	"fmt"
	"github.com/rickmvi/go/pkg/util/type"
	"os"
	"strconv"
	"strings"
)

// Line reads user input from the console after displaying a prompt message and trims surrounding whitespace.
// Returns the input string and an error, if any, occurred during input reading.
func Line(message string) (input string, err error) {
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// Formatted formats a message using the provided arguments, displays it as a prompt, and returns the user input and any error.
func Formatted(message string, args ...any) (input string, err error) {
	msg := fmt.Sprintf(message, args...)
	return Line(msg)
}

// Int prompts the user with a given message, reads input, and converts it to the integer type T.
// Returns the parsed integer value of type T or an error if input parsing fails.
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

// IntUntil prompts the user with a message and repeatedly requests an integer input until a valid integer of type T is entered.
// It returns the valid integer value of type T after successful parsing.
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

// FloatUntil repeatedly prompts the user with a message until valid input is provided and parses it as a float of type T.
// Returns the parsed floating-point value of type T.
func FloatUntil[T _type.Float](message string) T {
	for {
		input, err := Float[T](message)
		if err == nil {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}
