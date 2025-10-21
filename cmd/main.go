package main

import (
	"fmt"
	"github.com/rickmvi/go/pkg/scan"
)

func main() {

	input, err := scan.Format[int64]("Ola %d\n\n", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(input)
}
