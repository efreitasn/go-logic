package adder

// Add adds two ints.
func Add(a, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}

	return a ^ b
}
