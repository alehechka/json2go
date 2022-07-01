package utils

import "time"

// TimeFormatMap map of time formats
var TimeFormatMap = map[string]string{
	"Layout":      time.Layout,
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
}

// GetTimeFormat attempts to find a standard time format, else will return the fallback on empty strings or unaltered.
func GetTimeFormat(format string, fallback string) string {
	if len(format) == 0 {
		return fallback
	}

	if standard, exists := TimeFormatMap[format]; exists {
		return standard
	}

	return format
}
