package cobjs1

import (
	"io"
)

// ModelA is a concurrent object that has some properties.
// An implementation may include `Close() error` call, i.e. io.Closer interface.
// If one of methods is called on a closed object, it will cause panic with `its closed` message.
type ModelA interface {
	// Prop1 returns prop1.
	Prop1() string
	// SetProp1 sets prop1 to `val`.
	SetProp1(val string)
}

// ModelAExt1 is extended ModelA with io.Closer interface.
type ModelAExt1 interface {
	ModelA
	io.Closer
}

// ModelAExt2 is extended ModelA with additional property.
type ModelAExt2 interface {
	ModelA
	// Prop2 returns prop2.
	Prop2() int
	// SetProp2 sets prop2 to `val`.
	SetProp2(val int)
}

// ModelB is a concurrent object that has some request processing methods.
// An implementation may include `Close() error` call, i.e. io.Closer interface.
// If one of methods is called on a closed object, it will cause panic with `its closed` message.
type ModelB interface {
	// Ping echoes `val`.
	// This is synchronous method so it will be blocked until echo request is processed.
	Ping(val string) (pong string)

	// PingAsync echoes `val`
	// This is asynchronous method so it will return immediately after echo request is accepted.
	// The result will be sent to `pongCh`.
	PingAsync(val string, pongCh chan<- string)
}
