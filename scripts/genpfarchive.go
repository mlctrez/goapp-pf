//go:build ignore

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	dir := filepath.Join("temp/node_pf")

	failOnErr(os.MkdirAll(dir, 0755))

	failOnErr(os.WriteFile(filepath.Join(dir, "package.json"), []byte(packageJson), 0755))
	failOnErr(os.WriteFile(filepath.Join(dir, "package-lock.json"), []byte(packageLockJson), 0755))

	if _, errStat := os.Stat(filepath.Join(dir, "node_modules")); os.IsNotExist(errStat) {
		command := exec.Command("npm", "install", "@patternfly/patternfly", "--save")
		command.Dir = dir
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
		}
		failOnErr(err)
	}

	zipPath := filepath.Join(dir, "patternfly.zip")
	if _, zipErr := os.Stat(zipPath); os.IsNotExist(zipErr) {

		zipFile, openErr := os.Create(zipPath)
		failOnErr(openErr)

		writer := zip.NewWriter(zipFile)

		prefix := filepath.Join(dir, "node_modules", "@patternfly", "patternfly")
		failOnErr(
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
		failOnErr(writer.Flush())
		failOnErr(writer.Close())

	}

}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageJson = `{
  "name": "goapp-pf",
  "version": "1.0.0",
  "description": "temp package json",
  "repository": {},
  "license": "Apache-2.0",
  "dependencies": {
    "@patternfly/patternfly": "^4.192.1"
  }
}
`
var packageLockJson = `{
  "name": "goapp-pf",
  "version": "1.0.0",
  "lockfileVersion": 1,
  "requires": true,
  "dependencies": {
    "@patternfly/patternfly": {
      "version": "4.192.1",
      "resolved": "https://registry.npmjs.org/@patternfly/patternfly/-/patternfly-4.192.1.tgz",
      "integrity": "sha512-eNJ3aI9mGfvwMtBwkI+CBJHPhZx1FoNN6QY36iYEvrEOIL5xuuKRDG2tbOzeucQOzNqZ1PO1Eoock5xTcCG86Q=="
    }
  }
}
`
