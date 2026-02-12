package scanner

import (
	"fmt"
	"os"
	"path/filepath"
)

type Scanner struct {
	Path string
}

func NewScanner(path string) *Scanner {
	return &Scanner{Path: path}
}

func (s *Scanner) Scan() (int64, error) {
	var totalSize int64

	err := filepath.WalkDir(s.Path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking path %s: %w", path, err)
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return fmt.Errorf("Error getting info for path %s: %w", path, err)
			}
			totalSize += info.Size()
		}
		return nil
	})

	return totalSize, err
}
