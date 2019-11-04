package adder

// Half performs a half adder on the given inputs.
func Half(a, b uint8) (c, s uint8) {
	return a & b, a ^ b
}
