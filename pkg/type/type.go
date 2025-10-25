package _type

// Integer is a constraint that matches all signed integer types including int, int8, int16, int32, and int64.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UInteger is a constraint that matches all unsigned integer types including uint, uint8, uint16, uint32, uint64, and uintptr.
type UInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float is a constraint that matches floating-point types, including float32 and float64.
type Float interface {
	~float32 | ~float64
}

// Number is a constraint that matches either integer or floating-point types.
type Number interface {
	Integer | Float | UInteger
}

// Any is a constraint that matches basic types like integers, unsigned integers, floats, strings, and any other type.
type Any interface {
	Number | ~string | any
}
