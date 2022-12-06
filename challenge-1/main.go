package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	previousMax := 0
	currentMax := 0

	snacks := strings.Split(string(b), "\n")
	for _, calories := range snacks {
		if calories == "" {
			if currentMax > previousMax {
				previousMax = currentMax
			}

			currentMax = 0
			continue
		}

		ivalue, _ := strconv.Atoi(calories)
		currentMax += ivalue
	}

	if previousMax > currentMax {
		currentMax = previousMax
	}

	log.Printf("Max: %d\n", currentMax)
}
