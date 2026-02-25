package scanner

import (
	"dirsize/internal/formatter"

	"github.com/fatih/color"
)

type Dir struct {
	Name  string
	Path  string
	Files []PrintSizer
	Dirs  []PrintSizerWithRecalc
	Size  int64
	depth int
}

func (d *Dir) GetSize() int64 {
	return d.Size
}

func (d *Dir) GetDepth() int {
	return d.depth
}

func (d *Dir) SetDepth(depth int) {
	d.depth = depth
}

func (d *Dir) AddFile(file PrintSizer) {
	d.Files = append(d.Files, file)
	d.Size += file.GetSize()
}

func (d *Dir) AddDir(dir PrintSizerWithRecalc) {
	d.Dirs = append(d.Dirs, dir)
}

func (d *Dir) RecalcSize() int64 {
	var total int64

	for _, f := range d.Files {
		total += f.GetSize()
	}
	for _, sub := range d.Dirs {
		total += sub.RecalcSize()
	}

	d.Size = total
	return total
}

func (d *Dir) Print(depth int, prefix string, human bool) {
	dirColor := color.New()

	r, g, b := colorByDepth(depth)
	dirColor = color.RGB(r, g, b)

	dirColor.Printf("%s📁 %s (%s)\n", prefix, d.Name, formatter.FormatSizeExternal(d.Size, human))
}
