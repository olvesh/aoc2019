package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	ADD            = 1
	MULT           = 2
	INPUT          = 3
	OUTPUT         = 4
	END            = 99

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
		INPUT: func(i int, intops map[int]int) int {
			inputPos := intops[i+1]

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Input: ")
			text, _ := reader.ReadString('\n')
			val, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
			if err != nil {
				log.Fatal(err)
			}

			intops[inputPos] = val
			return i + 2
		},
		OUTPUT: func(i int, intops map[int]int) int {
			outPos := intops[i+1]

			fmt.Println("Output: ", intops[outPos])

			return i + 2
		},

		END: func(i int, intops map[int]int) int { return -1 },
	}
)

type Instruction struct {
	opcode Operation
	paramModes []int
}
//
func NewInstruction(inst int) Instruction {
	instruction := strconv.Itoa(inst)
	if len(instruction) < 3 {
		// Only params
		atoi, _ := strconv.Atoi(instruction)
		return Instruction{opcode: instructions[atoi]}
	} else {
		opcode, _ := strconv.Atoi(instruction[len(instruction)-2:])
		paramModes := instruction[:len(instruction)-2]
		for i:= len(paramModes);

	}

}

type Operation func(i int, intops map[int]int) int
//
//type Instruction struct {
//	Opcode    int
//	NumParams int
//	apply     func(int, map[int]int)
//}

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
