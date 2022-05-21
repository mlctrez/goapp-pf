package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"

	"github.com/mlctrez/bolt"
	"github.com/mlctrez/goapp-pf/scripts"
	"github.com/mlctrez/goapp-pf/typescript/recorder"
	"go.etcd.io/bbolt"
)

func main() {

	defer scripts.Recover()
	noErr := scripts.NoErr

	db := scripts.Bolt(scripts.InspectDatabase)
	defer func() { scripts.BoltClose(db) }()

	buckets := bolt.Keys{"components", "layouts", "typescript"}
	noErr(db.CreateBuckets(buckets))

	buckets = bolt.Keys{"components"}

	rootNode := &Node{Name: "sourceFile"}

	ntr := &recorder.NodeTypesRecorder{}

	noErr(db.View(func(tx *bbolt.Tx) error {
		for _, bk := range buckets {
			bucket := tx.Bucket(bk.B())
			cursor := bucket.Cursor()
			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				m := make(map[string]interface{})
				noErr(json.Unmarshal(v, &m))
				rootNode.Attr(m, ntr.AttrRecorder)
			}
		}
		return nil
	}))

	noErr(os.MkdirAll("typescript", 0755))

	ntr.PluralFix()

	buf := &bytes.Buffer{}
	buf.WriteString("package typescript\n")
	buf.WriteString("import (\n")
	buf.WriteString(fmt.Sprintf("\t%q", "github.com/mlctrez/goapp-pf/typescript/kind"))
	buf.WriteString(")\n")

	ntr.WriteCode(buf)
	formatted, err := format.Source(buf.Bytes())
	noErr(err)
	noErr(os.WriteFile("typescript/types.go", formatted, 0755))

}

type Node struct {
	Name       string
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Children   map[string]*Node       `json:"children,omitempty"`
}

func (n *Node) initAttributes() {
	if n.Attributes == nil {
		n.Attributes = make(map[string]interface{})
	}
}

func (n *Node) initChildren() {
	if n.Children == nil {
		n.Children = make(map[string]*Node)
	}
}

func (n *Node) Attr(m map[string]interface{}, r recorder.AttrRecorder) {

	for k, v := range m {
		r(n.Name, k, v)
		switch vt := v.(type) {
		case []interface{}:
			for _, av := range vt {
				key := "[]" + k
				n.initChildren()
				child := n.Children[key]
				if child == nil {
					child = &Node{Name: k}
					n.Children[key] = child
				}
				child.Attr(av.(map[string]interface{}), r)
			}
		case map[string]interface{}:
			n.initChildren()
			child := n.Children[k]
			if child == nil {
				child = &Node{Name: k}
				n.Children[k] = child
			}
			child.Attr(vt, r)
		default:
			n.initAttributes()
			n.Attributes[k] = vt
		}
	}
}
