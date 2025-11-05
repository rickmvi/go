package dialog

import (
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/console/scan"
	"github.com/rickmvi/go/pkg/util/function"
	"strings"
)

// Prompt represents a structure that holds prompt data, including lines and input, stored as lists of any type.
type Prompt struct {
	lines *array.List[any]
	input *array.List[any]
	err   error
}

// New initializes and returns a new instance of the Prompt struct with empty lines and input buffers.
func New() *Prompt {
	return &Prompt{lines: array.New[any](), input: array.New[any]()}
}

// Builder initializes and returns a new Prompt object with specified content and an empty input list.
func Builder(content any) *Prompt {
	return &Prompt{lines: array.Of[any](content), input: array.New[any]()}
}

// InsertLine adds a single line of any type to the prompt output buffer and returns the updated Prompt object.
func (p *Prompt) InsertLine(line any) *Prompt {
	p.lines.Add(line)
	return p
}

// InsertLines appends multiple lines of arbitrary content to the prompt output buffer and returns the modified Prompt object.
func (p *Prompt) InsertLines(lines ...any) *Prompt {
	if p.len() == 0 || len(lines) == 0 {
		return p
	}
	p.lines.AddAll(lines...)
	return p
}

// InsertTitle adds a title with separators made of a repeated character to the prompt output and returns the updated Prompt.
func (p *Prompt) InsertTitle(title string, char string, count int) *Prompt {
	p.Separator(char, count)
	p.InsertLine(title)
	p.Separator(char, count)
	return p
}

// Formatted adds a formatted string to the prompt output buffer and returns the modified Prompt object.
func (p *Prompt) Formatted(format string, args ...any) *Prompt {
	p.lines.Add(fmt.Sprintf(format, args...))
	return p
}

// Separator adds a repeated character sequence to the prompt output buffer and returns the modified Prompt object.
func (p *Prompt) Separator(char string, count int) *Prompt {
	if count <= 0 {
		return p
	}

	separator := strings.Repeat(char, count)
	p.lines.Add(separator)
	return p
}

// GetInput retrieves an input value at the specified index, ensuring it is of type string, or returns an error if not.
func (p *Prompt) GetInput(index int) (string, error) {
	res, err := p.input.GetSafe(index)
	if err != nil {
		return "", err
	}

	str, ok := res.(string)
	if !ok {
		return "", fmt.Errorf("input at index %d is not a string, got %T", index, res)
	}
	return str, nil
}

// --- Input Methods ---

// Input prompts the user with a message, reads a string input, stores it in the `Prompt`'s input list, and returns the Prompt.
func (p *Prompt) Input(message string) *Prompt {
	res, err := scan.Line(message)

	if err != nil {
		return p.recordError(err)
	}
	p.input.Add(res)
	return p
}

// InputInt prompts the user with a message, reads an integer input, stores it in the `Prompt`'s input, and returns the Prompt.
func (p *Prompt) InputInt(message string) *Prompt {
	res, err := scan.Int[int](message)

	if err != nil {
		return p.recordError(err)
	}
	p.input.Add(res)
	return p
}

// InputFloat prompts the user for a floating-point number input, stores the result, and returns the Prompt object.
func (p *Prompt) InputFloat(message string) *Prompt {
	res, err := scan.Float[float64](message)

	if err != nil {
		return p.recordError(err)
	}
	p.input.Add(res)
	return p
}

// InputUntil repeatedly prompts the user with a message until input satisfies the given condition and returns the Prompt.
// If input is invalid, retryMessage is displayed. If condition is unmet, feedbackMessage is displayed.
func (p *Prompt) InputUntil(message, retryMessage, feedbackMessage string, condition function.Predicate[string]) *Prompt {
	for {
		res, err := scan.Line(message)

		if err != nil {
			fmt.Println(retryMessage)
			continue
		}

		if condition(res) {
			p.input.Add(res)
			return p
		}

		fmt.Println(feedbackMessage)
	}
}

