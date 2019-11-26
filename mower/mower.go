package main

import (
	"github.com/ouesfa/go-mower/models"
	"github.com/ouesfa/go-mower/parser"
	"io/ioutil"
	"log"
)

func main() {

	input := readFile("mowers")

	boundaries, mowersWithIndices, actionsWithIndices := parser.Parse(input)

	movedMowers := moveAll(mowersWithIndices, actionsWithIndices, boundaries)

	log.Println("input:\n" + input)
	log.Println("\nfinal positions:\n", movedMowers)
}

func moveAll(mowersWithIndices []models.MowerWithIndex, actionsWithIndices []models.ActionWithIndex, boundaries models.Boundaries) []models.Mower {
	var movedMowers []models.Mower
	for _, mowerWithIndex := range mowersWithIndices {
		for _, actionWithIndex := range actionsWithIndices {
			// bind a mower to its own action
			if mowerCorrespondsToAction := actionWithIndex.Index-mowerWithIndex.Index == 1; mowerCorrespondsToAction {
				for _, instruction := range actionWithIndex.Instructions {
					mowerWithIndex.Mower = move(instruction, mowerWithIndex.Mower, boundaries)
				}
				movedMowers = append(movedMowers, mowerWithIndex.Mower)
			}
		}
	}
	return movedMowers
}

func move(instruction rune, mower models.Mower, boundaries models.Boundaries) models.Mower {
	switch instruction {
	case 'L':
		mower.Orientation = leftOrientation[mower.Orientation]
	case 'R':
		mower.Orientation = rightOrientation[mower.Orientation]
	case 'F':
		mower = moveForward(mower, boundaries)
	}
	return mower
}

func moveForward(mower models.Mower, boundaries models.Boundaries) models.Mower {
	switch mower.Orientation {
	case "N":
		if mower.Y < boundaries.NorthLimit {
			mower.Y += 1
		}
	case "S":
		if mower.Y > 0 {
			mower.Y -= 1
		}
	case "E":
		if mower.X < boundaries.EastLimit {
			mower.X += 1
		}
	case "W":
		if mower.X > 0 {
			mower.X -= 1
		}
	}
	return mower
}

func readFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	// just pass the file name
	if err != nil {
		log.Fatal("failed to parse input file, ", err)
	}
	return string(b)
}

var rightOrientation = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}
var leftOrientation = map[string]string{
	"N": "W",
	"W": "S",
	"S": "E",
	"E": "N",
}
