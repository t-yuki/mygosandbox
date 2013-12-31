package cobjs1

import (
	"io"
)

type ModelA interface {
	Prop1() string
	SetProp1(string)
}

type ModelACloser interface {
	ModelA
	io.Closer
}

type ModelB interface {
	Ping(string) (pong string)
}
