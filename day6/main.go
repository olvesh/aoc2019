package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const testOrbits2 = `COM)B
B)C
C)D`

const testOrbits = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

func main() {

	puzzle, err := os.Open("./day6/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(wires *os.File) {
		err := wires.Close()
		log.Fatal("Unable to close file", err)
	}(puzzle)

	massMap := NewMassFromReader(puzzle)

	log.Print(massMap.COM.NumDirectAndIndirect())

}

func NewMassFromReader(puzzle *os.File) *MassMap {
	massMap := &MassMap{
		idx: make(map[string]*Mass, 1057),
	}
	scanner := bufio.NewScanner(puzzle)
	for scanner.Scan() {
		link := scanner.Text()
		massMap.add(link)
	}
	return massMap

}

func NewMass(oorbits string) *MassMap {
	orbitsRaw := strings.Split(oorbits, "\n")

	massMap := &MassMap{
		idx: make(map[string]*Mass, len(orbitsRaw)),
	}

	for _, s := range orbitsRaw {
		massMap.add(s)
	}
	return massMap
}
