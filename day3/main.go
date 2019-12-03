package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	wires, err := os.Open("./day3/wires.txt")
	defer func(wires *os.File) {
		err := wires.Close()
		log.Fatal("Unable to close file", err)
	}(wires)

	if err != nil {
		log.Panic(err)
	}

	lines := make([]string, 0, 2)
	scanner := bufio.NewScanner(wires)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	distance := Distance(NewWire(lines[0]), NewWire(lines[1]))
	log.Printf("Distance is %v", distance)

	steps := Steps(NewWire(lines[0]), NewWire(lines[1]))
	log.Printf("Steps is %v", steps)

}
