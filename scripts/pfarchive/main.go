package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlctrez/goapp-pf/scripts"
)

func main() {
	defer scripts.Recover()
	noErr := scripts.NoErr

	dir := scripts.NodeTempDir()

	zipPath := filepath.Join(dir, "patternfly.zip")
	if _, zipErr := os.Stat(zipPath); os.IsNotExist(zipErr) {

		zipFile, openErr := os.Create(zipPath)
		noErr(openErr)

		writer := zip.NewWriter(zipFile)

		prefix := filepath.Join(dir, "node_modules", "@patternfly", "patternfly")
		noErr(
			filepath.Walk(prefix, func(path string, info fs.FileInfo, err error) error {
				if info == nil {
					return nil
				}
				name := info.Name()
				if info.IsDir() ||
					strings.HasSuffix(name, ".scss") ||
					strings.HasSuffix(name, ".md") {
					return nil
				}

				newPath := strings.TrimPrefix(path, prefix+"/")
				create, createErr := writer.Create(newPath)
				if createErr != nil {
					return createErr
				}
				src, srcErr := os.Open(path)
				if srcErr != nil {
					return srcErr
				}
				_, copyErr := io.Copy(create, src)
				if copyErr != nil {
					return copyErr
				}
				return src.Close()

			}),
		)
		noErr(writer.Flush())
		noErr(writer.Close())

	}

}
