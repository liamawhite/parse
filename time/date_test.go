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

package time_test

import (
	"testing"
	"time"

	. "github.com/liamawhite/parse/test"
	. "github.com/liamawhite/parse/time"
)

func TestYearMonthDay(t *testing.T) {
	tests := []ParserTest[time.Time]{
		{
			Name:          "match",
			Input:         "2021-01-01",
			Parser:        YearMonthDay,
			ExpectedMatch: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			ExpectedOK:    true,
		},
		{
			Name:       "invalid day high",
			Input:      "2021-01-32",
			Parser:     YearMonthDay,
			ExpectedOK: false,
			WantErr:    true,
		},
		{
			Name:       "invalid day low",
			Input:      "2021-01-00",
			Parser:     YearMonthDay,
			ExpectedOK: false,
			WantErr:    true,
		},
		{
			Name:       "invalid month high",
			Input:      "2021-13-01",
			Parser:     YearMonthDay,
			ExpectedOK: false,
			WantErr:    true,
		},
		{
			Name:       "invalid month low",
			Input:      "2021-00-01",
			Parser:     YearMonthDay,
			ExpectedOK: false,
			WantErr:    true,
		},
		{
			Name:          "leap year",
			Input:         "2024-02-29",
			Parser:        YearMonthDay,
			ExpectedMatch: time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			ExpectedOK:    true,
		},
		{
			Name:       "not a leap year",
			Input:      "2021-02-29",
			Parser:     YearMonthDay,
			ExpectedOK: false,
			WantErr:    true,
		},
	}
	RunTests(t, tests)
}

