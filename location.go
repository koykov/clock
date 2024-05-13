package clock

import (
	"sync"
	"time"
)

var (
	locMux      sync.RWMutex
	locRegistry map[string]*time.Location
)

func init() {
	locMux.Lock()
	locRegistry = make(map[string]*time.Location)
	locMux.Unlock()
}

// LoadLocation returns Location for given name.
func LoadLocation(name string) (*time.Location, error) {
	locMux.RLock()
	loc, ok := locRegistry[name]
	locMux.RUnlock()
	if ok {
		return loc, nil
	}
	var err error
	loc, err = time.LoadLocation(name)
	if err != nil {
		return nil, err
	}
	locMux.Lock()
	locRegistry[name] = loc
	locMux.Unlock()
	return loc, nil
}

var _ = LoadLocation
