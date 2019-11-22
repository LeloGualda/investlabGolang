package querys

import (
	"time"
)

func newDate() string {
	start := time.Now()
	date := start.Format("2006-01-02")

	return date
}
