package cli

import (
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/console/scan"
	"github.com/rickmvi/go/pkg/util/function"
	"log"
	"strings"
)

type Prompt struct {
	lines *array.List[any]
	input *array.List[any]
}

// New creates a new instance of Prompt with initialized lines and input slices.
func New() *Prompt {
	return &Prompt{lines: array.New[any](), input: array.New[any]()}
}

// Builder creates a new Prompt instance initialized with the provided message as the first line.
func Builder(content any) *Prompt {
	return &Prompt{lines: array.Of[any](content), input: array.New[any]()}
}

// InsertLine adds a single line of arbitrary content to the prompt output buffer.
func (p *Prompt) InsertLine(line any) *Prompt {
	p.lines.Add(line)
	return p
}

// InsertLines adds multiple lines of arbitrary content to the prompt output buffer.
func (p *Prompt) InsertLines(lines ...any) *Prompt {
	if p.len() == 0 || len(lines) == 0 {
		return p
	}
	p.lines.AddAll(lines...)
	return p
}

// Formatted adds a formatted string to the prompt output buffer using fmt.Sprintf.
func (p *Prompt) Formatted(format string, args ...any) *Prompt {
	p.lines.Add(fmt.Sprintf(format, args...))
	return p
}

// LineSeparator repeats a character/string N times and adds it as a new line.
// This is useful for creating visual separators.
func (p *Prompt) LineSeparator(char string, times int) *Prompt {
	if times <= 0 {
		return p
	}

	separator := strings.Repeat(char, times)
	p.lines.Add(separator)
	return p
}

// GetLine returns the content of the line at the specified index, or nil if index is out of bounds.
func (p *Prompt) GetLine(index int) any {
	if index < 0 || index >= p.len() {
		return nil
	}

	res, err := p.lines.GetSafe(index)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return res
}

// --- Input Methods ---

// Input prompts the user for a string input and stores the result.
func (p *Prompt) Input(message string) *Prompt {
	res, err := scan.Line(message)

	if err != nil {
		log.Fatal(err)
	}
	p.input.Add(res)
	return p
}

// InputInt prompts the user for an integer input and stores the result.
func (p *Prompt) InputInt(message string) *Prompt {
	res, err := scan.Int[int](message)

	if err != nil {
		log.Fatal(err)
	}
	p.input.Add(res)
	return p
}

// InputFloat prompts the user for a float64 input and stores the result.
func (p *Prompt) InputFloat(message string) *Prompt {
	res, err := scan.Float[float64](message)

	if err != nil {
		log.Fatal(err)
	}
	p.input.Add(res)
	return p
}

// InputUntil repeatedly prompts the user until the input string
// satisfies the provided Predicate condition.
func (p *Prompt) InputUntil(message string, condition function.Predicate[string]) *Prompt {
	for {
		res, err := scan.Line(message)

		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		if condition(res) {
			p.input.Add(res)
			return p
		}

		fmt.Println("Invalid input. Please meet the required condition.")
	}
}

// InputIntUntil repeatedly prompts the user until a valid integer is entered AND
// the Predicate condition on that integer is satisfied.
func (p *Prompt) InputIntUntil(message string, condition function.Predicate[int]) *Prompt {
	for {
		res, err := scan.Int[int](message)

		if err != nil {
			fmt.Println("Invalid input. Please enter a whole number.")
			continue
		}

		if condition(res) {
			p.input.Add(res)
			return p
		}

		fmt.Println("Invalid number. Please meet the required condition.")
	}
}

// InputFloatUntil works similarly for float64 inputs, accepting a Predicate on float64.
func (p *Prompt) InputFloatUntil(message string, condition function.Predicate[float64]) *Prompt {
	for {
		res, err := scan.Float[float64](message)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid decimal number.")
			continue
		}

		if condition(res) {
			p.input.Add(res)
			return p
		}

		fmt.Println("Invalid value. Please meet the required condition.")
	}
}

// Format retrieves collected inputs by index, formats a string, and adds it as a line.
func (p *Prompt) Format(format string, params ...int) *Prompt {
	var args []any

	for _, index := range params {
		if index < 0 || index >= p.lenInput() {
			log.Fatalf("Invalid input index %d provided to FormattedFromInput. Available inputs: 0 to %d.", index, p.lenInput()-1)
		}

		res, err := p.input.GetSafe(index)
		if err != nil {
			log.Fatal(err)
		}
		args = append(args, res)
	}

	formattedLine := fmt.Sprintf(format, args...)

	p.lines.Add(formattedLine)
	return p
}

// --- Final Methods ---

// DisplayLines executes the final action: prints all lines and then clears the state.
func (p *Prompt) DisplayLines() {
	for _, line := range p.lines.ToSlice() {
		if p.len() == 0 {
			return
		}
		fmt.Println(line)
	}
	p.clear()
}

// DisplayLine prints the line at the specified index from the prompt output buffer if it isn't empty.
func (p *Prompt) DisplayLine(index int) {
	if p.len() == 0 {
		return
	}

	fmt.Println(p.GetLine(index))
	p.clear()
}

// String returns the concatenated string of all lines, separated by newlines.
// It does NOT clear the state, allowing the string to be retrieved multiple times.
func (p *Prompt) String() string {
	if p.len() == 0 {
		return ""
	}

	res := make([]string, p.len())
	for i, line := range p.lines.ToSlice() {
		res[i] = fmt.Sprintf("%v", line)
	}
	return strings.Join(res, "\n")
}

// Do executes the given action. This is a final method and does not return the Prompt object.
func (p *Prompt) Do(action function.Runnable) {
	action()
	p.clear()
}

// ForEachInput executes the given Consumer on all collected inputs. This is a final method.
func (p *Prompt) ForEachInput(consumer function.Consumer[any]) {
	for _, input := range p.input.ToSlice() {
		consumer(input)
	}
	p.clear()
}

// Clear explicitly resets the prompt and input state.
func (p *Prompt) Clear() {
	p.clear()
}

// --- Utils Methods (Internal) ---

func (p *Prompt) len() int {
	return p.lines.Len()
}

func (p *Prompt) lenInput() int {
	return p.input.Len()
}

// clear resets the prompt and input state.
func (p *Prompt) clear() {
	p.lines.Clear()
	p.input.Clear()
}
