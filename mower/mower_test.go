package mower

import (
	"github.com/ouesfa/go-mower/models"
	"reflect"
	"testing"
)

func Test_move(t *testing.T) {
	type args struct {
		instruction rune
		mower       models.Mower
		boundaries  models.Boundaries
	}
	defaultMower := *models.NewMower("1", "1", "N")
	defaultBoundaries := *models.NewBoundaries("5", "5")
	tests := []struct {
		name string
		args args
		want models.Mower
	}{
		{"tototo", args{'L', defaultMower, defaultBoundaries}, *models.NewMower("1", "1", "W")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.instruction, tt.args.mower, tt.args.boundaries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveForward(t *testing.T) {
	type args struct {
		mower      models.Mower
		boundaries models.Boundaries
	}
	tests := []struct {
		name string
		args args
		want models.Mower
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveForward(tt.args.mower, tt.args.boundaries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveForward() = %v, want %v", got, tt.want)
			}
		})
	}
}
