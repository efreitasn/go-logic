package adder

// Full performs a full adder on the given inputs.
func Full(a, b, cIn uint8) (cOut, s uint8) {
	abC, abS := Half(a, b)
	abScInC, abScInS := Half(abS, cIn)

	return abC | abScInC, abScInS
}
