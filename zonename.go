package main

import (
	"strings"
)

type (
	// ZoneName is a zone name without the traling dot.
	ZoneName string
)

// NewZoneName will return a new ZoneName. If zoneName contains a trailing dot,
// it will be removed.
func NewZoneName(zoneName string) ZoneName {
	return ZoneName(strings.Trim(zoneName, "."))
}
