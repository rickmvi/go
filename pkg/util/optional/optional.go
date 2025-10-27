package optional

import (
	"fmt"
	"github.com/rickmvi/go/pkg/util/function"
	_type "github.com/rickmvi/go/pkg/util/type"
	"reflect"
)

// Optional represents a container object which may or may not contain a non-nil value of the specified type T.
// It provides a way to express the presence or absence of a value without returning nil directly.
type Optional[T _type.Any] struct {
	value     T
	isPresent bool
}

// --- FACTORIES (Constructors) ---

// Of creates a non-empty Optional containing the specified value of type T.
func Of[T _type.Any](value T) *Optional[T] {
	return &Optional[T]{value: value, isPresent: true}
}

// Empty returns a new Optional instance with no value present (isPresent is false by default).
func Empty[T _type.Any]() *Optional[T] {
	return &Optional[T]{}
}

// OfNullable creates an Optional containing the given value if it's non-nil; otherwise, returns an empty Optional.
func OfNullable[T comparable](value T) *Optional[T] {
	// Usar "value == nil" é a forma mais idiomática e performática no Go, mas só funciona
	// se T for restrito a 'comparable' (ou uma interface que suporte 'nil',
	// o que é coberto por 'comparable' quando T é um ponteiro ou tipo interface).
	if value == nil {
		return Empty[T]()
	}
	return Of[T](value)
}

// --- QUERY METHODS (Accessors) ---

// IsPresent returns true if a value is present in the Optional, otherwise returns false.
func (o *Optional[T]) IsPresent() bool {
	return o.isPresent
}

// IsEmpty returns true if the Optional does not contain a value, otherwise returns false.
func (o *Optional[T]) IsEmpty() bool {
	return !o.isPresent
}

// Get retrieves the value stored in the Optional if it is present; otherwise, it panics.
func (o *Optional[T]) Get() T {
	if !o.isPresent {
		panic("Optional is empty (No value present)")
	}
	return o.value
}

// --- ACTION METHODS (Consumers) ---

// IfPresent invokes the provided Consumer with the value if it exists within the Optional. No action is taken if empty.
func (o *Optional[T]) IfPresent(consumer function.Consumer[T]) {
	if o.isPresent {
		consumer(o.value)
	}
}

// IfPresentOrElse executes the action with the Optional's value if present; otherwise, executes the emptyAction.
func (o *Optional[T]) IfPresentOrElse(action function.Consumer[T], emptyAction function.Runnable) {
	if o.isPresent {
		action(o.value)
		return
	}
	emptyAction()
}

// --- TRANSFORMATION METHODS (Mappers) ---

// Filter returns an Optional containing the current value if the predicate evaluates to true; otherwise, it returns empty.
func (o *Optional[T]) Filter(predicate function.Predicate[T]) *Optional[T] {
	if o.isPresent && predicate(o.value) {
		return o
	}
	return Empty[T]()
}

// Map applies the provided mapper function to the value in the current Optional if present, returning a new Optional.
func Map[T, R _type.Any](o *Optional[T], mapper function.Function[T, R]) *Optional[R] {
	if o.isPresent {
		return Of[R](mapper(o.value))
	}
	return Empty[R]()
}

// FlatMap applies the provided mapper function to the value if present, returning the resulting Optional.
// Se T e o resultado do mapper são o mesmo tipo Optional[T], a assinatura pode ser simplificada.
// Se T e o resultado do mapper são tipos diferentes, precisamos de uma função auxiliar:
func FlatMap[T, R _type.Any](o *Optional[T], mapper function.Function[T, *Optional[R]]) *Optional[R] {
	if o.isPresent {
		return mapper(o.value)
	}
	return Empty[R]()
}

// --- ALTERNATIVE VALUE METHODS (Recovery) ---

// Or returns the current Optional if a value is present, otherwise returns an Optional created from the supplied value.
func (o *Optional[T]) Or(supplier function.Supplier[T]) *Optional[T] {
	if o.isPresent {
		return o
	}
	return Of[T](supplier())
}

// OrElse returns the value if present; otherwise, it returns the specified default value.
func (o *Optional[T]) OrElse(value T) T {
	if o.isPresent {
		return o.value
	}
	return value
}

// OrElseGet returns the value if present; otherwise, it invokes the provided Supplier to generate a fallback value.
func (o *Optional[T]) OrElseGet(valueSupplier function.Supplier[T]) T {
	if o.isPresent {
		return o.value
	}
	return valueSupplier()
}

// OrElseThrow returns the value if it is present; otherwise, it invokes the provided Supplier to generate an error.
func (o *Optional[T]) OrElseThrow(exceptionSupplier function.Supplier[error]) (T, error) {
	if o.isPresent {
		// Retorna o valor e nil para erro
		return o.value, nil
	}

	var zero T
	return zero, exceptionSupplier()
}

// OrElsePanic returns the value if present; otherwise, it panics with a formatted message using the provided format and args.
func (o *Optional[T]) OrElsePanic(format string, args ...any) T {
	if o.isPresent {
		return o.value
	}
	panic(fmt.Sprintf(format, args...))
}

// MustGet retrieves the value stored in the Optional if present; otherwise, it panics.
// É um alias para Get() para seguir a convenção MustX do Go para pânico.
func (o *Optional[T]) MustGet() T {
	return o.Get()
}

// Equal compares two Optionals for equality based on their presence status and values, using deep equality for comparisons.
func (o *Optional[T]) Equal(other *Optional[T]) bool {
	if o.isPresent != other.isPresent {
		return false
	}
	if !o.isPresent {
		return true
	}

	return reflect.DeepEqual(o.value, other.value)
}

// String returns a string representation of the Optional, describing whether it contains a value or is empty.
func (o *Optional[T]) String() string {
	if o.isPresent {
		return fmt.Sprintf("Optional[%v]", o.value)
	}
	return "Optional.Empty"
}

// _ is compile-time asserted ensuring that *Optional[int] satisfies the fmt.Stringer interface.
var _ fmt.Stringer = (*Optional[int])(nil)
