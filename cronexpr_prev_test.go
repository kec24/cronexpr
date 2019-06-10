package cronexpr

/******************************************************************************/

import (
	"testing"
	"time"
)

/******************************************************************************/

var cronprevtests = []crontest{

	// Seconds
	{
		"* * * * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:01", "2013-01-01 00:00:00" },
			{"2013-01-01 00:01:00", "2013-01-01 00:00:59"},
			{"2013-01-01 01:00:00", "2013-01-01 00:59:59" },
			{"2013-01-02 00:00:00", "2013-01-01 23:59:59"},
			{"2013-03-01 00:00:00", "2013-02-28 23:59:59"},
			{"2016-02-29 00:00:00", "2016-02-28 23:59:59"},
			{"2013-01-01 00:00:00", "2012-12-31 23:59:59"},
		},
	},

	// every 5 Second
	{
		"*/5 * * * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:05", "2013-01-01 00:00:00"},
			{"2013-01-01 00:00:01", "2013-01-01 00:00:00"},
			{"2013-01-01 00:59:59", "2013-01-01 00:59:55"},
		},
	},

	// Minutes
	{
		"* * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2012-12-31 23:59:00"},
			{"2013-01-01 00:01:59", "2013-01-01 00:01:00"},
			{"2013-01-01 00:00:03", "2013-01-01 00:00:00"},
			{"2013-03-01 00:00:00", "2013-02-28 23:59:00"},
			{"2013-02-28 00:00:00", "2013-02-27 23:59:00"},
			{"2016-03-01 00:00:00", "2016-02-29 23:59:00"},
			{"2012-01-01 17:00:00", "2012-01-01 16:59:00"},
		},
	},

	// Minutes with interval
	{
		"17-43/5 * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:00:00", "2012-12-31 23:42:00"},
			{"2013-01-01 00:16:59", "2012-12-31 23:42:00"},
			{"2013-01-01 00:30:00", "2013-01-01 00:27:00"},
			{"2013-01-01 00:12:00", "2012-12-31 23:42:00"},
			{"2013-01-01 23:10:00", "2013-01-01 22:42:00"},
			{"2013-03-01 00:08:00", "2013-02-28 23:42:00"},
			{"2016-03-01 00:05:00", "2016-02-29 23:42:00"},
		},
	},

	// Minutes interval, list
	{
		"15-30/4,55 * * * *",
		"2006-01-02 15:04:05",
		[]crontimes{
			{"2013-01-01 00:59:59", "2013-01-01 00:55:00"},
			{"2013-01-01 00:54:00", "2013-01-01 00:27:00"},
			{"2013-01-01 00:59:00", "2013-01-01 00:55:00"},
			{"2013-01-01 00:55:00", "2013-01-01 00:27:00"},
			{"2013-01-02 00:15:00", "2013-01-01 23:55:00"},
			{"2013-03-01 00:00:00", "2013-02-28 23:55:00"},
			{"2016-03-01 00:00:05", "2016-02-29 23:55:00"},
			{"2013-01-01 00:00:09", "2012-12-31 23:55:00"},
			{"2013-01-01 00:15:00", "2012-12-31 23:55:00"},
		},
	},

	// Days of week
	{
		"0 0 * * MON",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-01 00:00:00", "Mon 2012-12-31 00:00"},
			{"2013-01-28 00:00:00", "Mon 2013-01-21 00:00"},
			{"2013-12-30 00:30:00", "Mon 2013-12-30 00:00"},
		},
	},
	{
		"0 0 * * friday",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-01 00:00:00", "Fri 2012-12-28 00:00"},
			{"2013-02-01 00:00:00", "Fri 2013-01-25 00:00"},
			{"2014-01-02 00:30:00", "Fri 2013-12-27 00:00"},
		},
	},
	{
		"0 0 * * 6,7",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-01-01 00:00:00", "Sun 2012-12-30 00:00"},
		},
	},

	// Specific days of week
	{
		"0 0 * * 6#5",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-08-22 00:00:00", "Sat 2013-06-29 00:00"}, // july doesn't have a 5th saturday
		},
	},

	// Work day of month
	{
		"0 0 6W * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-08-05 00:00:00", "Fri 2013-07-05 00:00"},
			{"2013-11-05 00:00:00", "Mon 2013-10-07 00:00"},
		},
	},

	// Work day of month -- end of month
	{
		"0 0 30W * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-03-02 00:00:00", "Wed 2013-01-30 00:00"},
			{"2013-06-02 00:00:00", "Thu 2013-05-30 00:00"},
			{"2013-03-02 00:00:00", "Wed 2013-01-30 00:00"},
		},
	},

	// Last day of month
	{
		"0 0 L * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-09-02 00:00:00", "Sat 2013-08-31 00:00"},
			{"2014-01-01 00:00:00", "Tue 2013-12-31 00:00"},
			{"2014-03-01 00:00:00", "Fri 2014-02-28 00:00"},
			{"2016-03-15 00:00:00", "Mon 2016-02-29 00:00"},
		},
	},

	// Last work day of month
	{
		"0 0 LW * *",
		"Mon 2006-01-02 15:04",
		[]crontimes{
			{"2013-09-02 00:00:00", "Fri 2013-08-30 00:00"},
			{"2013-11-02 00:00:00", "Thu 2013-10-31 00:00"},
			{"2013-07-15 00:00:00", "Fri 2013-06-28 00:00"},
		},
	},

}

func TestPrevExpressions(t *testing.T) {
	for _, test := range cronprevtests {
		for _, times := range test.times {
			from, _ := time.Parse("2006-01-02 15:04:05", times.from)
			expr, err := Parse(test.expr)
			if err != nil {
				t.Errorf(`Parse("%s") returned "%s"`, test.expr, err.Error())
			}
			prev := expr.Prev(from)
			prevstr := prev.Format(test.layout)
			if prevstr != times.next {
				t.Errorf(`("%s").Prev("%s") = "%s", got "%s"`, test.expr, times.from, times.next, prevstr)
			}
		}
	}
}
