package parser

import (
	"bytes"
	"github.com/ouesfa/go-mower/models"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestParseBoundaries(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		want      models.Boundaries
		wantPanic string
	}{
		{name: "should return boundaries given valid input", args: struct{ input string }{input: "2 2"}, want: models.Boundaries{2, 2}},
		{name: "should panic given invalid input", args: struct{ input string }{input: "invalid"}, wantPanic: "failed to Parse boundaries"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r != nil && r != tt.wantPanic {
					t.Errorf("recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := parseBoundaries(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBoundaries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAction(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *models.Action
		log  string
	}{
		{name: "should return action given valid input", args: args{input: "RRRLFF"}, want: &models.Action{Instructions: "RRRLFF"}},
		{name: "should return nil given invalid input", args: args{input: "&é4556"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			got := parseAction(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) && strings.HasSuffix(buf.String(), tt.log) {
				t.Errorf("parseAction() = %v, want %v", got, tt.want)
				t.Errorf("actual log = %v, expected log %v", buf.String(), tt.log)
			}
		})
	}
}

func TestParseMower(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *models.Mower
	}{
		{"should return mower given valid input", args{input: "1 0 N"}, &models.Mower{X: 1, Orientation: "N"}},
		{name: "should return nil given invalid orientation", args: args{input: "0 0 Z"}},
		{"should return nil given invalid coordinates", args{input: "22 0 N"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMower(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		input string
	}
	lines := `5 5
1 2 N
LFLFLFLFF`
	tests := []struct {
		name  string
		args  args
		want  models.Boundaries
		want1 []models.MowerWithIndex
		want2 []models.ActionWithIndex
	}{
		{"should Parse valid input into boundaries, mowers & actions",
			args{lines},
			models.Boundaries{5, 5},
			[]models.MowerWithIndex{{0, models.Mower{1, 2, "N"}}},
			[]models.ActionWithIndex{{1, models.Action{Instructions: "LFLFLFLFF"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Parse(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Parse() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
