package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func String2Time(dt string) (time.Time, error) {
	data := strings.Split(dt, "/")
	if len(data) != 3 {
		return time.Time{}, fmt.Errorf("Date invalid")
	}
	y, err := strconv.Atoi(data[2])
	if err != nil {
		return time.Time{}, err
	}
	m, err := strconv.Atoi(data[1])
	if err != nil {
		return time.Time{}, err
	}
	d, err := strconv.Atoi(data[0])

	if err != nil {
		return time.Time{}, err
	}
	result := time.Date(y, time.Month(m), d, 12, 15, 5, 5, time.Local)
	return result, nil
}
