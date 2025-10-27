package _type

import "golang.org/x/exp/constraints"

// Any is an empty interface that can hold values of any type, serving as a placeholder for generic implementations.
type Any interface{}

// Number is a type constraint that includes all type satisfying constraints.Integer or constraints.Float.
type Number interface {
	constraints.Integer | constraints.Float
}

// Float is a type constraint interface that restricts type to floating-point numbers (float32 and float64).
type Float interface {
	constraints.Float
}

// Integer is a generic constraint that represents all integer type, including signed and unsigned integers.
type Integer interface {
	constraints.Integer
}

// Ordered is a constraint that includes all type that are either Number or string.
type Ordered interface {
	Number | string
}

type Comparable interface {
	comparable
}
