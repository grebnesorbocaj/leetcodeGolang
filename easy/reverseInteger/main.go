package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(0))
	fmt.Println(reverse(-873))
	fmt.Println(reverse(-128))
	fmt.Println(reverse(1234))
	fmt.Println(reverse(1511))
	fmt.Println(reverse(2147483647))
	fmt.Println(reverse(-2147483648))

}

func reverse(i int) int {
	x := int32(i)
	y := reverseInteger(x)
	return int(y)
}

func reverseInteger(i int32) int32 {
	if i < 0 {
		i = -i
		var si string  // stringedInteger
		var rsi string // stringedInteger
		si = String(i)
		rsi = Reverse(si) // reversedStringInteger
		ri, err := strconv.ParseInt(rsi, 10, 32)
		if err != nil {
			return int32(0)
		}
		return int32(-ri)
	} else {
		var si string  // stringedInteger
		var rsi string // stringedInteger
		si = String(i)
		rsi = Reverse(si) // reversedStringInteger
		ri, err := strconv.ParseInt(rsi, 10, 32)
		if err != nil {
			return int32(0)
		}
		return int32(ri)
	}
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
