package main

import (
	"reflect"
)

const (
	ADD  = 1
	MULT = 2
	END  = 99
)

var (
	instructions = map[int]Operation{
		ADD: func(i int, intops map[int]int) int {
			resultPos := intops[i+3]
			sum1Pos := intops[i+1]
			sum2Pos := intops[i+2]
			sum := intops[sum1Pos] + intops[sum2Pos]
			intops[resultPos] = sum
			return i + 4
		},
		MULT: func(i int, intops map[int]int) int {
			productPos := intops[i+3]
			coefficient1Pos := intops[i+1]
			coefficient2Pos := intops[i+2]
			product := intops[coefficient1Pos] * intops[coefficient2Pos]
			intops[productPos] = product
			return i + 4
		},
		END: func(i int, intops map[int]int) int { return -1 },
	}
)

type Operation func(i int, intops map[int]int) int

type Instruction struct {
	Opcode    int
	NumParams int
	apply     func(int, map[int]int)
}

var mapify = func(intopsProgram []int) map[int]int {
	m := make(map[int]int, len(intopsProgram))
	for i, val := range intopsProgram {
		m[i] = val
	}
	return m
}

func CalcIntOpsSlice(intops []int) []int {
	rt := reflect.TypeOf(intops)
	switch rt.Kind() {
	case reflect.Array:
		intops = intops[:]
	}

	intopsMap := CalcIntOps(mapify(intops))
	result := make([]int, len(intopsMap))
	for i, i2 := range intopsMap {
		result[i] = i2
	}
	return result
}

func CalcIntOps(intops map[int]int) map[int]int {
	for i := 0; i != -1 && i < len(intops); {
		instruction := intops[i]
		operation := instructions[instruction]
		i = operation(i, intops)

	}
	return intops

}
