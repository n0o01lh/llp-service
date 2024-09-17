package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func SanitizeNumericParam(param string) (string, error) {
	if _, err := strconv.Atoi(param); err != nil {
		return "", fmt.Errorf("param #{param} is not a number")
	}
	return param, nil
}

func SanitizeArrayParam(arrayParam string) []string {
	sanitizedParams := make([]string, 0)
	params := strings.Split(arrayParam, ",")
	for _, param := range params {
		if _, err := SanitizeNumericParam(param); err == nil {
			sanitizedParams = append(sanitizedParams, param)
		}
	}
	return sanitizedParams
}

func SanitizeArrayParamString(arrayParam string) []string {
	sanitizedParams := make([]string, 0)
	params := strings.Split(arrayParam, ",")
	for _, param := range params {
		if param != "" {
			sanitizedParams = append(sanitizedParams, param)
		}
	}
	return sanitizedParams
}

func SanitizeDateParam(datePparam string) (time.Time, error) {
	return time.Parse("02/01/2006", datePparam)
}
