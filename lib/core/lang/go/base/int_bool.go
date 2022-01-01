package base

const (
	IntBoolTrue IntBool = 1
)

type IntBool int8

func (i IntBool) SwitchInt8() int8 {
	return int8(i)
}

func (i IntBool) SwitchBool() bool {
	switch i {
	case IntBoolTrue:
		return true
	default:
		return false
	}
}
