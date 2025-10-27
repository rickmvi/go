package function

import (
	_type "github.com/rickmvi/go/pkg/util/type"
)

// Package function provides functional interfaces (type aliases) and factory functions
// to simplify the use of generics and functional programming patterns in Go,
// inspired by Java's java.util.function package.

// --- Functional Types ---

// Predicate represents a function that accepts one argument and returns a boolean result.
// It is typically used for filtering or condition checking.
type Predicate[T _type.Any] func(T) bool

// BiPredicate represents a function that accepts two arguments of the same type and returns a boolean result.
// It is typically used for comparison logic, like in sorting or matching pairs.
type BiPredicate[T _type.Any] func(T, T) bool

// Consumer represents an operation that accepts a single input argument and returns no result.
// It is typically used for side effects, like printing or logging.
type Consumer[T _type.Any] func(T)

// BiConsumer represents an operation that accepts two input arguments and returns no result.
// It is typically used for side effects involving two values.
type BiConsumer[T _type.Any] func(T, T)

// Function represents a function that accepts one argument of type T and produces a result of type R.
// It is typically used in mapping operations.
type Function[T _type.Any, R _type.Any] func(T) R

// BiFunction represents a function that accepts two arguments of type T and produces a result of type R.
type BiFunction[T _type.Any, R _type.Any] func(T, T) R

// Supplier represents a supplier of results, typically used to lazily generate a value.
type Supplier[T _type.Any] func() T

// BooleanSupplier represents a function that supplies a boolean value, typically used for conditional or state evaluation.
type BooleanSupplier func() bool

// UnaryOperator represents an operation on a single operand that produces a result of the same type T.
// It is used where the input and output types of a Function are the same, common in Map operations.
type UnaryOperator[T _type.Any] func(T) T

// BinaryOperator represents an operation upon two operands of the same type T, producing a result of the same type T.
// It is used in aggregation operations like Reduce.
type BinaryOperator[T _type.Any] func(T, T) T

// Runnable represents a function with no parameters and no return value, intended to be used as a lightweight task or action.
type Runnable func()

// IntBinaryOperator is a specialized BinaryOperator for int inputs and outputs.
type IntBinaryOperator func(int, int) int

// LongBinaryOperator is a specialized BinaryOperator for int64 inputs and outputs.
type LongBinaryOperator func(int64, int64) int64

// DoubleBinaryOperator is a specialized BinaryOperator for float64 inputs and outputs.
type DoubleBinaryOperator func(float64, float64) float64

// --- Core Utility Functions (Factories) ---

// EqualsTo checks if two comparable values are equal.
func EqualsTo[T comparable](a, b T) bool {
	return a == b
}

// NotEqualsTo checks if two comparable values are not equal.
func NotEqualsTo[T comparable](a, b T) bool {
	return a != b
}

// Equals returns a Predicate that tests if the input value is equal to the provided target.
func Equals[T comparable](target T) Predicate[T] {
	return func(value T) bool {
		return target == value
	}
}

// NotEquals returns a Predicate that tests if the input value is not equal to the provided target.
func NotEquals[T comparable](target T) Predicate[T] {
	return func(value T) bool {
		return target != value
	}
}

// And returns a composite Predicate that represents a logical AND of two Predicates.
// The composite Predicate is true only if both input Predicates are true.
func And[T _type.Any](predicate1 Predicate[T], predicate2 Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return predicate1(value) && predicate2(value)
	}
}

// Or returns a composite Predicate that represents a logical OR of two Predicates.
// The composite Predicate is true if at least one of the input Predicates is true.
func Or[T _type.Any](predicate1 Predicate[T], predicate2 Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return predicate1(value) || predicate2(value)
	}
}

// Not returns a Predicate that is the logical negation of the input Predicate.
func Not[T _type.Any](predicate Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return !predicate(value)
	}
}

// Identity returns a UnaryOperator that always returns its input argument.
func Identity[T _type.Any]() UnaryOperator[T] {
	return func(value T) T {
		return value
	}
}

// AndThen returns a composite Function that applies the first function (T->R) and then
// applies the second function (R->R) to the result of the first. Execution order: first -> second.
func AndThen[T _type.Any, R _type.Any](first Function[T, R], second Function[R, R]) Function[T, R] {
	return func(value T) R {
		return second(first(value))
	}
}

