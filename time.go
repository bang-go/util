package util

import "time"

// GetLocation attempts to load a location by name.
// If it fails (e.g. missing tzdata), it constructs a fixed zone with the given offset in seconds.
// Example: GetLocation("Asia/Shanghai", 8*3600)
func GetLocation(name string, offsetSec int) *time.Location {
	location, err := time.LoadLocation(name)
	if err != nil {
		return time.FixedZone(name, offsetSec)
	}
	return location
}
