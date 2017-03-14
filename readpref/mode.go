package readpref

import (
	"fmt"
	"strings"
)

// Mode indicates the user's preference on reads.
type Mode uint8

// Mode constants
const (
	// PrimaryMode indicates that only a primary is
	// considered for reading. This is the default
	// mode.
	PrimaryMode Mode = iota
	// PrimaryPreferredMode indicates that if a primary
	// is available, use it; otherwise, eligible
	// secondaries will be considered.
	PrimaryPreferredMode
	// SecondaryMode indicates that only secondaries
	// should be considered.
	SecondaryMode
	// SecondaryPreferredMode indicates that only secondaries
	// should be considered when one is available. If none
	// are available, then a primary will be considered.
	SecondaryPreferredMode
	// NearestMode indicates that all primaries and secondaries
	// will be considered.
	NearestMode
)

// ModeFromString returns a mode corresponding to
// mode.
func ModeFromString(mode string) (Mode, error) {
	switch strings.ToLower(mode) {
	case "primary":
		return PrimaryMode, nil
	case "primarypreferred":
		return PrimaryPreferredMode, nil
	case "secondary":
		return SecondaryMode, nil
	case "secondarypreferred":
		return SecondaryPreferredMode, nil
	case "nearest":
		return NearestMode, nil
	}
	return Mode(uint8(0)), fmt.Errorf("unknown read preference %v", mode)
}
