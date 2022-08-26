# Clock

Fast analogue of time package methods.

## Benchmarks
```
BenchmarkClock/clock.Now()-8         	1000000000          0.8397 ns/op
BenchmarkClock/time.Now()-8          	20795817	        57.93 ns/op
```

## Format

| pattern | description                                                                             |
|:--------|:----------------------------------------------------------------------------------------|
| %A      | full weekday name                                                                       |
| %a      | abbreviated weekday name                                                                |
| %B      | full month name                                                                         |
| %b      | abbreviated month name                                                                  |
| %C      | two digit representation of the century                                                 |
| %c      | equivalent to %a %b %e %H:%M:%S %Y                                                      |
| %D      | equivalent to %m/%d/%y                                                                  |
| %d      | day of the month with leading zero (01-31)                                              |
| %e      | day of the month with leading space (1-31)                                              |
| %F      | equivalent to %Y-%m-%d                                                                  |
| %H      | the hour (24-hour clock) as a decimal number [00-23]                                    |
| %h      | alias of %b                                                                             |
| %I      | hour (12-hour clock) as a decimal number [01-12]                                        |
| %j      | day of the year with leading zero [001-366]                                             |
| %k      | hour (24-hour clock) with leading space [0-23]                                          |
| %l      | hour (12-hour clock) with leading space [1-12]                                          |
| %M      | minute with leading zero [00-59]                                                        |
| %m      | month with leading zero [01-12]                                                         |
| %N      | nanoseconds (9 digits)                                                                  |
| %n      | nanoseconds (7 digits)                                                                  |
| %o      | microseconds (6 digits)                                                                 |
| %i      | milliseconds (3 digits)                                                                 |
| %P      | am/pm                                                                                   |
| %p      | AM/PM                                                                                   |
| %R      | equivalent to %H:%M                                                                     |
| %r      | equivalent to %I:%M:%S %p                                                               |
| %S      | second with leading zero [00-60]                                                        |
| %T      | equivalent to %H:%M:%S                                                                  |
| %U      | week number of the year (Sunday as the first day of the week) with leading zero [00-53] |
| %u      | weekday (Monday as the first day of the week) [1-7]                                     |
| %V      | week number of the year (Monday as the first day of the week) with leading zero [01-53] |
| %v      | equivalent to %e-%b-%Y                                                                  |
| %W      | week number of the year (Monday as the first day of the week) with leading zero [00-53] |
| %w      | the weekday (Sunday as the first day of the week) as a decimal number [0-6]             |
| %X      | equivalent to %H:%M:%S                                                                  |
| %x      | equivalent to %m/%d/%y                                                                  |
| %Y      | year with century as a decimal number                                                   |
| %y      | year without century as a decimal number [00-99]                                        |
| %Z      | time zone name                                                                          |
| %z      | time zone offset from UTC                                                               |
| %%      | symbol '%'                                                                              |
