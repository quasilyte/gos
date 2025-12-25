package gos

import "path/filepath"

func AbsPath(p string) string {
	abs, err := filepath.Abs(p)
	if err != nil {
		return ""
	}
	return filepath.ToSlash(abs)
}