func TestDayOfWeek(t *testing.T) {
	tests := []ParserTest[time.Weekday]{
		{
			Name:          "monday long",
			Input:         CaPiTaLiZe("monday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Monday,
			ExpectedOK:    true,
		},
		{
			Name:          "monday short",
			Input:         CaPiTaLiZe("mon"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Monday,
			ExpectedOK:    true,
		},
		{
			Name:          "tuesday long",
			Input:         CaPiTaLiZe("tuesday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Tuesday,
			ExpectedOK:    true,
		},
		{
			Name:          "tuesday short",
			Input:         CaPiTaLiZe("tues"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Tuesday,
			ExpectedOK:    true,
		},
		{
			Name:          "tuesday shortest",
			Input:         CaPiTaLiZe("tue"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Tuesday,
			ExpectedOK:    true,
		},
		{
			Name:          "wednesday long",
			Input:         CaPiTaLiZe("wednesday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Wednesday,
			ExpectedOK:    true,
		},
		{
			Name:          "wednesday short",
			Input:         CaPiTaLiZe("weds"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Wednesday,
			ExpectedOK:    true,
		},
		{
			Name:          "wednesday shortest",
			Input:         CaPiTaLiZe("wed"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Wednesday,
			ExpectedOK:    true,
		},
		{
			Name:          "thursday long",
			Input:         CaPiTaLiZe("thursday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Thursday,
			ExpectedOK:    true,
		},
		{
			Name:          "thursday short",
			Input:         CaPiTaLiZe("thurs"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Thursday,
			ExpectedOK:    true,
		},
		{
			Name:          "thursday shorter",
			Input:         CaPiTaLiZe("thur"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Thursday,
			ExpectedOK:    true,
		},
		{
			Name:          "thursday shortest",
			Input:         CaPiTaLiZe("thu"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Thursday,
			ExpectedOK:    true,
		},
		{
			Name:          "friday long",
			Input:         CaPiTaLiZe("friday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Friday,
			ExpectedOK:    true,
		},
		{
			Name:          "friday short",
			Input:         CaPiTaLiZe("fri"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Friday,
			ExpectedOK:    true,
		},
		{
			Name:          "saturday long",
			Input:         CaPiTaLiZe("saturday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Saturday,
			ExpectedOK:    true,
		},
		{
			Name:          "saturday short",
			Input:         CaPiTaLiZe("sat"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Saturday,
			ExpectedOK:    true,
		},
		{
			Name:          "sunday long",
			Input:         CaPiTaLiZe("sunday"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Sunday,
			ExpectedOK:    true,
		},
		{
			Name:          "sunday short",
			Input:         CaPiTaLiZe("sun"),
			Parser:        DayOfWeek,
			ExpectedMatch: time.Sunday,
			ExpectedOK:    true,
		},
		{
			Name:           "no match",
			Input:          "foo",
			Parser:         DayOfWeek,
			ExpectedOK:     false,
			ExpectedMatch:  -1,
			RemainingInput: "foo",
		},
	}
	RunTests(t, tests)
}

func TestDaysOfWeek(t *testing.T) {
	tests := []ParserTest[[]time.Weekday]{
		{
			Name:          "monday",
			Input:         CaPiTaLiZe("monday"),
			Parser:        DaysOfWeek,
			ExpectedMatch: []time.Weekday{time.Monday},
			ExpectedOK:    true,
		},
		{
			Name:          "every day of the week",
			Input:         CaPiTaLiZe("mon tues wed thu fri sat sun"),
			Parser:        DaysOfWeek,
			ExpectedMatch: []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday},
			ExpectedOK:    true,
		},
		{
			Name:           "not a day",
			Input:          "not a day",
			Parser:         DaysOfWeek,
			ExpectedOK:     false,
			RemainingInput: "not a day",
		},
	}
	RunTests(t, tests)
}

func TestMonthDay(t *testing.T) {
	tests := []ParserTest[int]{
		{
			Name:          "1st",
			Input:         CaPiTaLiZe("1st"),
			Parser:        MonthDay,
			ExpectedMatch: 1,
			ExpectedOK:    true,
		},
		{
			Name:          "1",
			Input:         "1",
			Parser:        MonthDay,
			ExpectedMatch: 1,
			ExpectedOK:    true,
		},
		{
			Name:          "2nd",
			Input:         CaPiTaLiZe("2nd"),
			Parser:        MonthDay,
			ExpectedMatch: 2,
			ExpectedOK:    true,
		},
		{
			Name:          "3rd",
			Input:         CaPiTaLiZe("3rd"),
			Parser:        MonthDay,
			ExpectedMatch: 3,
			ExpectedOK:    true,
		},
		{
			Name:          "4th",
			Input:         CaPiTaLiZe("4th"),
			Parser:        MonthDay,
			ExpectedMatch: 4,
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestMonthOfYear(t *testing.T) {
	tests := []ParserTest[time.Month]{
		{
			Name:          "january",
			Input:         CaPiTaLiZe("january"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.January,
			ExpectedOK:    true,
		},
		{
			Name:          "jan",
			Input:         CaPiTaLiZe("jan"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.January,
			ExpectedOK:    true,
		},
		{
			Name:          "february",
			Input:         CaPiTaLiZe("february"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.February,
			ExpectedOK:    true,
		},
		{
			Name:          "feb",
			Input:         CaPiTaLiZe("feb"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.February,
			ExpectedOK:    true,
		},
		{
			Name:          "march",
			Input:         CaPiTaLiZe("march"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.March,
			ExpectedOK:    true,
		},
		{
			Name:          "mar",
			Input:         CaPiTaLiZe("mar"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.March,
			ExpectedOK:    true,
		},
		{
			Name:          "april",
			Input:         CaPiTaLiZe("april"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.April,
			ExpectedOK:    true,
		},
		{
			Name:          "apr",
			Input:         CaPiTaLiZe("apr"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.April,
			ExpectedOK:    true,
		},
		{
			Name:          "may",
			Input:         CaPiTaLiZe("may"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.May,
			ExpectedOK:    true,
		},
		{
			Name:          "june",
			Input:         CaPiTaLiZe("june"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.June,
			ExpectedOK:    true,
		},
		{
			Name:          "jun",
			Input:         CaPiTaLiZe("jun"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.June,
			ExpectedOK:    true,
		},
		{
			Name:          "july",
			Input:         CaPiTaLiZe("july"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.July,
			ExpectedOK:    true,
		},
		{
			Name:          "jul",
			Input:         CaPiTaLiZe("jul"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.July,
			ExpectedOK:    true,
		},
		{
			Name:          "august",
			Input:         CaPiTaLiZe("august"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.August,
			ExpectedOK:    true,
		},
		{
			Name:          "aug",
			Input:         CaPiTaLiZe("aug"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.August,
			ExpectedOK:    true,
		},
		{
			Name:          "september",
			Input:         CaPiTaLiZe("september"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.September,
			ExpectedOK:    true,
		},
		{
			Name:          "sept",
			Input:         CaPiTaLiZe("sept"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.September,
			ExpectedOK:    true,
		},
		{
			Name:          "sep",
			Input:         CaPiTaLiZe("sep"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.September,
			ExpectedOK:    true,
		},
		{
			Name:          "october",
			Input:         CaPiTaLiZe("october"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.October,
			ExpectedOK:    true,
		},
		{
			Name:          "oct",
			Input:         CaPiTaLiZe("oct"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.October,
			ExpectedOK:    true,
		},
		{
			Name:          "november",
			Input:         CaPiTaLiZe("november"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.November,
			ExpectedOK:    true,
		},
		{
			Name:          "nov",
			Input:         CaPiTaLiZe("nov"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.November,
			ExpectedOK:    true,
		},
		{
			Name:          "december",
			Input:         CaPiTaLiZe("december"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.December,
			ExpectedOK:    true,
		},
		{
			Name:          "dec",
			Input:         CaPiTaLiZe("dec"),
			Parser:        MonthOfYear,
			ExpectedMatch: time.December,
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}
