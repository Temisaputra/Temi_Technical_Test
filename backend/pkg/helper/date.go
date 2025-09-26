package helper

import (
	"fmt"
	"time"
)

func GetListMonthOfYear(year int) []int {
	if year == time.Now().Year() {
		return GetListMonthOfCurrentYear()

	}

	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
}

func GetListMonthOfCurrentYear() []int {
	month := time.Now().Month()
	months := make([]int, 0, int(month))

	for i := 1; i <= int(month); i++ {
		months = append(months, i)
	}

	return months
}

func GetFirstDayAndLastDayOfTheMonth(year, month int) (string, string) {
	months := []string{
		"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12",
	}

	monthStr := months[month-1]

	firstDay := fmt.Sprintf("%d-%s-01", year, monthStr)
	lastDay := fmt.Sprintf("%d-%s-%d", year, monthStr, GetLastDayOfMonth(month, year))

	return firstDay, lastDay
}

func GetLastDayOfMonth(month, year int) int {
	if month == 2 {
		if year%4 == 0 {
			return 29
		}

		return 28
	}

	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}

	return 31
}

func GetFirstDateAndLastDateOfTheYear(year int) (string, string) {
	firstDate := fmt.Sprintf("%d-01-01", year)
	lastDate := fmt.Sprintf("%d-12-31", year)

	return firstDate, lastDate
}

func GetFirstDateAndLastDateOfTwoMonth() (string, string, string, string) {
	year := time.Now().Year()
	month := time.Now().Month()

	firstDate := fmt.Sprintf("%d-%02d-01", year, month)
	lastDate := fmt.Sprintf("%d-%02d-%d", year, month, GetLastDayOfMonth(int(month), year))

	if month == 1 {
		year--
		month = 12
	} else {
		month--
	}

	firstDate2 := fmt.Sprintf("%d-%02d-01", year, month)
	lastDate2 := fmt.Sprintf("%d-%02d-%d", year, month, GetLastDayOfMonth(int(month), year))

	return firstDate, lastDate, firstDate2, lastDate2
}

func GetFirstDateAndLastDateOfTwoYear() (string, string, string, string) {
	year := time.Now().Year()

	firstDate := fmt.Sprintf("%d-01-01", year)
	lastDate := fmt.Sprintf("%d-12-31", year)

	year--

	firstDate2 := fmt.Sprintf("%d-01-01", year)
	lastDate2 := fmt.Sprintf("%d-12-31", year)

	return firstDate, lastDate, firstDate2, lastDate2
}

func GetDay(day int) string {
	return fmt.Sprintf("%02d", day)
}
