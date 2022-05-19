package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlctrez/bolt"
	"github.com/mlctrez/goapp-pf/scripts"
)

func main() {

	defer scripts.Recover()
	noErr := scripts.NoErr

	db, err := bolt.New(filepath.Join("temp", "types.bolt"))
	noErr(err)
	defer func() {
		errClose := db.Close()
		if errClose != nil {
			log.Println("errClose : ", err)
		}
	}()

	componentsBucket := bolt.Key("components")

	val := &bolt.Value{K: "node_modules/@patternfly/react-core/src/components/AboutModal/AboutModal.tsx"}
	noErr(db.Get(componentsBucket, val))

	counter := 1
	jsonDir := filepath.Join("temp", "json")
	noErr(os.MkdirAll(jsonDir, 0755))

	scanner := bufio.NewScanner(bytes.NewBuffer(val.V))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "\t")

		m := make(map[string]interface{})
		noErr(json.Unmarshal([]byte(parts[1]), &m))

		var formatted []byte
		formatted, err = json.MarshalIndent(m, "", "  ")
		noErr(err)

		fileName := fmt.Sprintf("%02d_%s.json", counter, parts[0])

		fullPath := filepath.Join(jsonDir, fileName)
		noErr(os.WriteFile(fullPath, formatted, 0666))

		counter++
	}

}
