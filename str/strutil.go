// Package contains string utilities
package strutil

// #include "strutil.h"
import "C"
import (
	"strconv"
	"bytes"
)

// Func check if char occurs once in string
// Estimate time: O(n) Estimate required memory: O(1)
func IsUniqueChars(str string) bool {
	if len(str) > 256 {
		return false
	}

	var charSet [256]bool
	for i := 0; i < len(str); i++ {
		val := int(str[i])
		if charSet[val] {
			return false
		}
		charSet[val] = true
	}

	return true
}

// Func that reverse string, which ending null
func Reverse(str string) string {
	ch := C.CString(str)
	C.reverse(ch)
	res := C.GoString(ch)
	return res
}

// Func that check if strings are permutations
// Estimate time: O(n) Estimate required memory: O(1)
func IsPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var letters [256]int
	for _, c := range s1 {
		letters[c]++;
	}

	for i := 0; i < len(s2); i++ {
		letters[s2[i]]--
		if letters[s2[i]] < 0 {
			return false
		}
	}

	return true
}

func Compress(str string) string {
	size := countCompression(str)
	if (size >= len(str)) {
		return str
	}

	var buffer bytes.Buffer

	last := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == last{
			count++
		} else {
			buffer.WriteString(string(last))
			buffer.WriteString(strconv.Itoa(count))
			last = str[i]
			count = 1
		}
	}	

	buffer.WriteString(string(last))
	buffer.WriteString(strconv.Itoa(count))
	return buffer.String()
}

func countCompression(str string) int {
	last := str[0]
	size := 0
	count := 1
	for i := 1; i < len(str); i++ {
		if (str[i] == last) {
			count++
		} else {
			last = str[i]
			size += 1 + len(string(count))
			count = 0
		}
	}	
	size += 1 + len(string(count))
	return size
}