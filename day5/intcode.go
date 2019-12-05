package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	ADD           = 1
	MULT          = 2
	INPUT         = 3
	OUTPUT        = 4
	JUMP_IF_TRUE  = 5
	JUMP_IF_FALSE = 6
	LESS_THAN     = 7
	EQUALS        = 8
	END           = 99
)

var (
	instructions = map[int]Operation{
		ADD: func(i int, intops map[int]int, mode paramMode) int {
			resultPos := intops[i+3]
			sum1Val := mode.valueFor(1, intops[i+1], intops)
			sum2Val := mode.valueFor(2, intops[i+2], intops)

			sum := sum1Val + sum2Val
			intops[resultPos] = sum
			return i + 4
		},
		MULT: func(i int, intops map[int]int, mode paramMode) int {
			productPos := intops[i+3]
			coeff1Val := mode.valueFor(1, intops[i+1], intops)
			coeff2Val := mode.valueFor(2, intops[i+2], intops)

			product := coeff1Val * coeff2Val
			intops[productPos] = product
			return i + 4
		},

		INPUT: func(i int, intops map[int]int, mode paramMode) int {
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
		OUTPUT: func(i int, intops map[int]int, mode paramMode) int {
			outPos := intops[i+1]

			fmt.Println("Output: ", intops[outPos])

			return i + 2
		},
		JUMP_IF_TRUE: func(i int, intops map[int]int, mode paramMode) int {

			testVal := mode.valueFor(1, intops[i+1], intops)
			if testVal != 0 {
				return mode.valueFor(2, intops[i+2], intops)
			}

			return i + 3

		},

		JUMP_IF_FALSE: func(i int, intops map[int]int, mode paramMode) int {

			testVal := mode.valueFor(1, intops[i+1], intops)
			if testVal == 0 {
				return mode.valueFor(2, intops[i+2], intops)
			}

			return i + 3

		},
		LESS_THAN: func(i int, intops map[int]int, mode paramMode) int {

			testVal1 := mode.valueFor(1, intops[i+1], intops)
			testVal2 := mode.valueFor(2, intops[i+2], intops)
			outpos := intops[i+3]

			if testVal1 < testVal2 {
				intops[outpos] = 1
			} else {
				intops[outpos] = 0
			}

			return i + 4

		},

		EQUALS: func(i int, intops map[int]int, mode paramMode) int {

			testVal1 := mode.valueFor(1, intops[i+1], intops)
			testVal2 := mode.valueFor(2, intops[i+2], intops)
			outpos := intops[i+3]
			if testVal1 == testVal2 {
				intops[outpos] = 1
			} else {
				intops[outpos] = 0
			}

			return i + 4

		},

		END: func(i int, intops map[int]int, mode paramMode) int { return -1 },
	}
)

type Instruction struct {
	opcode     Operation
	paramModes []int
}

func NewInstruction(inst int) Instruction {
	newInstruction := Instruction{paramModes: []int{0, 0, 0, 0}}
	stringOpCode := strconv.Itoa(inst)
	if len(stringOpCode) < 3 {
		// Only params
		atoi, _ := strconv.Atoi(stringOpCode)
		newInstruction.opcode = instructions[atoi]

	} else {
		opcode, _ := strconv.Atoi(stringOpCode[len(stringOpCode)-2:])
		newInstruction.opcode = instructions[opcode]
		paramModes := stringOpCode[:len(stringOpCode)-2]

		for i, j := len(paramModes)-1, 0; i >= 0; i, j = i-1, j+1 {
			newInstruction.paramModes[j], _ = strconv.Atoi(paramModes[i : i+1])
		}

	}
	return newInstruction
}

type paramMode []int

func (m paramMode) valueFor(paramIdx int, paramValue int, intops map[int]int) int {
	mode := m.idx(paramIdx)
	if mode == 0 {
		return intops[paramValue]
	} else if mode == 1 {
		return paramValue
	}

	return 0
}
func (m paramMode) idx(i int) int {
	i = i - 1
	if i >= 0 && i < len(m) {
		return m[i]
	}
	return 0
}

type Operation func(i int, intops map[int]int, mode paramMode) int

//type OperationMode func(i int, intops map[int]int) int

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

func CalcIntOpsSlice(intops []int) {
	rt := reflect.TypeOf(intops)
	switch rt.Kind() {
	case reflect.Array:
		intops = intops[:]
	}

	CalcIntOps(mapify(intops))
	//result := make([]int, len(intopsMap))
	//for i, i2 := range intopsMap {
	//  result[i] = i2
	//}
	//return intops
}

func CalcIntOps(intops map[int]int) map[int]int {
	for i := 0; i != -1 && i < len(intops); {
		inst := NewInstruction(intops[i])

		//instruction := intops[i]
		//operation := instructions[instruction]
		i = inst.opcode(i, intops, inst.paramModes)

	}
	return intops

}
