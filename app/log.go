package main

import (
	"fmt"
	"time"
)

func filenameLog() string {
	var timenow = time.Now()
	year := timenow.Year()
	month := timenow.Month()
	day := timenow.Day()
	minute := timenow.Minute()
	second := timenow.Second()

	return fmt.Sprintf("log_%d%02d%02d%02d%02d.log", year, month, day, minute, second)
}
