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

func NoErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Errorf("%s:%d : %s", file, line, err))
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
