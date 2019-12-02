package main

import (
	"aoc2020/day1/part1"
	"log"
	"os"
)

func main() {
	modules, err := os.Open("./modules.txt")
	defer func(modules *os.File) {
		err := modules.Close()
		log.Fatal("Unable to lose file", err)
	}(modules)

	if err != nil {
		log.Panic(err)
	}
	part1.CalculateFuelForAllModules(modules)
	part1.CalculateFuelForAllModules(modules)


}
