package cron

import (
	"reflect"
	"testing"
)

func TestParseRange(t *testing.T) {
	testList := []struct {
		name       string
		input      string
		want       []int
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid input",
			input:   "2-5",
			want:    []int{2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:       "Invalid input",
			input:      "6-3",
			wantErr:    true,
			errMessage: "left value cannot be greater than right value for expression element: 6-3",
		},
	}
	for _, test := range testList {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseRange(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if err != nil && err.Error() != test.errMessage {
				t.Errorf("Parse() error message = %v, wantErr %v", err.Error(), test.errMessage)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Parse() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestParseList(t *testing.T) {
	testList := []struct {
		name       string
		input      string
		want       []int
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid input",
			input:   "1,4,8",
			want:    []int{1, 4, 8},
			wantErr: false,
		},
		{
			name:       "Invalid input",
			input:      "-1,4,8",
			wantErr:    true,
			errMessage: "values cannot be negative for expression element: -1,4,8",
		},
		{
			name:       "Invalid input",
			input:      "4",
			wantErr:    true,
			errMessage: "invalid list expression for element: 4 it should contain atleast two values e.g a,b",
		},
	}
	for _, test := range testList {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseList(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if err != nil && err.Error() != test.errMessage {
				t.Errorf("Parse() error message = %v, wantErr %v", err.Error(), test.errMessage)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Parse() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestParseStep(t *testing.T) {
	testList := []struct {
		name       string
		input      string
		want       []int
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid input",
			input:   "1-10/2",
			want:    []int{1, 3, 5, 7, 9},
			wantErr: false,
		},
	}
	for _, test := range testList {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseStep(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if err != nil && err.Error() != test.errMessage {
				t.Errorf("Parse() error message = %v, wantErr %v", err.Error(), test.errMessage)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Parse() = %v, want %v", got, test.want)
			}
		})
	}
}
