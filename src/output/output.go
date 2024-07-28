package output

import "fmt"

var schedules = []string{"minute", "hour", "day of month", "month", "day of week"}

func Display(commandToExec string, scheduleList [][]int) {
	if len(scheduleList) != 5 {
		return
	}
	for position, schedule := range scheduleList {
		sch := ""
		for _, s := range schedule {
			sch += fmt.Sprintf("%d ", s)
		}
		fmt.Printf("%-14s%s\n", schedules[position], sch)
	}
	fmt.Printf("%-14s%s\n", "command", commandToExec)
}
