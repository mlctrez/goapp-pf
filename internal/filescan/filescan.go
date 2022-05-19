package filescan

import (
	"os"
	"path/filepath"
	"sort"
)

func SubDirectories(dir string) (result []string, err error) {

	var files []string
	if files, err = filepath.Glob(filepath.Join(dir, "*")); err != nil {
		panic(err)
	}
	for _, s := range files {
		var fi os.FileInfo
		if fi, err = os.Stat(s); err != nil {
			return
		}
		if fi.IsDir() {
			result = append(result, s)
		}
	}
	sort.Strings(result)
	return
}
