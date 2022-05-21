package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mlctrez/bolt"
	"github.com/mlctrez/goapp-pf/internal/util"
	"github.com/mlctrez/goapp-pf/scripts"
	"github.com/mlctrez/goapp-pf/typescript"
	"github.com/mlctrez/goapp-pf/typescript/kind"
	"go.etcd.io/bbolt"
)

func main() {
	defer scripts.Recover()
	noErr := scripts.NoErr
	_ = noErr

	db := scripts.Bolt(scripts.InspectDatabase)
	defer func() { scripts.BoltClose(db) }()

	pfx := "node_modules/@patternfly/react-core/src/components"

	outputGo := false
	// pfx = "node_modules/@patternfly/react-core/src/components/Dropdown/InternalDropdownItem.tsx"

	var sources []*typescript.SourceFile

	components := bolt.Key("components")
	noErr(db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(components.B())
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			if strings.HasPrefix(string(k), pfx) {
				sf := &typescript.SourceFile{}
				noErr(json.Unmarshal(v, sf))
				sources = append(sources, sf)
				//scripts.JsonIndentBytesToFile(v, "convert/"+filepath.Base(pfx)+".json")
			}
		}
		return nil
	}))

	buf := &bytes.Buffer{}
	pl := func(format string, args ...any) {
		buf.WriteString(fmt.Sprintf(format, args...) + "\n")
	}

	for _, file := range sources {

		fmt.Printf("file:  " + filepath.Join("temp", "node", file.FileName) + ":1\n")

		for _, statement := range file.Statements {
			switch statement.Kind {
			case kind.ImportDeclaration:
			case kind.InterfaceDeclaration:

				if len(statement.Members) == 1 && statement.Members[0].Kind == kind.IndexSignature {
					m := statement.Members[0]
					// convert this to a map type declaration
					pl("\ntype %s map[%s]%s", statement.Name.TextString(), m.Parameters[0].Type.GoType(), m.Type.GoType())
					continue
				}

				pl("\ntype %s struct {", statement.Name.TextString())
				for _, member := range statement.Members {
					pl(member.Docstring())
					pl("\t%s %s", util.Title(member.Name.TextString()), member.Type.GoType())
				}

				pl("}")
			case kind.ClassDeclaration:

			}
		}
		if outputGo {
			fmt.Printf(buf.String())
		}
		buf.Reset()
	}

}