// AndThenMap returns a composite Function that applies the first function (T->R) and then
// applies the second function (R->Z) to the result of the first. Execution order: first -> second.
// This is more flexible than AndThen, allowing the final type Z to be different from the intermediate R.
func AndThenMap[T _type.Any, R _type.Any, Z _type.Any](first Function[T, R], second Function[R, Z]) Function[T, Z] {
	return func(value T) Z {
		return second(first(value))
	}
}

// Compose returns a composite Function that applies the first function (R->R) to the result
// of the second function (T->R). Execution order: second -> first.
func Compose[T _type.Any, R _type.Any](first Function[R, R], second Function[T, R]) Function[T, R] {
	return func(value T) R {
		return first(second(value))
	}
}

// --- Numeric Methods Function (Predicate Factories) ---

// GreaterThan returns a Predicate that tests if the input value is greater than the provided target.
func GreaterThan[T _type.Number](target T) Predicate[T] {
	return func(value T) bool {
		return value > target
	}
}

// GreaterThanOrEqual returns a Predicate that tests if the input value is greater than or equal to the provided target.
func GreaterThanOrEqual[T _type.Number](target T) Predicate[T] {
	return func(value T) bool {
		return value >= target
	}
}

// LessThan returns a Predicate that tests if the input value is less than the provided target.
func LessThan[T _type.Number](target T) Predicate[T] {
	return func(value T) bool {
		return value < target
	}
}

// LessThanOrEqual returns a Predicate that tests if the input value is less than or equal to the provided target.
func LessThanOrEqual[T _type.Number](target T) Predicate[T] {
	return func(value T) bool {
		return value <= target
	}
}

// Sum returns a BinaryOperator that computes the sum of two operands (x + y).
func Sum[T _type.Number]() BinaryOperator[T] {
	return func(x, y T) T {
		return x + y
	}
}

// Subtract returns a BinaryOperator that computes the difference of two operands (x - y).
func Subtract[T _type.Number]() BinaryOperator[T] {
	return func(x, y T) T {
		return x - y
	}
}

// Multiply returns a BinaryOperator that computes the product of two operands (x * y).
func Multiply[T _type.Number]() BinaryOperator[T] {
	return func(x, y T) T {
		return x * y
	}
}

// Divide returns a BinaryOperator that computes the quotient of two operands (x / y).
func Divide[T _type.Number]() BinaryOperator[T] {
	return func(x, y T) T {
		return x / y
	}
}

// Modulo returns a BinaryOperator that computes the remainder of the division of two operands (x % y).
// Requires the type T to be an Integer.
func Modulo[T _type.Integer]() BinaryOperator[T] {
	return func(x, y T) T {
		return x % y
	}
}

// IsNegative returns a Predicate that tests if the input value is less than zero (value < 0).
func IsNegative[T _type.Number]() Predicate[T] {
	return func(value T) bool {
		return value < 0
	}
}

// IsPositive returns a Predicate that tests if the input value is greater than zero (value > 0).
func IsPositive[T _type.Number]() Predicate[T] {
	return func(value T) bool {
		return value > 0
	}
}

// IsZero returns a Predicate that tests if the input value is equal to zero (value == 0).
func IsZero[T _type.Number]() Predicate[T] {
	return func(value T) bool {
		return value == 0
	}
}

// IsNotZero returns a Predicate that tests if the input value is not equal to zero (value != 0).
func IsNotZero[T _type.Number]() Predicate[T] {
	return func(value T) bool {
		return value != 0
	}
}

// Min returns a BinaryOperator that returns the smaller of two operands.
func Min[T _type.Number]() BinaryOperator[T] {
	return func(x, y T) T {
		if x < y {
			return x
		}
		return y
	}
}

// Max returns a BinaryOperator that returns the larger of two operands.
func Max[T _type.Number]() BinaryOperator[T] {
	return func(x, y T) T {
		if x > y {
			return x
		}
		return y
	}
}

// Abs returns a UnaryOperator that computes the absolute value of the input.
func Abs[T _type.Number]() UnaryOperator[T] {
	return func(value T) T {
		if value < 0 {
			return -value
		}
		return value
	}
}
