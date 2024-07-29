package cron

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testList := []struct {
		name       string
		input      []string
		want       [][]int
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid input",
			input:   []string{"0", "0,12", "1", "*/2", "*"},
			want:    [][]int{{0}, {0, 12}, {1}, {1, 3, 5, 7, 9, 11}, {0, 1, 2, 3, 4, 5, 6}},
			wantErr: false,
		},
		{
			name:    "Valid input",
			input:   []string{"*/5", "4,6,8", "1-20/4", "*/2", "6"},
			want:    [][]int{{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55}, {4, 6, 8}, {1, 5, 9, 13, 17}, {1, 3, 5, 7, 9, 11}, {6}},
			wantErr: false,
		},
	}
	for _, test := range testList {
		t.Run(test.name, func(t *testing.T) {
			got, err := Parse(test.input)
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
