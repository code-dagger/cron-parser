package cron

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	DelimiterList  = ","
	DelimiterRange = "-"
	DelimiterStep  = "/"
	DelimiterFixed = ""
	Wildcard       = "*"
)

const (
	ElementList  = "ElementList"
	ElementRange = "ElementRange"
	ElementStep  = "ElementStep"
	ElementFixed = "FixedElement"
)

// defining enum for element types
type ElementType int

const (
	TypeInvalid ElementType = iota
	TypeList
	TypeRange
	TypeStep
	TypeFixed
)

// function will parse cron expression that has hypen "-" in it
func parseRange(element string) ([]int, error) {
	rangeList := strings.Split(element, DelimiterRange)
	if len(rangeList) != 2 {
		return nil, fmt.Errorf("invalid range expression for element: %s it should contain two values e.g a-b", element)
	}
	left, leftErr := strconv.Atoi(rangeList[0])
	if leftErr != nil {
		return nil, leftErr
	}
	right, rightErr := strconv.Atoi(rangeList[1])
	if rightErr != nil {
		return nil, rightErr
	}
	if left < 0 {
		return nil, fmt.Errorf("left value cannot be negative for expression element: %s", element)
	} else if right < 0 {
		return nil, fmt.Errorf("right value cannot be negative for expression element: %s", element)
	} else if left > right {
		return nil, fmt.Errorf("left value cannot be greater than right value for expression element: %s", element)
	}
	valList := make([]int, 0)
	for i := left; i <= right; i++ {
		valList = append(valList, i)
	}
	return valList, nil
}

// function will parse cron expression that has comma "," in it
func parseList(element string) ([]int, error) {
	list := strings.Split(element, DelimiterList)
	if len(list) < 2 {
		return nil, fmt.Errorf("invalid list expression for element: %s it should contain atleast two values e.g a,b", element)
	}
	valList := make([]int, 0)
	for _, v := range list {
		val, valErr := strconv.Atoi(v)
		if valErr != nil {
			return nil, valErr
		}
		if val < 0 {
			return nil, fmt.Errorf("values cannot be negative for expression element: %s", element)
		}
		valList = append(valList, val)
	}
	// note: the expression values separated by "," should be in sorted (increasing) order
	sort.Ints(valList)
	return valList, nil
}

// function will parse cron expression that has slash "/" in it
func parseStep(element string) ([]int, error) {
	// note: left value of "/" can be range also
	steplist := strings.Split(element, DelimiterStep)
	if len(steplist) != 2 {
		return nil, fmt.Errorf("invalid step expression for element: %s it should contain exactly two values e.g a/b", element)
	}

	// parsing the left segment of the expression
	leftList := strings.Split(steplist[0], "-")
	leftRange := make([]int, 0)
	switch len(leftList) {
	case 1:
		val, valErr := strconv.Atoi(leftList[0])
		if valErr != nil {
			return nil, valErr
		}
		leftRange = append(leftRange, val)
	case 2:
		rangeList, rangeErr := parseRange(steplist[0])
		if rangeErr != nil {
			return nil, rangeErr
		}
		leftRange = rangeList
	default:
		return nil, fmt.Errorf("left value of step expression can contain atmost 2 range values, element: %s", element)
	}
	// parsing the right segment of the expression
	right, rightErr := strconv.Atoi(steplist[1])
	if rightErr != nil {
		return nil, rightErr
	}

	valList := make([]int, 0)
	for i := leftRange[0]; i <= leftRange[len(leftRange)-1]; i += right {
		valList = append(valList, i)
	}
	return valList, nil
}

func parseFixed(element string) ([]int, error) {
	val, valErr := strconv.Atoi(element)
	if valErr != nil {
		return nil, valErr
	}
	return []int{val}, nil
}
