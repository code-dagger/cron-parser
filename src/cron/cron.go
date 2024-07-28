package cron

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

/*
Expression: * * * * *

	Index 0: minute  (allowed range: 0 to 59)
	Index 1: hour (allowed range: 0 to 23)
	Index 2: day of month (allowed range: 1 to 31)
	Index 4: month (allowed range: 1 to 12)
	Index 5: day of week (allowed range: 0 to 6)
*/
var allowedRange = [][2]int{{0, 59}, {0, 23}, {1, 31}, {1, 12}, {0, 6}}
var delimiterRegex = regexp.MustCompile(`[*\/\-,]+`)

func Parse(elementList []string) ([][]int, error) {
	if len(elementList) != 5 {
		return nil, errors.New("expression should contain exactly 5 elements")
	}
	scheduleList := make([][]int, 5)
	for position, element := range elementList {
		valList, parseErr := parseElement(position, element)
		if parseErr != nil {
			return nil, parseErr
		}
		// checking the range allowed based on element index
		if valList[0] < allowedRange[position][0] {
			errStr := fmt.Sprintf("left value %d cannot be less than %d for element: %s", valList[0], allowedRange[position][0], element)
			return nil, errors.New(errStr)
		}
		if valList[len(valList)-1] > allowedRange[position][1] {
			errStr := fmt.Sprintf("right value %d cannot be greater than %d for element: %s", valList[len(valList)-1], allowedRange[position][1], element)
			return nil, errors.New(errStr)
		}
		scheduleList[position] = valList
	}
	return scheduleList, nil
}

func parseElement(position int, element string) ([]int, error) {
	// first we will replace all the wildcard with a range of value based on the element index,
	// this will modify all the */2 or * based expression with range 1-24/2 or 1-24 respectively
	// based on position of the element
	element = strings.ReplaceAll(element, Wildcard, fmt.Sprintf("%d-%d", allowedRange[position][0], allowedRange[position][1]))
	// now we need to get the type of the delimiter used to parse the element with the appropriate values
	elementType, typeErr := getElementType(element)
	if typeErr != nil {
		return nil, typeErr
	}
	switch elementType {
	case TypeList:
		return parseList(element)
	case TypeRange:
		return parseRange(element)
	case TypeStep:
		return parseStep(element)
	case TypeFixed:
		return parseFixed(element)
	}
	return nil, fmt.Errorf("invalid expression element: %s no valid element type matches", element)
}

func getElementType(element string) (ElementType, error) {
	// getting all the delimiter which are present in the current element of expression using regex matches
	matches := delimiterRegex.FindAllString(element, -1)
	delimiterList := append([]string{}, matches...)
	delimiterListLen := len(delimiterList)

	if delimiterListLen > 1 && delimiterList[1] == DelimiterStep {
		return TypeStep, nil
	} else if delimiterListLen >= 1 && delimiterList[0] == DelimiterList {
		return TypeList, nil
	}
	if delimiterListLen == 0 {
		return TypeFixed, nil
	}
	switch delimiterList[0] {
	case DelimiterList:
		return TypeList, nil
	case DelimiterRange:
		return TypeRange, nil
	default:
		return TypeInvalid, fmt.Errorf("invalid delimiter: %s in expression element: %s", delimiterList[0], element)
	}
}