// InputIntUntil repeatedly prompts the user with a message until a valid integer satisfying the given condition is entered.
// If input is invalid, retryMessage is displayed. If the condition is unmet, feedbackMessage is displayed.
func (p *Prompt) InputIntUntil(message, retryMessage, feedbackMessage string, condition function.Predicate[int]) *Prompt {
	for {
		res, err := scan.Int[int](message)

		if err != nil {
			fmt.Println(retryMessage)
			continue
		}

		if condition(res) {
			p.input.Add(res)
			return p
		}

		fmt.Println(feedbackMessage)
	}
}

// InputFloatUntil repeatedly prompts for a floating-point value until the input satisfies the given condition and returns the Prompt.
// If input is invalid, retryMessage is displayed. If the condition is unmet, feedbackMessage is displayed.
func (p *Prompt) InputFloatUntil(message, retryMessage, feedbackMessage string, condition function.Predicate[float64]) *Prompt {
	for {
		res, err := scan.Float[float64](message)

		if err != nil {
			fmt.Println(retryMessage)
			continue
		}

		if condition(res) {
			p.input.Add(res)
			return p
		}

		fmt.Println(feedbackMessage)
	}
}

// Format formats a string using stored input values at the specified indices and adds it to the prompt output buffer.
func (p *Prompt) Format(format string, params ...int) *Prompt {
	var args []any

	for _, index := range params {
		if index < 0 || index >= p.lenInput() {
			p.err = fmt.Errorf("Invalid input index %d provided to Format. Available inputs: 0 to %d.", index, p.lenInput()-1)
			return p
		}

		res, err := p.input.GetSafe(index)
		if err != nil {
			p.err = fmt.Errorf("error retrieving input at index %d: %w", index, err)
			return p
		}
		args = append(args, res)
	}

	formattedLine := fmt.Sprintf(format, args...)

	p.lines.Add(formattedLine)
	return p
}

// --- Final Methods ---

// Render outputs all lines stored in the prompt, clears the state, and ensures no operation occurs if no lines exist.
func (p *Prompt) Render() {
	for _, line := range p.lines.ToSlice() {
		if p.len() == 0 {
			return
		}
		fmt.Println(line)
	}
	p.clear()
}

// Line prints the content of a specific line at the given index and clears the prompt state.
func (p *Prompt) Line(index int) {
	if p.len() == 0 {
		return
	}

	fmt.Println(p.GetInput(index))
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

// Do executes the provided Runnable action and then clears the prompt state.
func (p *Prompt) Do(action function.Runnable) {
	action()
	p.clear()
}

// IterateInputs executes the given Consumer on all collected inputs. This is a final method.
func (p *Prompt) IterateInputs(consumer function.Consumer[any]) {
	for _, input := range p.input.ToSlice() {
		consumer(input)
	}
	p.clear()
}

// Clear removes all stored lines and inputs from the prompt buffer, resetting its state.
func (p *Prompt) Clear() {
	p.clear()
}

// Error retrieves the current error stored in the Prompt object, if any, and returns it.
func (p *Prompt) Error() error {
	return p.err
}

// --- Utils Methods (Internal) ---

// len returns the total number of lines currently stored in the prompt output buffer.
func (p *Prompt) len() int {
	return p.lines.Len()
}

// lenInput returns the total number of elements currently stored in the input buffer.
func (p *Prompt) lenInput() int {
	return p.input.Len()
}

// clear resets the prompt and input state.
func (p *Prompt) clear() {
	p.lines.Clear()
	p.input.Clear()
	p.err = nil
}

// recordError stores an error message in the `Prompt` object and returns the updated `Prompt` instance.
func (p *Prompt) recordError(err error) *Prompt {
	p.err = fmt.Errorf("dialog input error: %w", err) // Armazena o erro
	return p
}
