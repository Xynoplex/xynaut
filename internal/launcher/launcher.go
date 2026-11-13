// Package launcher implements xynaut commands for managing Semiont stacks.
package launcher

// These defaults identify an unstamped build; release builds replace them
// with -ldflags.
const (
	developmentVersion = "dev"
	unstampedCommit  = "none"
	unstampedDate    = "unknown"
)

// Set via -ldflags at release time.
var (
	BuildVersion = "0.9.0"
	BuildCommit  = unstampedCommit
	BuildDate    = unstampedDate
)
