package scanner

import (
	"fmt"
	"strings"

	"dirsize/internal/formatter"
)

type File struct {
	Name string
	Path string
	Size int64
}

type Dir struct {
	Name  string
	Path  string
	Files []File
	Dirs  []*Dir
	Size  int64
}

func (d *Dir) AddFile(file File) {
	d.Files = append(d.Files, file)
	d.Size += file.Size
}

func (d *Dir) AddDir(dir *Dir) {
	d.Dirs = append(d.Dirs, dir)
}

func (d *Dir) RecalcSize() int64 {
	var total int64

	for _, f := range d.Files {
		total += f.Size
	}
	for _, sub := range d.Dirs {
		total += sub.RecalcSize()
	}

	d.Size = total
	return total
}

func (d *Dir) Print(indent int, human bool, fullTree bool) {
	prefix := strings.Repeat("  ", indent)

	fmt.Printf("%s📁 %s (%s)\n", prefix, d.Name, formatter.FormatSizeExternal(d.Size, human))

	for _, file := range d.Files {
		fmt.Printf("%s  📄 %s (%s)\n", prefix, file.Name, formatter.FormatSizeExternal(file.Size, human))
	}

	for _, dir := range d.Dirs {
		if !fullTree {
			fmt.Printf("%s  📁 %s (%s)\n", prefix, dir.Name, formatter.FormatSizeExternal(dir.Size, human))
			continue
		}
		dir.Print(indent+1, human, fullTree)
	}
}
