package util

import "time"

var Months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

func GetMonthByInt(monthInt int) string {
	return Months[(monthInt - 1)]
}

func DaysInMonth(year, month int) []int {

	t := time.Date(year, time.Month(month), 32, 0, 0, 0, 0, time.Local)
	daysInMonth := 32 - t.Day()
	days := make([]int, daysInMonth)
	for i := range days {
		days[i] = i + 1
	}

	return days

}
