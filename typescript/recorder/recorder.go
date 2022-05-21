package recorder

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type AttrRecorder func(nodeName, property string, v interface{})

type NodeTypesRecorder struct {
	NodeTypesMap map[string]map[string]string
}

func (n *NodeTypesRecorder) AttrRecorder(nodeName, property string, v interface{}) {
	if n.NodeTypesMap == nil {
		n.NodeTypesMap = make(map[string]map[string]string)
	}
	if n.NodeTypesMap[nodeName] == nil {
		n.NodeTypesMap[nodeName] = make(map[string]string)
	}

	var tv string
	switch v.(type) {
	case nil:
		tv = "any"
	case []interface{}:
		n.AttrRecorder(property, "pos", 0)
		tv = fmt.Sprintf("[]*%s", title(property))
	case map[string]interface{}:
		n.AttrRecorder(property, "pos", 0)
		tv = fmt.Sprintf("*%s", title(property))
	case float64:
		tv = "int"
	case string:
		tv = "string"
		// the one and only quirk where this attribute can be a string or an array
		if nodeName == "jsDoc" && property == "comment" {
			tv = "any"
		}
	default:
		tv = reflect.TypeOf(v).String()
	}

	n.NodeTypesMap[nodeName][property] = tv

}

func (n *NodeTypesRecorder) PluralFix() {
	var plurals []string
	for nodeName := range n.NodeTypesMap {
		// the one and only conflict that we cannot de-duplicate
		if strings.HasSuffix(nodeName, "s") {
			plurals = append(plurals, nodeName)
		}
	}

	wouldReduce := make(map[string]int)
	for _, nodeMap := range n.NodeTypesMap {
		for nodeProp, nodeVal := range nodeMap {
			for _, plural := range plurals {
				if plural == nodeProp && fmt.Sprintf("[]*%s", title(plural)) == nodeVal {
					wouldReduce[plural]++
				}
			}
		}
	}

	for pluralName := range wouldReduce {

		var singularName string
		if strings.HasSuffix(pluralName, "ies") {
			singularName = strings.TrimSuffix(pluralName, "ies") + "y"
		} else {
			singularName = strings.TrimSuffix(pluralName, "s")
		}

		// rename the top level element
		topLevel := n.NodeTypesMap[pluralName]
		delete(n.NodeTypesMap, pluralName)
		if prev, ok := n.NodeTypesMap[singularName]; ok {
			// merge types onto type
			for k, v := range topLevel {
				prev[k] = v
			}
		} else {
			n.NodeTypesMap[singularName] = topLevel
		}

		// find all arrays that need the type renamed
		for _, nodeMap := range n.NodeTypesMap {
			if nodeMap[pluralName] == fmt.Sprintf("[]*%s", title(pluralName)) {
				nodeMap[pluralName] = fmt.Sprintf("[]*%s", title(singularName))
			}
		}

	}

}

func (n *NodeTypesRecorder) WriteCode(b *bytes.Buffer) {

	pl := func(format string, args ...any) {
		b.WriteString(fmt.Sprintf(format, args...))
		b.WriteString("\n")
	}
	nodeAttributes := []string{"pos", "end", "kind", "flags", "transformFlags", "modifierFlagsCache"}
	n.NodeTypesMap["node"] = map[string]string{}
	for _, attribute := range nodeAttributes {
		if attribute == "kind" {
			n.NodeTypesMap["node"][attribute] = "kind.Kind"
		} else {
			n.NodeTypesMap["node"][attribute] = "int"
		}
	}
	for _, name := range sortedNodes(n) {
		pl("\ntype %s struct {", title(name))
		if name != "node" {
			pl("\tNode")
			for _, delAttr := range nodeAttributes {
				delete(n.NodeTypesMap[name], delAttr)
			}
			delete(n.NodeTypesMap[name], "parent")
		}
		nodeTypes := n.NodeTypesMap[name]
		for _, k := range sortedProperties(nodeTypes) {
			v := nodeTypes[k]
			var omitSuffix = ""
			if strings.Contains(v, "*") || v == "any" {
				omitSuffix = ",omitempty"
			}
			pl("\t %s %s `json:\"%s%s\"`", title(k), v, k, omitSuffix)
		}
		pl("}")
	}

}

func sortedNodes(n *NodeTypesRecorder) []string {
	var sorted []string
	for k := range n.NodeTypesMap {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}

func sortedProperties(m map[string]string) []string {
	var sorted []string
	for k := range m {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}

func title(text string) string {
	if len(text) < 2 {
		return text
	}
	return strings.ToUpper(text[0:1]) + text[1:]
}
