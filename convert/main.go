package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlctrez/bolt"
	"github.com/mlctrez/goapp-pf/internal/util"
	"github.com/mlctrez/goapp-pf/scripts"
	"github.com/mlctrez/goapp-pf/typescript"
	"github.com/mlctrez/goapp-pf/typescript/kind"
	"go.etcd.io/bbolt"
)

var currentFile string

var saveOffCurrent bool
var writeOutputFile bool

func main() {
	defer scripts.Recover()

	defer func() {
		if saveOffCurrent && currentFile != "" {
			fileData, err := os.ReadFile(currentFile)
			if err != nil {
				fmt.Println("error reading current file", err)
				return
			}
			err = os.WriteFile("convert/current.tsx", fileData, 0644)
			if err != nil {
				fmt.Println("error writing current file", err)
				return
			}
			if currentSourceStream != nil {
				err = os.WriteFile("convert/currentSourceStream.txt",
					currentSourceStream.Bytes(), 0644)
				if err != nil {
					fmt.Println("error writing currentSourceStream file", err)
					return
				}
			}
		}
	}()

	noErr := scripts.NoErr
	_ = noErr

	db := scripts.Bolt(scripts.InspectDatabase)
	defer func() { scripts.BoltClose(db) }()

	progressBucket := bolt.Key("progress")
	noErr(db.CreateBuckets(bolt.Keys{progressBucket}))

	//scripts.JsonIndentBytesToFile(v, "convert/"+filepath.Base(pfx)+".json")

	ignoreCurrent := true
	onlyOnce := false
	saveOffCurrent = false
	writeOutputFile = true

	buckets := &bolt.Keys{"components", "layouts"}
	for _, buck := range *buckets {
		currentKey := string(buck.B())
		current := &bolt.Value{K: bolt.Key(currentKey)}

		if ignoreCurrent {
			noErr(db.Delete(progressBucket, current.K))
		}

		err := db.Get(progressBucket, current)
		if err != nil && err != bolt.ErrValueNotFound {
			scripts.NoErr(err)
		}
		if string(current.V) == "completedWithoutError" {
			continue
		}

		fmt.Printf("bucket %q starting at %q\n", string(buck.B()), string(current.V))

		noErr(db.View(func(tx *bbolt.Tx) error {
			bucket := tx.Bucket(buck.B())
			cursor := bucket.Cursor()

			for k, v := cursor.Seek(current.V); k != nil; k, v = cursor.Next() {

				if !onlyOnce {
					current.V = k
					noErr(db.Put(progressBucket, current))
				}

				sf := &typescript.SourceFile{}
				noErr(json.Unmarshal(v, sf))

				if saveOffCurrent {
					scripts.JsonIndentBytesToFile(v, "convert/current.json")
					currentFile = filepath.Join("temp", "node", sf.FileName)
					fmt.Printf("file:  " + currentFile + ":1\n")
				}
				maybeWriteOutput(sf)
			}
			return nil
		}))

		if !onlyOnce {
			current.V = []byte("completedWithoutError")
			noErr(db.Put(progressBucket, current))
		}

	}

}

var currentSourceStream *bytes.Buffer

func maybeWriteOutput(file *typescript.SourceFile) {

	fmt.Println(file.FileName)

	var match *DirectoryMatch
	for _, matcher := range matchers {
		if matcher.Matches(file.FileName) {
			match = matcher
			break
		}
	}
	if match == nil {
		scripts.NoErr(fmt.Errorf("no matcher for FileName %s", file.FileName))
	}
	pkg, finalPath := match.GoPackageAndPath(file.FileName)

	buf := &bytes.Buffer{}
	currentSourceStream = buf
	pl := func(format string, args ...any) {
		buf.WriteString(fmt.Sprintf(format, args...) + "\n")
	}

	pl("package %s", pkg)

	var importStatements []*typescript.Statement

	for _, statement := range file.Statements {
		switch statement.Kind {
		case kind.ImportDeclaration:

			if statement.ImportClause != nil && statement.ImportClause.Name != nil {
				if statement.ImportClause.Name.TextString() == "styles" {
					fmt.Println(statement.ModuleSpecifier.Text)
				}
			}

			importStatements = append(importStatements, statement)
		case kind.InterfaceDeclaration:

			typeName := statement.Name.TextString()
			if strings.HasPrefix(strings.ToLower(typeName), pkg) {
				typeName = typeName[len(pkg):]
			}

			if len(statement.Members) == 1 && statement.Members[0].Kind == kind.IndexSignature {
				m := statement.Members[0]
				// convert this to a map type declaration
				pl("\ntype %s map[%s]%s", typeName, m.Parameters[0].Type.GoType(), m.Type.GoType())
				continue
			}

			pl("\ntype %s struct {", typeName)
			for _, member := range statement.Members {
				pl(member.Docstring())
				goType := member.Type.GoType()
				if strings.HasPrefix(goType, "any") {
					goType = "any // " + strings.TrimPrefix(goType, "any")
				}

				pl("\t%s %s", util.Title(member.Name.TextString()), goType)
			}

			pl("}")
		case kind.ClassDeclaration:

		}

	}
	scripts.JsonIndentToFile(importStatements, "convert/imports.json")

	scripts.NoErr(fmt.Errorf("halt"))

	pl("")
	pl("// contents of %s", strings.TrimPrefix(file.FileName, "node_modules/@patternfly/"))
	scanner := bufio.NewScanner(bytes.NewBufferString(file.Text))
	for scanner.Scan() {
		pl(fmt.Sprintf("// %s", scanner.Text()))
	}
	scripts.NoErr(scanner.Err())

	if writeOutputFile {
		scripts.NoErr(os.MkdirAll(filepath.Dir(finalPath), 0755))
		scripts.NoErr(os.WriteFile(finalPath, buf.Bytes(), 0644))
		//jsonFile := strings.TrimSuffix(finalPath, ".go") + ".json"
		//file.Text = "omitted for space"
		//scripts.JsonIndentToFile(file, jsonFile)
	}
}

type DirectoryMatch struct {
	PathPart string
	Dir      string
}

func (dm *DirectoryMatch) Matches(path string) bool {
	return strings.Contains(path, dm.PathPart)
}

func (dm *DirectoryMatch) GoPackageAndPath(path string) (string, string) {
	parts := strings.Split(path, dm.PathPart)
	if len(parts) != 2 {
		scripts.NoErr(fmt.Errorf("unable to split path parts for %q using %q", path, dm.PathPart))
	}

	split := strings.Split(parts[1], "/")
	pkg := strings.ToLower(split[0])
	if pkg == "select" || pkg == "switch" {
		pkg = "pf" + pkg
	}
	goFileName := strings.TrimSuffix(strings.ToLower(split[1]), "tsx") + "go"

	if strings.HasPrefix(goFileName, pkg) {
		goFileName = strings.TrimPrefix(goFileName, pkg)
		if goFileName == ".go" {
			goFileName = "component.go"
		}
	}

	finalPath := filepath.Join(dm.Dir, pkg, goFileName)
	return pkg, finalPath
}

var matchers = []*DirectoryMatch{
	{PathPart: "/react-core/src/components/", Dir: "core/components"},
	{PathPart: "/react-core/src/layouts/", Dir: "core/layouts"},
}
