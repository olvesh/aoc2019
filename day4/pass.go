package main

import (
	"bytes"
	"sort"
	"strconv"
)

func PossiblePasses(from, to int) int {
	numPasses := 0
	for i := from; i < to; i = i + 1 {
		if CheckPass(strconv.Itoa(i)) {
			numPasses = numPasses + 1
		}
	}
	return numPasses
}

func CheckPass(pass string) bool {
	bytePass := []byte(pass)
	return CriteriaSorted(bytePass) && CriteriaDoubleDigit(bytePass)
}

func CriteriaSorted(pass []byte) bool {
	intPass := intSlice(pass)
	return sort.IsSorted(intPass)

}

func intSlice(pass []byte) sort.IntSlice {
	var intPass sort.IntSlice = make([]int, len(pass))
	for i, b := range pass {
		intPass[i] = int(b)
	}
	return intPass
}

func CriteriaDoubleDigit(pass []byte) bool {
	//slice := intSlice(pass)
	//if ! sort.IsSorted(slice) {
	//  return false
	//}
	//
	//sort.IndexSearch()
	match := false
	for i := 0; i < len(pass)-1; i++ {
		idx := bytes.LastIndexByte(pass, pass[i])

		if idx == 1+i {
			//if (idx - i ) % 2 != 0 {
			match = true
			//  i = idx
			//}  else {
			//
			//}
		}
		i = idx
	}
	return match
}

func CriteriaDoubleDigitPart1(pass []byte) bool {
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {

			return true
		}
	}

	return false
}
