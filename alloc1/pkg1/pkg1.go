package pkg1

type Maxer interface {
	Max1(a, b int) int
	Max2(a, b int) int
	Max3(a, b int) int
	Max4(a, b int) int
	Max5(a, b int) int
}

type MaxerImpl struct {
}

// Max is an exported package function and it calls unexported package function `max`.
func Max(a, b int) int {
	return max(a, b)
}

// max is an unexported package function.
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Max1 is an exported function with receiver and it calls exported packagee function `Max`.
func (*MaxerImpl) Max1(a, b int) int {
	return Max(a, b)
}

// Max2 is an exported function with receiver and it calls unexported package function `max`.
func (*MaxerImpl) Max2(a, b int) int {
	return max(a, b)
}

// max3 is an unexported function.
func (*MaxerImpl) max3(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Max3 is an exported function with receiver and it calls unexported function `max3` with receiver.
func (impl *MaxerImpl) Max3(a, b int) int {
	return impl.max3(a, b)
}

// Max4 is an exported function with receiver.
func (*MaxerImpl) Max4(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Max5 is an exported function with receiver it calls exported function `Max4` with receiver.
func (impl *MaxerImpl) Max5(a, b int) int {
	return impl.Max4(a, b)
}

// NewMaxer returns MaxerImpl by interface.
func NewMaxer() Maxer {
	return &MaxerImpl{}
}

// NewMaxerImpl returns MaxerImpl by struct pointer.
func NewMaxerImpl() *MaxerImpl {
	return &MaxerImpl{}
}
