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
		{name: "should return boundaries given valid input", args: struct{ input string }{input: string("2 2")}, want: *models.NewBoundaries("2", "2")},
		{name: "should panic given invalid input", args: struct{ input string }{input: string("invalid")}, wantPanic: "failed to parse boundaries"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r != nil && r != tt.wantPanic {
					t.Errorf("recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := ParseBoundaries(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBoundaries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseActions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *models.Action
		log  string
	}{
		{"should return action given valid input", args{input: "RRRLFF"}, &models.Action{Instructions: "RRRLFF"}, ""},
		{"should return nil given invalid input", args{input: "&é4556"}, nil, "toto"},
		//{"should log specific message given invalid input", args{input: "&é4556"}, nil, "toto"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			got := ParseActions(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) && strings.HasSuffix(buf.String(), tt.log) {
				t.Errorf("ParseActions() = %v, want %v", got, tt.want)
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
		{"should return mower given valid input", args{input: "1 0 N"}, models.NewMower("1", "0", "N")},
		{"should return nil given invalid orientation", args{input: "0 0 Z"}, nil},
		{"should return nil given invalid coordinates", args{input: "22 0 N"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseMower(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMower() = %v, want %v", got, tt.want)
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
		want1 []models.Mower
		want2 []models.Action
	}{
		{"should parse valid input into boundaries, mowers & actions",
			args{lines},
			*models.NewBoundaries("5", "5"),
			[]models.Mower{*models.NewMower("1", "2", "N")},
			[]models.Action{{Instructions: "LFLFLFLFF"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := parse(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("parse() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
