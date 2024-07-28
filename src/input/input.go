package input

import (
	"errors"
	"os"
	"strings"
)

// this struct will hold the expression and command that is to be executed
type Data struct {
	elementList   []string // private data member
	commandToExec string   // private data member
}

func (d *Data) GetElementList() []string {
	return d.elementList
}

func (d *Data) GetCommandToExec() string {
	return d.commandToExec
}

// this function will be called from the main file to get the command line cron expression related data
func GetData() (*Data, error) {
	/*
		The length of the command line argument should be atleast '2', since
		the program is expecting the entire cron job as a single string.
		Ex:
		a) go run main.go "* * * * * /etc/bin/program.sh"
		b) ./<executable> "* * * * * /etc/bin/program.sh"
	*/
	args := os.Args
	if len(args) < 2 {
		return nil, errors.New("invalid number of argument")
	}
	// splitting the argument on whitespace
	parts := strings.Split(args[1], " ")
	if len(parts) < 6 {
		return nil, errors.New("invalid cron expression")
	}

	// getting the list of expression elements
	elementList := make([]string, 0)
	for i := 0; i <= 4; i++ {
		elementList = append(elementList, parts[i])
	}
	// getting the command to execute
	commandToExec := strings.Join(parts[5:], " ")

	d := &Data{
		elementList:   elementList,
		commandToExec: commandToExec,
	}
	return d, nil
}
