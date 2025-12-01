package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Move struct {
	Direction string
	Distance  int
}

func parseCombination(input []string) []Move {

	moves := make([]Move, 0, len(input))

	for _, line := range input {
		direction := string(line[0])            // L or R
		distance, err := strconv.Atoi(line[1:]) // everything else

		if err != nil {
			fmt.Printf("invalid distance in %s", line)
		}

		moves = append(moves, Move{
			Direction: direction,
			Distance:  distance,
		})
	}

	return moves
}

func readInputs() []string {
	// Read the input file line by line

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// close file when done reading
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func unlock(moves []Move) (int, int) {
	timesEndedAtZero := 0
	totalClicksAtZero := 0
	currentPosition := 50

	for _, move := range moves {
		switch move.Direction {
		case "L":

			clicksAt0 := 0

			if move.Distance >= currentPosition && currentPosition > 0 {
				// (part 1) reset to 99, subtract the remainder & count how many times 0 is passed
				remaining := move.Distance - currentPosition
				clicksAt0 = 1 + (remaining / 100)
			} else if currentPosition == 0 && move.Distance > 0 {
				// (part 2) starting at 0, we come back to 0 every 100 clicks
				clicksAt0 = move.Distance / 100
			}

			totalClicksAtZero += clicksAt0

			currentPosition = currentPosition - move.Distance
			if currentPosition < 0 {
				currentPosition = ((currentPosition % 100) + 100) % 100
			}

		case "R":

			clicksAt0 := 0

			totalDistance := currentPosition + move.Distance
			// (part 2) determine how many times we exceed 99
			clicksAt0 = totalDistance / 100

			totalClicksAtZero += clicksAt0

			currentPosition = totalDistance % 100
		}

		// (part 1) count when we end on 0
		if currentPosition == 0 {
			timesEndedAtZero++
		}
	}

	return timesEndedAtZero, totalClicksAtZero
}

func main() {

	instructions := readInputs()

	moves := parseCombination(instructions)

	for _, move := range moves {
		fmt.Printf("Direction: %s, Distance: %d\n", move.Direction, move.Distance)
	}

	timesAt, timesPassed := unlock(moves)
	fmt.Printf("Times at 0: %d\nTimes passed 0: %d\n", timesAt, timesPassed)

}
