package scan

import (
	"fmt"
	"github.com/rickmvi/go/pkg/type"
)

func Any[T _type.Any](message string) (input T, err error) {
	fmt.Println(message)
	_, err = fmt.Scan(&input)
	return input, err
}

func Format[T _type.Any](format string, a ...any) (input T, err error) {
	format = fmt.Sprintf(format, a...)
	fmt.Print(format)
	_, err = fmt.Scan(&input)
	return input, err
}

func String(message string) (input string, err error) {
	fmt.Printf(message)
	_, err = fmt.Scan(&input)
	return input, err
}

func Int[T _type.Integer](message string) (input T, err error) {
	fmt.Printf(message)
	_, err = fmt.Scan(&input)
	return input, err
}

func Float[T _type.Float](message string) (input T, err error) {
	fmt.Printf(message)
	_, err = fmt.Scan(&input)
	return input, err
}
