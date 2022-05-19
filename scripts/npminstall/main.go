package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mlctrez/goapp-pf/scripts"
)

func main() {
	defer scripts.Recover()

	tempNode := scripts.NodeTempDir()

	if _, errStat := os.Stat(filepath.Join(tempNode, "node_modules")); os.IsNotExist(errStat) {
		toInstall := writeEmbedFiles(tempNode)

		npmArgs := []string{"install", "--save"}
		npmArgs = append(npmArgs, toInstall...)
		command := exec.Command("npm", npmArgs...)

		command.Dir = tempNode
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
			scripts.NoErr(err)
		}

	}

}

func writeEmbedFiles(tempNode string) (installNames []string) {

	noErr := scripts.NoErr

	var err error
	var readDir []fs.DirEntry
	readDir, err = nodeJson.ReadDir(".")
	noErr(err)

	for _, entry := range readDir {
		var data []byte
		data, err = nodeJson.ReadFile(entry.Name())
		noErr(err)
		if entry.Name() == "package.json" {
			p := &packageJson{}
			noErr(json.Unmarshal(data, p))
			for k := range p.Dependencies {
				installNames = append(installNames, k)
			}
		}
		err = os.WriteFile(filepath.Join(tempNode, entry.Name()), data, 0644)
		noErr(err)
	}
	return installNames
}

type packageJson struct {
	Dependencies map[string]string `json:"dependencies"`
}

//go:embed package.json package-lock.json
var nodeJson embed.FS
