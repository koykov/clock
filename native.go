package clock

import "time"

// Native is a wrapper over default time.Now() function implements Interface.
type Native struct{}

func (c Native) Now() time.Time {
	return time.Now()
}
