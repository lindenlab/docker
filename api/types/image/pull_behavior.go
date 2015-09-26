package image

import (
	"fmt"
)

// PullBehavior can be one of: never, always, or missing
type PullBehavior int

const (
	PullNever PullBehavior = iota
	PullAlways
	PullMissing
)

func ParsePullBehavior(pullVal string) (PullBehavior, error) {
	switch pullVal {
	case "never":
		return PullNever, nil
	case "always":
		return PullAlways, nil
	case "missing", "":
		return PullMissing, nil
	}
	return PullNever, fmt.Errorf("Invalid pull behavior '%s'", pullVal)
}

func (p PullBehavior) String() string {
	switch p {
	case PullNever:
		return "never"
	case PullAlways:
		return "always"
	case PullMissing:
		return "missing"
	}
	return ""
}
