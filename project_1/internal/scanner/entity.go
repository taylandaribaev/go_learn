package scanner

import (
	"strings"
)

type PrintSizer interface {
	Print(depth int, prefix string, human bool)
	GetSize() int64
}

type PrintSizerWithRecalc interface {
	PrintSizer
	GetDepth() int
	SetDepth(depth int)
	AddFile(file PrintSizer)
	AddDir(dir PrintSizerWithRecalc)
	RecalcSize() int64
	PrintTree(human bool, fullTree bool)
}

func colorByDepth(depth int) (int, int, int) {
	// базовый светлый цвет
	baseR := 120
	baseG := 180
	baseB := 255

	// затемнение
	factor := depth * 20

	r := max(baseR-factor, 0)
	g := max(baseG-factor, 0)
	b := max(baseB-factor, 0)

	return r, g, b
}

func (d *Dir) PrintTree(human bool, fullTree bool) {
	prefix := strings.Repeat("  ", d.GetDepth())

	d.Print(d.GetDepth(), prefix, human)
	for _, file := range d.Files {
		file.Print(d.GetDepth(), prefix, human)
	}

	for _, childDir := range d.Dirs {
		if !fullTree {
			childDir.Print(d.GetDepth(), prefix+"  ", human)
			continue
		}
		childDir.SetDepth(d.GetDepth() + 1)
		childDir.PrintTree(human, fullTree)
	}
}
