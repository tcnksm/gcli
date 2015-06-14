package skeleton

import "time"

// dateFunc returns formated date. Format follows
// https://github.com/olivierlacan/keep-a-changelog format
func dateFunc() func() string {
	return func() string {
		return time.Now().Format("2006-01-02")
	}
}
