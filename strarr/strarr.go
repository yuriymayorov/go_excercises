// Package contains string utilities
package strarr

// #include "strarr.h"
import "C"
import (
	"strconv"
	"bytes"
	"strings"
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

// Func that compress string like sooodaa -> so3da2
// Estimate time: O(n) Estimate required memory: 0(n)
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

// Func that turn matrix N*N
func Rotate(matrix [][]int) [][]int {
	size := len(matrix)
	var result = make([][]int, size)
	
	for i := range matrix {
	    result[i] = make([]int, size)
	    copy(result[i], matrix[i])
	}
	
	for layer := 0; layer < size / 2; layer++{
		//fmt.printf("layer: %v\n", layer)
		first := layer
		last := size - 1 - layer
		//fmt.printf("last: %v\n", last)
		for i := first; i < last; i++ {
			offset := i - first
			//fmt.printf("offset: %v\n", offset/2)
			top := result[first][i]

			// left -> top
			result[first][i] = result[last-offset][first]

			// bottom -> left
			result[last-offset][first] = result[last][last - offset]

			// right -> bottom
			result[last][last - offset] = result[i][last]

			// top -> right
			result[i][last] = top
		}
	}
	return result
}

// Func that set all M-col and all N-row values to 0, if value in point MxN is 0 
func SetZeros(matrix [][]int) [][]int {
	var result = make([][]int, len(matrix))
	
	for i := range matrix {
	    result[i] = make([]int, len(matrix[i]))
	    copy(result[i], matrix[i])
	}

	var row = make([]bool, len(result))
	var column = make([]bool, len(result[0]))

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			if (result[i][j] == 0) {
				row[i] = true
				column[j] = true
			}
		}
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++{
			if (row[i] || column[j]) {
				result[i][j] = 0
			}
		}
	}

	return result
}

// Func check if s2 is cyclic shift of s1
func IsRotation(s1 string, s2 string) bool {
	len1 := len(s1)
	if len1 == len(s2) && len1 > 0 {
		s1s1 := s1 + s1
		return strings.Contains(s1s1, s2)
	}
	return false
}

func MaxSum([]int arr) {
	length := len(arr)
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum := 
	}
}

//--- private ----------------------------------------

func countCompression(str string) int {
	last := str[0]
	size := 0
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == last {
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