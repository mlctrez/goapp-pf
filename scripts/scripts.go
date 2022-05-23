package scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mlctrez/bolt"
)

func NodeTempDir() string {
	dir := filepath.Join("temp/node")
	NoErr(os.MkdirAll(dir, 0755))
	return dir
}

func NoErr(err error, extra ...interface{}) {
	if err != nil {
		for _, i := range extra {
			fmt.Println(JsonIndent(i))
			fmt.Println()
		}
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Errorf("runtime error %s:%d : %s", file, line, err))
	}
}

func Recover() {
	e := recover()
	if e != nil {
		switch err := e.(type) {
		case error:
			errString := err.Error()
			if strings.Contains(errString, "runtime error") {
				panic(err)
			}
		}
		fmt.Printf("error: %+v\n", e)
		os.Exit(1)
	}
}

func JsonIndentBytesToFile(data []byte, path string) {
	m := make(map[string]interface{})
	NoErr(json.Unmarshal(data, &m))
	JsonIndentToFile(m, path)
}

func JsonIndentToFile(i interface{}, path string) {
	NoErr(os.WriteFile(path, []byte(JsonIndent(i)), 0755))
}

func JsonIndent(i interface{}) string {
	indent, err := json.MarshalIndent(i, "", "  ")
	NoErr(err)
	return string(indent)
}

func Bolt(dbName string) *bolt.Bolt {
	db, err := bolt.New(filepath.Join("temp", dbName))
	NoErr(err)
	return db
}

const InspectDatabase = "tsInspect.bolt"

func BoltClose(db *bolt.Bolt) {
	NoErr(db.Close())
}
