package scan

import (
	"time"
)

// Options of the scan
type Options struct {
	Timeout time.Duration
	Retries int
	Rate    int
	Debug   bool
	Proxy   string
	Stream  bool
}
