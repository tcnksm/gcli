package skeleton

import (
	"fmt"
)

// Debugf outputs debug infomation
func (s *Skeleton) Debugf(format string, args ...interface{}) {
	if s.Verbose {
		fmt.Fprintf(s.LogWriter, "[DEBUG] "+format+"\n", args...)
	}
}
