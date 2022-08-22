package clock

const (
	Layout      = "%m/%d %h:%m:%s%p '%y %z"
	ANSIC       = "%a %b %d %h:%m:%s %Y"
	UnixDate    = "%a %b %d %h:%m:%s %Z %Y"
	RubyDate    = "%a %b %d %h:%m:%s %z %Y"
	RFC822      = "%d %b %y %h:%m %Z"
	RFC822Z     = "%d %b %y %h:%m %z"
	RFC850      = "%A, %d-%b-%y %h:%m:%s %Z"
	RFC1123     = "%a, %d %b %Y %h:%m:%s %Z"
	RFC1123Z    = "%a, %d %b %Y %h:%m:%s %z"
	RFC3339     = "%Y-%m-%dT%h:%m:%dZ%o"
	RFC3339Nano = "%Y-%m-%dT%h:%m:%d.%nZ%o"
	Kitchen     = "%h:%m%p"
	Stamp       = "%b %d %h:%m:%s"
	StampMilli  = "b %d %h:%m:%s.%i"
	StampMicro  = "b %d %h:%m:%s.%u"
	StampNano   = "b %d %h:%m:%s.%n"
)
