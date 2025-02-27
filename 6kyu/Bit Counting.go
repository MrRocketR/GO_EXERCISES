package _kyu

import (
	"math/bits"
	"strconv"
	"strings"
)

func CountBits(n uint) int {
	binaryStr := strconv.FormatInt(int64(n), 2)
	count := 0
	for _, v := range binaryStr {
		if v == '1' {
			count++
		}
	}
	return count
}

func CountBits2(n uint) int {
	return bits.OnesCount(n)
}

func CountBits3(n uint) int {
	var res int = 0
	for n > 0 {
		if n&1 == 1 {
			res = res + 1
		}
		n = n >> 1
	}
	return res
}

func CountBits4(n uint) int {
	return strings.Count(strconv.FormatInt(int64(n), 2), "1")
}
