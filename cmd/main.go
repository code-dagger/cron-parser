package main

import (
	"fmt"

	"github.com/code-dagger/cron-parser/src/cron"
	"github.com/code-dagger/cron-parser/src/input"
	"github.com/code-dagger/cron-parser/src/output"
)

func main() {
	// getting the input data
	data, inputErr := input.GetData()
	if inputErr != nil {
		fmt.Println(inputErr.Error())
		return
	}

	// cron parsing
	scheduleList, parseErr := cron.Parse(data.GetElementList())
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		return
	}

	// display
	output.Display(data.GetCommandToExec(), scheduleList)
}

// go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"
// go run main.go "0 0,12 1 */2 * /usr/bin/find"
// go run main.go "0 4 8-32 * * /usr/bin/find"
