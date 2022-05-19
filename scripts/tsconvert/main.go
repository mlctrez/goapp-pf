package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlctrez/bolt"
	"github.com/mlctrez/goapp-pf/attic/typescript"
	"github.com/mlctrez/goapp-pf/scripts"
	"go.etcd.io/bbolt"
)

func main() {
	defer scripts.Recover()
	noErr := scripts.NoErr

	db := scripts.Bolt(scripts.InspectDatabase)
	defer func() { scripts.BoltClose(db) }()

	componentsBucket := bolt.Key("components")

	single := "node_modules/@patternfly/react-core/src/components/NumberInput/NumberInput.tsx"
	if os.Getenv("foo") == "asd" {
		single = ""
	}

	if single != "" {

		v := &bolt.Value{K: bolt.Key(single)}
		noErr(db.Get(componentsBucket, v))

		m := make(map[string]interface{})
		noErr(json.Unmarshal(v.V, &m))
		formatted, err := json.MarshalIndent(m, "", "  ")
		noErr(err)

		jsonFile := strings.ReplaceAll(filepath.Base(single), ".tsx", ".json")
		jsonFilePath := filepath.Join("temp", "json", jsonFile)
		noErr(os.WriteFile(jsonFilePath, formatted, 0666))

		sf := &typescript.SourceFile{}
		noErr(json.Unmarshal(v.V, &sf))

		tsxFile := filepath.Join("temp", "json", filepath.Base(single))
		noErr(os.WriteFile(tsxFile, []byte(sf.Text), 0666))

		sf.ParseStatements()
		sf.Dump()
		return
	}

	noErr(db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(componentsBucket.B())
		if bucket == nil {
			return fmt.Errorf("invalid bucket name %s", componentsBucket)
		}
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			m := &typescript.SourceFile{}
			noErr(json.Unmarshal(v, &m))

			m.ParseStatements()
			m.Dump()
		}
		return nil
	}))

}
