package part2

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Module struct {
	Mass int
}

func CalculateFuelNeededForMass(module Module) int {
	mass := module.Mass
	return recursiveFuelCalc(mass)
}

func recursiveFuelCalc(mass int) int {
	i := (mass / 3) - 2
	if (i > 0 ) {
		i = i + recursiveFuelCalc(i)
	}
	if (i <= 0) {
		return 0
	}

	return i
}

func CalculateFuelForAllModules(modules *os.File) {


	fuelChan := make(chan int)

	// Go over a file line by line and queue up a ton of work
	go func() {

		scanner := bufio.NewScanner(modules)
		for scanner.Scan() {
			//go func() {
			moduleMass, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Panic(err)
			}
			fuelChan <- CalculateFuelNeededForMass(Module{Mass: moduleMass})
			//}()
		}
		close(fuelChan)
	}()

	// Collect all the results...
	// First, make sure we close the result channel when everything was processed
	//go func() {
	//	wg.Wait()
	//	close(results)
	//}()

	// Now, add up the results from the results channel until closed
	sum := 0
	for v := range fuelChan {
		sum += v
	}
	log.Printf("Part2: Fuel needed for total mass %v", sum)
}

