package util

import "time"

// LoadLocationOrFixed loads a location by name.
// If it fails, it returns a fixed zone with the provided offset in seconds.
func LoadLocationOrFixed(name string, offsetSec int) *time.Location {
	location, err := time.LoadLocation(name)
	if err != nil {
		return time.FixedZone(name, offsetSec)
	}
	return location
}
