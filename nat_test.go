package clock

import "testing"

func TestRelativeNatural(t *testing.T) {
	spans := []struct {
		key, exp string
	}{
		// now
		{"now", "..."},
		{"right now", "..."},
		{"  right  now  ", "..."},

		// seconds
		{"1 second", "..."},
		{"next second", "..."},
		{"last second", "..."},
		{"one second", "..."},
		{"1 second ago", "..."},
		{"5 seconds ago", "..."},
		{"five seconds ago", "..."},
		{"   5    seconds  ago   ", "..."},
		{"2 seconds from now", "..."},
		{"two seconds from now", "..."},
		{"Message me in 2 seconds", "..."},
		{"Message me in 2 seconds from now", "..."},

		// minutes
		{"1 minute", "..."},
		{"next minute", "..."},
		{"last minute", "..."},
		{"one minute", "..."},
		{"1 minute ago", "..."},
		{"5 minutes ago", "..."},
		{"five minutes ago", "..."},
		{"   5    minutes  ago   ", "..."},
		{"2 minutes from now", "..."},
		{"two minutes from now", "..."},
		{"Message me in 2 minutes", "..."},
		{"Message me in 2 minutes from now", "..."},

		// hours
		{"1 hour", "..."},
		{"last hour", "..."},
		{"next hour", "..."},
		{"1 hour ago", "..."},
		{"6 hours ago", "..."},
		{"an hour ago", "..."},
		{"twelve hours ago", "..."},
		{"1 hour from now", "..."},
		{"Remind me in 1 hour", "..."},
		{"Remind me in 1 hour from now", "..."},
		{"Remind me in 1 hour and 3 minutes from now", "..."},
		{"Remind me in an hour", "..."},
		{"Remind me in an hour from now", "..."},

		// days
		{"1 day", "..."},
		{"next day", "..."},
		{"1 day ago", "..."},
		{"3 days ago", "..."},
		{"3 days ago at 11:25am", "..."},
		{"1 day from now", "..."},
		{"Remind me one day from now", "..."},
		{"Remind me in a day", "..."},
		{"Remind me in one day", "..."},
		{"Remind me in one day from now", "..."},

		// weeks
		{"1 week", "..."},
		{"1 week ago", "..."},
		{"2 weeks ago", "..."},
		{"2 weeks ago at 8am", "..."},
		{"next week", "..."},
		{"Message me in a week", "..."},
		{"Message me in one week", "..."},
		{"Message me in one week from now", "..."},
		{"Message me in two weeks from now", "..."},
		{"Message me two weeks from now", "..."},
		{"Message me in two weeks", "..."},

		// months
		{"1 month ago", "..."},
		{"a month ago", "..."},
		{"eleven months ago", "..."},
		{"a month ago", "..."},
		{"last month", "..."},
		{"next month", "..."},
		{"1 month ago at 9:30am", "..."},
		{"2 months ago", "..."},
		{"12 months ago", "..."},
		{"1 month from now", "..."},
		{"next 2 months", "..."},
		{"2 months from now", "..."},
		{"12 months from now at 6am", "..."},
		{"Remind me in 12 months from now at 6am", "..."},
		{"Remind me in a month", "..."},
		{"Remind me in 2 months", "..."},
		{"Remind me in a month from now", "..."},
		{"Remind me in 2 months from now", "..."},

		// years
		{"last year", "..."},
		{"next year", "..."},
		{"one year ago", "..."},
		{"one year from now", "..."},
		{"two years ago", "..."},
		{"2 years ago", "..."},
		{"Remind me in one year from now", "..."},
		{"Remind me in a year", "..."},
		{"Remind me in a year from now", "..."},

		// today
		{"today", "..."},
		{"today at 10am", "..."},

		// yesterday
		{"yesterday", "..."},
		{"yesterday 10am", "..."},
		{"yesterday at 10am", "..."},
		{"yesterday at 10:15am", "..."},

		// tomorrow
		{"tomorrow", "..."},
		{"tomorrow 10am", "..."},
		{"tomorrow at 10am", "..."},
		{"tomorrow at 10:15am", "..."},

		// past weekdays
		{"sunday", "..."},
		{"monday", "..."},
		{"tuesday", "..."},
		{"wednesday", "..."},
		{"thursday", "..."},
		{"friday", "..."},
		{"saturday", "..."},

		{"last sunday", "..."},
		{"past sunday", "..."},
		{"last monday", "..."},
		{"last tuesday", "..."},
		{"last wednesday", "..."},
		{"last thursday", "..."},
		{"last friday", "..."},
		{"last saturday", "..."},

		// future weekdays
		{"next tuesday", "..."},
		{"next wednesday", "..."},
		{"next thursday", "..."},
		{"next friday", "..."},
		{"next saturday", "..."},
		{"next sunday", "..."},
		{"next monday", "..."},

		// months
		{"last january", "..."},
		{"next january", "..."},
		{"january", "..."},
		{"february", "..."},
		{"march", "..."},
		{"april", "..."},
		{"may", "..."},
		{"june", "..."},
		{"july", "..."},
		{"august", "..."},
		{"september", "..."},
		{"october", "..."},
		{"november", "..."},

		// ordinal dates
		{"november 15th", "..."},
		{"december 1st", "..."},
		{"december 2nd", "..."},
		{"december 3rd", "..."},
		{"december 4th", "..."},
		{"december 15th", "..."},
		{"december 23rd", "..."},
		{"december 23rd 5pm", "..."},
		{"december 23rd at 5pm", "..."},
		{"december 23rd at 5:25pm", "..."},

		// 12-hour clock
		{"10am", "..."},
		{"10 am", "..."},
		{"5pm", "..."},
		{"10:25am", "..."},
		{"1:05pm", "..."},
		{"10:25:10am", "..."},
		{"1:05:10pm", "..."},

		// 24-hour clock
		{"10", "..."},
		{"10:25", "..."},
		{"10:25:30", "..."},
		{"17", "..."},
		{"17:25:30", "..."},

		// case sensitivity
		{"December 23rd AT 5:25 PM", "..."},
		{"next December 23rd AT 5:25 PM", "..."},

		// QA
		{"Restart the server in 2 days from now", "..."},
		{"Remind me on the 5th of next month", "..."},
		{"Remind me on the 5th of next month at 7am", "..."},
		{"Remind me at 7am on the 5th of next month", "..."},
		{"Remind me in one month from now", "..."},
		{"Remind me in one month from now at 7am", "..."},
	}
	for _, span := range spans {
		_ = span
		// ...
	}
}
