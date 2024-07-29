package input

import (
	"os"
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	testList := []struct {
		name       string
		args       []string
		want       *Data
		wantErr    bool
		errMessage string
	}{
		{
			name:       "No Arguments",
			args:       []string{"./executable"},
			wantErr:    true,
			errMessage: "invalid number of argument",
		},
		{
			name:       "Invalid Cron Expression",
			args:       []string{"./executable", "* * * *"},
			wantErr:    true,
			errMessage: "invalid cron expression",
		},
		{
			name: "Valid Cron Expression",
			args: []string{"./executable", "* * * * * /etc/bin/program.sh"},
			want: &Data{
				elementList:   []string{"*", "*", "*", "*", "*"},
				commandToExec: "/etc/bin/program.sh",
			},
			wantErr: false,
		},
	}

	for _, test := range testList {
		t.Run(test.name, func(t *testing.T) {
			os.Args = test.args
			got, err := GetData()
			if (err != nil) != test.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if err != nil && err.Error() != test.errMessage {
				t.Errorf("GetData() error message = %v, want %v", err.Error(), test.errMessage)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("GetData() = %v, want %v", got, test.want)
			}
		})
	}
}
