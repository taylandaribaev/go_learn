package scanner

import (
	"dirsize/internal/formatter"

	"github.com/fatih/color"
)

type File struct {
	Name string
	Path string
	Size int64
}

func (f *File) Print(depth int, prefix string, human bool) {
	fileColor := color.New()

	r, g, b := colorByDepth(depth)
	fileColor = color.RGB(r, g, b)

	fileColor.Printf("%s  📄 %s (%s)\n", prefix, f.Name, formatter.FormatSizeExternal(f.Size, human))
}

func (f *File) GetSize() int64 {
	return f.Size
}
