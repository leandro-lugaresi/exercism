//Package leap implements one function to
//discover if a year is leap or not
package leap

const testVersion = 2

/*
IsLeapYear take a year and returns if that year is a leap year
with these rules:
	on every year that is evenly divisible by 4
	except every year that is evenly divisible by 100
	unless the year is also evenly divisible by 400
*/
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
