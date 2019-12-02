package main

import "log"

//1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,9,23,1,5,23,27,1,27,9,31,1,6,31,35,2,35,9,39,1,39,6,43,2,9,43,47,1,47,6,51,2,51,9,55,1,5,55,59,2,59,6,63,1,9,63,67,1,67,10,71,1,71,13,75,2,13,75,79,1,6,79,83,2,9,83,87,1,87,6,91,2,10,91,95,2,13,95,99,1,9,99,103,1,5,103,107,2,9,107,111,1,111,5,115,1,115,5,119,1,10,119,123,1,13,123,127,1,2,127,131,1,131,13,0,99,2,14,0,0
var gravityAssistProgramInput = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 19, 9, 23, 1, 5, 23, 27, 1, 27, 9, 31, 1, 6, 31, 35, 2, 35, 9, 39, 1, 39, 6, 43, 2, 9, 43, 47, 1, 47, 6, 51, 2, 51, 9, 55, 1, 5, 55, 59, 2, 59, 6, 63, 1, 9, 63, 67, 1, 67, 10, 71, 1, 71, 13, 75, 2, 13, 75, 79, 1, 6, 79, 83, 2, 9, 83, 87, 1, 87, 6, 91, 2, 10, 91, 95, 2, 13, 95, 99, 1, 9, 99, 103, 1, 5, 103, 107, 2, 9, 107, 111, 1, 111, 5, 115, 1, 115, 5, 119, 1, 10, 119, 123, 1, 13, 123, 127, 1, 2, 127, 131, 1, 131, 13, 0, 99, 2, 14, 0, 0}

var gravityAssistProgram = func() map[int]int {
	m := make(map[int]int, len(gravityAssistProgramInput))
	for i, val := range gravityAssistProgramInput {
		m[i] = val
	}
	return m
}

func main() {

	noun := 12
	verb := 2

	res := execIntcode(gravityAssistProgram(), noun, verb)

	log.Printf("Intops #1: %v", res)

	for noun := 0; noun < 100; noun = noun + 1 {
		for verb := 0; verb < 100; verb = verb + 1 {
			res := execIntcode(gravityAssistProgram(), noun, verb)
			if res == 19690720 {
				log.Printf("19690720 reached, noun: %v verb %v", noun, verb)
			}
		}
	}

}

func execIntcode(gravityAssistProgram map[int]int, noun int, verb int) int {
	gravityAssistProgram[1] = noun
	gravityAssistProgram[2] = verb
	CalcIntOps(gravityAssistProgram)
	return gravityAssistProgram[0]
}
