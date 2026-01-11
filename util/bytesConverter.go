package util

import (
	"math"
	"slices"
)

func ConvMultiByteToSingleByte(data []byte, endianness bool) int {
	res := 0
	if endianness {
		// big endian
		slices.Reverse(data)
	}
	for dataI, dataByte := range data {
		res += int(dataByte) * int(math.Pow(16, float64(dataI*2)))
	}
	return res
}
