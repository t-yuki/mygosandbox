package goenumtest

type Enum1 uint8

const (
	Enum1A Enum1 = iota
	Enum1B
)

type Enum2 struct {
	val uint8
}

// Enum2{iota} is not constant initializer. so use var.
var (
	Enum2A = Enum2{0}
	Enum2B = Enum2{1}
)

type Enum3 struct {
	uint8
}

func Enum3A() Enum3 {
	return Enum3{0}
}

func Enum3B() Enum3 {
	return Enum3{1}
}

type Enum4 struct {
	val uint8
}

func Enum4A() Enum4 {
	return Enum4{0}
}

func Enum4B() Enum4 {
	return Enum4{1}
}
