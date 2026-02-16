package scanner

import (
	"os"
	"path/filepath"
)

type Scanner struct {
	path string
}

func NewScanner(path string) *Scanner {
	cleanPath := filepath.Clean(path)
	return &Scanner{path: cleanPath}
}

func (s *Scanner) Scan() (*Dir, error) {
	root := &Dir{
		Name: filepath.Base(s.path),
		Path: s.path,
	}

	dirMap := map[string]*Dir{
		s.path: root,
	}

	err := filepath.WalkDir(s.path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path == s.path {
				return nil
			}

			parentPath := filepath.Dir(path)
			parent, ok := dirMap[parentPath]

			if !ok {
				return nil
			}

			newDir := &Dir{
				Name: d.Name(),
				Path: path,
			}

			parent.AddDir(newDir)
			dirMap[path] = newDir
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		parentPath := filepath.Dir(path)
		parent := dirMap[parentPath]

		file := File{
			Name: d.Name(),
			Path: path,
			Size: info.Size(),
		}

		parent.AddFile(file)
		return nil
	})

	if err != nil {
		return nil, err
	}
	root.RecalcSize()
	return root, nil
}
