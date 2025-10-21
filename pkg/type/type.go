package _type

type Integer interface {
	~int | int8 | int16 | int32 | int64
}

type UInteger interface {
	~uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Float interface {
	float32 | ~float64
}

type Any interface {
	Integer | UInteger | Float | string
}
