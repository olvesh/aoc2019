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

const testOrbits = `B)C
C)D
D)E
COM)B
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

const testDistanceOrbits = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

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

	//massMap := NewMassMapStrings(testDistanceOrbits)
	log.Print(massMap.COM.NumDirectAndIndirect())
	log.Print(massMap.distance("YOU", "SAN"))
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

func NewMassMapStrings(oorbits string) *MassMap {
	orbitsRaw := strings.Split(oorbits, "\n")
	massMap := NewMassMap(len(orbitsRaw))

	for _, s := range orbitsRaw {
		massMap.add(s)
	}
	return massMap
}

func NewMassMap(cap int) *MassMap {
	massMap := &MassMap{
		idx: make(map[string]*Mass, cap),
	}
	com := &Mass{
		Id:      "COM",
		Orbits:  nil,
		InOrbit: make([]*Mass, 0),
	}
	massMap.idx["COM"] = com
	massMap.COM = com
	return massMap
}
