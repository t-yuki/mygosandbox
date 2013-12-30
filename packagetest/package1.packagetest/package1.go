package packagetest

func Pub() {
	// since _impl.go has same package name, `private` can be referenced by symbolic link
	private()
}
