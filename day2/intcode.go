package main

const (
	ADD = 1
	MULTIPLY = 2
	END = 99
)

func CalcIntOps(intops map[int]int) map[int]int {
	//index := make([]int,len(intops))


	for i := 0; i < len(intops); i = i + 4 {
		operation := intops[i]
		if operation == ADD {
			resultPos := intops[i+3]
			sum1Pos := intops[i+1]
			sum2Pos := intops[i+2]
			intops[resultPos] =  intops[sum1Pos] + intops[sum2Pos]
		}

		if operation == MULTIPLY {
			resultPos := intops[i+3]
			sum1Pos := intops[i+1]
			sum2Pos := intops[i+2]
			intops[resultPos] =  intops[sum1Pos] * intops[sum2Pos]
		}

		if operation== END{
			break
		}
	}
	return intops

}
