package bco

// AnyTrue set target to true if b is true
func AnyTrue(target *bool, b bool) {
	if b {
		*target = true
	}
}

// AnyFalse set target to false if b is false
func AnyFalse(target *bool, b bool) {
	if !b {
		*target = false
	}
}
