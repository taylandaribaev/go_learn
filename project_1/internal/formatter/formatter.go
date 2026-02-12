package formatter

import (
	"fmt"
)

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%d bytes", size)
	}
	var unit string
	var value float64
	switch {
	case size >= 1024*1024*1024:
		unit = "GB"
		value = float64(size) / (1024 * 1024 * 1024)
	case size >= 1024*1024:
		unit = "MB"
		value = float64(size) / (1024 * 1024)
	case size >= 1024:
		unit = "KB"
		value = float64(size) / 1024
	default:
		unit = "bytes"
		value = float64(size)
	}
	return fmt.Sprintf("%.2f %s", value, unit)
}
