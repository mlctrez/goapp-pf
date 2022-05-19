package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mlctrez/bolt"
	"github.com/mlctrez/goapp-pf/scripts"
)

var tsInspect = "tsInspect.ts"

func main() {
	defer scripts.Recover()
	noErr := scripts.NoErr

	nodeDir := scripts.NodeTempDir()

	copyTypeScript(nodeDir, tsInspect)

	db := scripts.Bolt(scripts.InspectDatabase)
	defer func() { scripts.BoltClose(db) }()

	componentBucket := bolt.Key("components")
	layoutBucket := bolt.Key("layouts")
	keys := bolt.Keys{componentBucket, layoutBucket}
	noErr(db.CreateBuckets(keys))

	noErr(os.Chdir("temp/node"))

	paths := make(map[bolt.Key][]string)

	for _, key := range keys {
		path := filepath.Join("node_modules/@patternfly/react-core/src", string(key))
		err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if info != nil && strings.HasSuffix(info.Name(), ".tsx") {
				if strings.Contains(path, "__tests__") || strings.Contains(path, "examples") {
					return nil
				}
				if strings.HasSuffix(path, "index.tsx") {
					return nil
				}
				paths[key] = append(paths[key], path)
			}
			return nil
		})
		noErr(err)
	}

	for key, pathList := range paths {

		var savePaths []string
		for _, path := range pathList {
			v := &bolt.Value{K: bolt.Key(path)}
			if err := db.Get(key, v); err == bolt.ErrValueNotFound {
				savePaths = append(savePaths, path)
			}
			if len(savePaths) > 9 {
				write(db, key, savePaths...)
				savePaths = []string{}
			}
		}
		if len(savePaths) > 9 {
			write(db, key, savePaths...)
		}

	}

}

func write(db *bolt.Bolt, key bolt.Key, savePaths ...string) {
	var data []byte
	var err error
	noErr := scripts.NoErr

	args := []string{tsInspect}
	args = append(args, savePaths...)
	fmt.Println(args)
	data, err = exec.Command("node", args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(data))
		noErr(err)
	}
	m := make(map[string]interface{})
	noErr(json.Unmarshal(data, &m))
	for p, dataJson := range m {
		data, err = json.Marshal(dataJson)
		noErr(err)
		v := &bolt.Value{K: bolt.Key(p), V: data}
		noErr(db.Put(key, v))
	}

}

//go:embed *.txt
var tsFiles embed.FS

func copyTypeScript(dir, name string) {
	data, err := tsFiles.ReadFile(name + ".txt")
	scripts.NoErr(err)

	err = os.WriteFile(filepath.Join(dir, name), data, 0644)
	scripts.NoErr(err)
}
