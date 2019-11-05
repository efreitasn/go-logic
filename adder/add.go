package adder

import (
	"strconv"
	"strings"
)

// Add adds two integers.
func Add(a, b int) int {
	carry := a & b
	sum := a ^ b

	for carry > 0 {
		b = carry << 1
		carry = sum & b
		sum ^= b
	}

	return sum
}

// AddFullAdder adds two integers using a full adder.
func AddFullAdder(a, b int) int {
	aBin := strconv.FormatInt(int64(a), 2)
	bBin := strconv.FormatInt(int64(b), 2)
	cIn := uint8(0)
	sumBin := ""

	// Padding
	bBin = padLeft(bBin, "0", len(aBin))
	aBin = padLeft(aBin, "0", len(bBin))

	for i := len(aBin) - 1; i >= 0; i-- {
		aBinChar := string(aBin[i])
		bBinChar := string(bBin[i])
		aBinCharInt, _ := strconv.Atoi(aBinChar)
		bBinCharInt, _ := strconv.Atoi(bBinChar)

		cOut, sum := Full(
			uint8(aBinCharInt),
			uint8(bBinCharInt),
			cIn,
		)
		sumBin = strconv.Itoa(int(sum)) + sumBin
		cIn = cOut
	}

	sumBin = strconv.Itoa(int(cIn)) + sumBin
	sum, _ := strconv.ParseInt(sumBin, 2, 64)

	return int(sum)
}

func padLeft(str, padChar string, length int) string {
	if len(str) >= length {
		return str
	}

	var b strings.Builder
	b.Grow(length - len(str))

	for b.Len() < b.Cap() {
		b.WriteString(padChar)
	}

	b.WriteString(str)

	return b.String()
}
