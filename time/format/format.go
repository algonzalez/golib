// Package format provides various time format constants
// as well as functions to work with time format strings.
package format

import "strings"

const (
	DATE     string = "2006-01-02"
	TIME     string = "15:04:05"
	DATETIME string = "2006-01-02T15:04:05"
	OFFSET   string = "-07:00"

	MILLISECONDS string = ".000"
	MICROSECONDS string = ".000000"
	NANOSECONDS  string = ".000000000"
)

// Function to convert .NET DateTime format string
// to the Golang Time style format string.
//
//  Month:
//  - MMMM for full month name (e.g. January)
//  - MMM for abbreviated month name (e.g. Jan)
//  - MM zero padded month number (e.g. Jan is 01)
//  - M non-zero padded month number (e.g. Jan is 1)
//  Day:
//  - dddd for full weekday name (e.g. Monday)
//  - ddd for abbreviated weekday name (e.g. Mon)
//  - dd zero padded day (e.g. 02)
//  - d non-zero padded day (e.g. 2)
//  Year:
//  - yyyy full four digit year (e.g. 2006)
//  - yy zero padded year (e.g. 2006 is 06)
//  - y non-zero padded year (e.g. 2006 is 6)
//  Hour:
//  - HH 24-hour format hour (e.g. 15)
//  - hh zero padded hour (e.g. 03)
//  - h non-zero padded hour (e.g. 3)
//  Minutes:
//  - mm zero padded minute (e.g. 04)
//  - m non-zero padded minute (e.g. 4)
//  Seconds:
//  - ss zero padded seconds (e.g. 05)
//  - s non-zero padded seconds (e.g. 5)
//  Seconds fractions:
//  - .fff milliseconds
//  - .ffffff microseconds
//  - .fffffffff nanoseconds
//  AM/PM:
//  - TT capitalized AM/PM designator (e.g. 3AM)
//  - tt lower cased AM/PM designator (e.g. 3am)
//  Timezone:
//  - TZ timezone abbreviation (e.g. MST for Mountain Standard Time)
//  Timezone offset:
//  - zzzz hour and minute offset (e.g. -07:00)
//  - zz just hour offset (e.g. -07)
func DotNetToGolangStyle(format string) string {
	r := strings.NewReplacer(
		"MMMM", "January",
		"MMM", "Jan",
		"MM", "01",
		"M", "1",
		"dddd", "Monday",
		"ddd", "Mon",
		"dd", "02",
		"d", "2",
		"HH", "15",
		"hh", "03",
		"h", "3",
		"mm", "04",
		"m", "4",
		"ss", "05",
		"s", "5",
		"yyyy", "2006",
		"yy", "06",
		"y", "6",
		"zzzz", "-07:00",
		"zz", "-07",
		"TZ", "MST",
		"TT", "PM",
		"tt", "pm",
		// NOTE: choosing to use zeros, but nines would also work -- nines drop any trailing zeros
		// Consider making this configurable
		".fff", ".000", // milliseconds
		".ffffff", ".000000", // microseconds
		".fffffffff", ".000000000", // nanoseconds
	)
	return r.Replace(format)
}

// Function to convert Golang Time style format string
// to the .NET DateTime format string.
func GolangToDotNetStyle(format string) string {
	r := strings.NewReplacer(
		"2006", "yyyy",
		"06", "yy",
		"6", "y",
		"January", "MMMM",
		"Jan", "MMM",
		"01", "MM",
		"1", "M",
		"Monday", "dddd",
		"Mon", "ddd",
		"02", "dd",
		"2", "d",
		"15", "HH",
		"03", "hh",
		"3", "h",
		"04", "mm",
		"4", "m",
		"05", "ss",
		"5", "s",
		"-07:00", "zzzz",
		"-07", "zz",
		"MST", "TZ",
		"PM", "TT",
		"pm", "tt",
		".000", ".fff", // milliseconds
		".000000", ".ffffff", // microseconds
		".000000000", ".fffffffff", // nanoseconds
		".999", ".fff", // milliseconds drop trailing zeros
		".999999", ".ffffff", // microseconds drop trailing zeros
		".999999999", ".fffffffff", // nanoseconds drop trailing zeros
	)
	return r.Replace(format)
}
