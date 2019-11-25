package parser

import (
	"github.com/alexflint/go-restructure"
	"github.com/ouesfa/go-mower/models"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) (models.Boundaries, []struct {
	Index int
	models.Mower
}, []struct {
	Index int
	models.Action
}) {
	lines := strings.Split(input, "\n")
	boundaries := ParseBoundaries(lines[0])
	var actions []struct {
		Index int
		models.Action
	}
	var mowers []struct {
		Index int
		models.Mower
	}
	for i, v := range lines[1:] {
		if action := ParseAction(v); action != nil && i%2 != 0 {
			actions = append(actions, struct {
				Index int
				models.Action
			}{i, *action})
		} else if mower := ParseMower(v); mower != nil && i%2 == 0 {
			mowers = append(mowers, struct {
				Index int
				models.Mower
			}{i, *mower})
		}
	}
	return boundaries, mowers, actions
}
func ParseBoundaries(input string) models.Boundaries {
	const boundariesPattern = `^(\d)\s(\d)$`
	re := regexp.MustCompile(boundariesPattern)
	if matchedBoundaries := re.FindStringSubmatch(input); matchedBoundaries == nil {
		panic("failed to Parse boundaries")
	} else {
		matchedNorthLimit, _ := strconv.Atoi(matchedBoundaries[1])
		matchedEastLimit, _ := strconv.Atoi(matchedBoundaries[2])
		return models.Boundaries{
			NorthLimit: matchedNorthLimit,
			EastLimit:  matchedEastLimit,
		}
	}
}
func ParseMower(input string) *models.Mower {
	const mowerPattern = `^(\d)\s(\d)\s([NSEW])$`
	re := regexp.MustCompile(mowerPattern)
	if matchedMower := re.FindStringSubmatch(input); matchedMower == nil {
		log.Println("failed to Parse mower from:", input)
		return nil
	} else {
		matchedX, _ := strconv.Atoi(matchedMower[1])
		matchedY, _ := strconv.Atoi(matchedMower[2])
		matchedOrientation := matchedMower[3]
		return &models.Mower{
			Orientation: matchedOrientation,
			X:           matchedX,
			Y:           matchedY,
		}
	}
}
func ParseAction(input string) *models.Action {
	var actions models.Action
	match, _ := restructure.Find(&actions, input)
	if !match {
		defer log.Println("failed to Parse actions")
		return nil
	}
	return &models.Action{Instructions: actions.Instructions}
}
