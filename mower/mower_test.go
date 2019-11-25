package main

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
	defaultMower := models.Mower{X: 1, Y: 1, Orientation: "N"}
	defaultBoundaries := models.Boundaries{NorthLimit: 5, EastLimit: 5}
	tests := []struct {
		name string
		args args
		want models.Mower
	}{
		{name: "should return West given instruction Left when orientation is North",
			args: args{
				instruction: 'L',
				mower:       defaultMower,
				boundaries:  defaultBoundaries},
			want: models.Mower{X: 1, Y: 1, Orientation: "W"}},
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
		{"should not overstep boundaries", args{
			mower:      models.Mower{X: 3, Y: 3, Orientation: "E"},
			boundaries: models.Boundaries{NorthLimit: 3, EastLimit: 3},
		}, models.Mower{X: 3, Y: 3, Orientation: "E"}},
		{"should decrement X if orientation is West", args{
			mower:      models.Mower{X: 3, Y: 3, Orientation: "W"},
			boundaries: models.Boundaries{NorthLimit: 3, EastLimit: 3},
		}, models.Mower{X: 2, Y: 3, Orientation: "W"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveForward(tt.args.mower, tt.args.boundaries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveForward() = %v, want %v", got, tt.want)
			}
		})
	}
}
