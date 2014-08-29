package itime

import "time"

var Now = func(format string) string {
	return time.Now().Format(format)
}
