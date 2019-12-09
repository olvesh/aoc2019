package main

import (
	"fmt"
)

func MaxAmplitude(prog []int) int {
	maxPhases := 5
	maxPhaseVal := 5
	maxPhaseValMod := 5

	res := make([]int, 255)

	//
	x := 0
	for i := 0; i < len(res); i = i + 1 {

		phases := make([]int, maxPhases)

		fmt.Printf("%v: ", i)

		for j := 0; j < maxPhases; j++ {
			phases[j] = (j + x) % maxPhaseValMod

			x++
		}
		x++
		maxPhaseValMod = maxPhaseVal

		fmt.Printf("%v \n", phases)

	}
	//  for j := 0; j < maxPhaseVal; j = j + 1 {
	//    for k := 0; k < maxPhaseVal; k = k + 1 {
	//      for l := 0; l < maxPhaseVal; l = l + 1 {
	//        for m := 0; m < maxPhaseVal; m = m + 1 {
	//          ints := []int{i, j, k, l, m}
	//          rack := NewAmpRack(ints, prog)
	//          res = append(res, rack)
	//        }
	//      }
	//    }
	//  }
	//}
	//sort.Ints(res)
	//return res[le
	//n(res)-1]
	return 0

	//
	//for i:= 0; i < len(res); i = i+1 {
	//  log.Printf("%v %v %v %v %v", i, i%4, (i+1)%8, (i+2)%, (i+3)%4)
	//}
}

//https://stackoverflow.com/questions/1817631/iterating-one-dimension-array-as-two-dimension-array
/*int[10] oneDim = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
int rows = 2;
int columns = 5;
for(int index = 0; index < 10; index++) {
    int column = index % columns;
    int row = (index - column) / columns;
    System.out.println(row + ", " + column + ": " + oneDim[index]);
}*/

func NewAmpRack(phases []int, prog []int) int {
	//amps := make([]Amp, len(phases))

	currentOut := 0
	for i, phase := range phases {

		callnum := 0

		curentIC := NewIntcodeComputerOverrideInOut(
			prog,
			func() int {
				defer func() { callnum++ }()
				if callnum == 1 {
					if i == 0 {
						return 0
					}
					return currentOut
				}
				return phase
			},
			func(out int) {
				currentOut = out
			})

		curentIC.Exec()
		callnum = 0
	}

	return currentOut
}

type Amp struct {
	IntcodeComputer IntcodeComputer
} /*

func NewAmp(phases []int, program []int) int {

  firstIn := func() int {
    return phases[0]
  }

  out := -1
  lastOut := func(val int) {
    out = val
  }

  return out
}

func createConnectedInOut() (In, Out) {

}
*/
