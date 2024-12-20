// Copyright 2024 Notedown Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package time

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/liamawhite/parse/core"
)

// Parse a date in the format yyyy-MM-dd.
var YearMonthDay = func(in Input) (time.Time, bool, error) {
	// Create parsers for year, month and day.
	year := StringFrom(Times(4, Digit))
	month := StringFrom(RuneIn("01"), Digit)
	day := StringFrom(RuneIn("0123"), Digit)

	// Create string parser for yyyy-MM-dd.
	date := StringFrom(All(year, Rune('-'), month, Rune('-'), day))

	s, ok, err := date(in)
	if err != nil || !ok {
		return time.Time{}, false, err
	}

	// Parse the date.
	match, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, false, fmt.Errorf("failed to parse date: %w", err)
	}

	return match, true, nil
}

// mon, monday, tue, tues, tuesday, wed, weds, wednesday, thu, thur, thurs, thursday, fri, friday, sat, saturday, sun, sunday
var DayOfWeek = func(in Input) (match time.Weekday, ok bool, err error) {
	m := map[time.Weekday]Parser[string]{
		// Ordering is important! Use more specific parsers first.
		time.Monday:    Any(StringInsensitive("monday"), StringInsensitive("mon")),
		time.Tuesday:   Any(StringInsensitive("tuesday"), StringInsensitive("tues"), StringInsensitive("tue")),
		time.Wednesday: Any(StringInsensitive("wednesday"), StringInsensitive("weds"), StringInsensitive("wed")),
		time.Thursday:  Any(StringInsensitive("thursday"), StringInsensitive("thurs"), StringInsensitive("thur"), StringInsensitive("thu")),
		time.Friday:    Any(StringInsensitive("friday"), StringInsensitive("fri")),
		time.Saturday:  Any(StringInsensitive("saturday"), StringInsensitive("sat")),
		time.Sunday:    Any(StringInsensitive("sunday"), StringInsensitive("sun")),
	}

	for day, parser := range m {
		_, ok, err := parser(in)
		if err != nil {
			return time.Weekday(-1), false, err
		}
		if ok {
			return day, true, nil
		}
	}

	return time.Weekday(-1), false, nil
}

// Space separated list of days of the week.
var DaysOfWeek = func(in Input) (match []time.Weekday, ok bool, err error) {
	var days []time.Weekday

	delimiter := RuneIn(" ")

	for {
		day, ok, err := DayOfWeek(in)
		if err != nil {
			return nil, false, err
		}
		if !ok {
			break
		}
		_, ok, err = delimiter(in)
		if err != nil {
			return nil, false, err
		}
		days = append(days, day)
	}

	if len(days) == 0 {
		return nil, false, nil
	}

	return days, true, nil
}

// Parse a number followed by an optional ordinal (st, nd, rd, th).
var MonthDay = func(in Input) (match int, ok bool, err error) {
	n, ok, err := StringFrom(AtLeast(1, Digit))(in)
	if err != nil {
		return 0, false, err
	}
	if !ok {
		return 0, false, nil
	}

	// Parse an optional ordinal.
	ordinal := Any(StringInsensitive("st"), StringInsensitive("nd"), StringInsensitive("rd"), StringInsensitive("th"))
	_, _, err = ordinal(in)
	if err != nil {
		return 0, false, err
	}

	// Convert the number to an integer.
	number, err := strconv.Atoi(n)
	if err != nil {
		return 0, false, fmt.Errorf("failed to parse number: %w", err)
	}

	return number, true, nil
}

var MonthOfYear = func(in Input) (match time.Month, ok bool, err error) {
	m := map[time.Month]Parser[string]{
		// Ordering is important! Use more specific parsers first.
		time.January:   Any(StringInsensitive("january"), StringInsensitive("jan")),
		time.February:  Any(StringInsensitive("february"), StringInsensitive("feb")),
		time.March:     Any(StringInsensitive("march"), StringInsensitive("mar")),
		time.April:     Any(StringInsensitive("april"), StringInsensitive("apr")),
		time.May:       StringInsensitive("may"),
		time.June:      Any(StringInsensitive("june"), StringInsensitive("jun")),
		time.July:      Any(StringInsensitive("july"), StringInsensitive("jul")),
		time.August:    Any(StringInsensitive("august"), StringInsensitive("aug")),
		time.September: Any(StringInsensitive("september"), StringInsensitive("sept"), StringInsensitive("sep")),
		time.October:   Any(StringInsensitive("october"), StringInsensitive("oct")),
		time.November:  Any(StringInsensitive("november"), StringInsensitive("nov")),
		time.December:  Any(StringInsensitive("december"), StringInsensitive("dec")),
	}

	for month, parser := range m {
		_, ok, err := parser(in)
		if err != nil {
			return time.Month(-1), false, err
		}
		if ok {
			return month, true, nil
		}
	}

	return 0, false, nil
}
