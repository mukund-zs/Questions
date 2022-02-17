package main

//checkLeapYear checks the year is leap year or not
func checkLeapYear(year int) string {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return "It's not a leap year"
		}
		return "It's a leap year"
	}
	return "It's not a leap year"
}
