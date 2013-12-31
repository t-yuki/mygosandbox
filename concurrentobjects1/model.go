package cobjs1

// ModelA is a concurrent object that has some properties.
// An implementation may include `Close() error` call, i.e. io.Closer interface.
// If one of methods is called on a closed object, it will cause panic with `its closed` message.
type ModelA interface {
	// Prop1 returns prop1.
	Prop1() string
	// SetProp1 sets prop1 to val.
	SetProp1(string)
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
