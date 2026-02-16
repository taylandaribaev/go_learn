package formatter

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func FormatSizeExternal(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%d bytes", size)
	}
	return humanize.Bytes(uint64(size))
}
