package typescript

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlctrez/goapp-pf/scripts"
	"github.com/mlctrez/goapp-pf/typescript/kind"
)

type SourceFile struct {
	Imports        []*ImportDeclaration
	Interfaces     []*InterfaceDeclaration
	Classes        []*ClassDeclaration
	TypeAliases    []*TypeAliasDeclaration
	Statements     []interface{} `json:"statements"`
	EndOfFileToken *Node         `json:"endOfFileToken"`
	FileName       string        `json:"fileName"`
	Text           string        `json:"text"`
}

func (s *SourceFile) ParseStatements() {
	noErr := scripts.NoErr

	for _, statement := range s.Statements {

		statementJson, err := json.Marshal(statement)
		noErr(err)

		imp := &ImportDeclaration{}
		noErr(json.Unmarshal(statementJson, imp))

		switch imp.Kind {
		case kind.ImportDeclaration:
			s.Imports = append(s.Imports, imp)
			continue

		case kind.InterfaceDeclaration:
			id := &InterfaceDeclaration{}
			noErr(json.Unmarshal(statementJson, id))
			s.Interfaces = append(s.Interfaces, id)
			continue

		case kind.ClassDeclaration:
			cd := &ClassDeclaration{}
			noErr(json.Unmarshal(statementJson, cd))
			s.Classes = append(s.Classes, cd)
			continue

		case kind.TypeAliasDeclaration:
			ta := &TypeAliasDeclaration{}
			noErr(json.Unmarshal(statementJson, ta))
			s.TypeAliases = append(s.TypeAliases, ta)
			continue

		case kind.VariableStatement,
			kind.ExpressionStatement,
			kind.EnumDeclaration,
			kind.FunctionDeclaration:
			// TODO: need implementation for each

		default:
			noErr(fmt.Errorf("unhandled kind %d", imp.Kind))
		}
	}
}

func (s *SourceFile) Dump() {

	noErr := scripts.NoErr

	fmt.Println("\n****", s.FileName)

	for j, id := range s.Interfaces {
		//break
		fmt.Println("ID", j, id.InterfaceName())
		for i, member := range id.Members {
			fmt.Println("ID", j, i, member.MemberName(), ":", member.Type.KindString())
		}
	}

	for j, id := range s.TypeAliases {
		fmt.Println("TA", j, id.Name.LowerText()+" = "+id.AliasValue())
	}

	for j, id := range s.Imports {
		break
		marshal, err := json.Marshal(id)
		noErr(err)
		fmt.Println("IMP", j, string(marshal))

	}

	for j, id := range s.Classes {
		break
		marshal, err := json.MarshalIndent(id, "", "  ")
		noErr(err)
		fmt.Println("CD", j, string(marshal))
	}

}

/*
TypeAliasDeclaration  258
*/

func (d *TypeAliasDeclaration) AliasValue() string {
	switch d.Type.Kind {
	case kind.UnionType:
		var parts []string
		for _, typ := range d.Type.Types {
			parts = append(parts, typ.Literal.LowerText())
		}
		return strings.Join(parts, " | ")
	case kind.FunctionType:

		var paramList []string
		for _, parameter := range d.Type.Parameters {
			name := parameter.Name.LowerText()
			var typeArgs []string
			for _, argument := range parameter.Type.TypeArguments {
				typeArgs = append(typeArgs, argument.LowerText())
			}
			var typeName string
			if parameter.Type.TypeName != nil {
				typeName = parameter.Type.TypeName.TypeName()
			} else {
				switch parameter.Type.Kind {
				case kind.StringKeyword:
					typeName = "string"
				case kind.UnionType:

					var unionParts []string

					for _, n := range parameter.Type.Types {
						unionParts = append(unionParts, n.Kind.ToType())
					}
					return strings.Join(unionParts, " | ")
				default:
					scripts.NoErr(fmt.Errorf("unhandled type %d", parameter.Type.Kind))
				}
			}

			param := fmt.Sprintf("%s: %s<%s>", name, typeName, strings.Join(typeArgs, ", "))
			paramList = append(paramList, param)
		}

		funcSig := fmt.Sprintf("func(%s)", strings.Join(paramList, ", "))

		if d.Type.Type.Kind != kind.VoidKeyword {
			scripts.NoErr(fmt.Errorf("unhandled type %d", d.Type.Type.Kind))
			return fmt.Sprintf("func() %s", "TODO")
		}

		return funcSig

	default:
		fmt.Println(scripts.JsonIndent(d.Type))
		scripts.NoErr(fmt.Errorf("unhandled type %d", d.Type.Kind))
		return ""
	}
}

type TypeAliasDeclaration struct {
	Node
	Modifiers []*Node        `json:"modifiers,omitempty"`
	Name      *Node          `json:"name"`
	Type      *TypeAliasType `json:"type"`
}

type TypeAliasType struct {
	Node
	Type       *Node                   `json:"types"`
	Types      []*TypeAliasTypeLiteral `json:"types"`
	Parameters []*TypeAliasParameter   `json:"parameters"`
}

type TypeAliasTypeLiteral struct {
	Node
	Literal *Node `json:"literal"`
}

type TypeAliasParameter struct {
	Node
	Name *Node                   `json:"name"`
	Type *TypeAliasParameterType `json:"type"`
}

type TypeAliasParameterType struct {
	Node
	TypeArguments []*TypeAliasParameterArguments `json:"typeArguments"`
	TypeName      *TypeName                      `json:"typeName"`
	Types         []*Node                        `json:"types"`
}

type TypeAliasParameterArguments struct {
	Node
	TypeName *TypeName `json:"typeName"`
}

/*
ImportDeclaration  265
*/

type ImportDeclaration struct {
	Node
	ImportClause    *ImportClause    `json:"importClause"`
	ModuleSpecifier *ModuleSpecifier `json:"moduleSpecifier"`
}

type ImportClause struct {
	Node
	NamedBindings *NamedBindings `json:"namedBindings"`
	IsTypeOnly    bool           `json:"isTypeOnly"`
	Name          *Node          `json:"name,omitempty"`
}

type ModuleSpecifier struct {
	Node
	HasExtendedUnicodeEscape bool `json:"hasExtendedUnicodeEscape"`
}

type NamedBindings struct {
	Node
	Name *Node `json:"name"`
}

/*
InterfaceDeclaration  257
*/

type InterfaceDeclaration struct {
	Name Node `json:"name"`
	Node
	Members []*InterfaceMember `json:"members"`
}

func (i *InterfaceDeclaration) InterfaceName() string {
	return i.Name.EscapedText
}

type InterfaceMember struct {
	Node
	Name          *Node                `json:"name"`
	QuestionToken *Node                `json:"questionToken,omitempty"`
	Type          *InterfaceMemberType `json:"type"`
	JsDoc         []*JsDoc             `json:"jsDoc"`
	Parameters    []*Parameter         `json:"parameters"`
}

type JsDoc struct {
	Node
	Comment string      `json:"comment"`
	Parent  interface{} `json:"parent,omitempty"`
}

type Parameter struct {
	Node
	Name *Node `json:"name"`
	Type *Node `json:"type"`
}

func (p *Parameter) Definition() string {
	key := p.Name.UpperText()
	if p.Type.Kind == kind.StringKeyword {
		return fmt.Sprintf("%s: string", key)
	}
	scripts.NoErr(fmt.Errorf("unhandled type %d", p.Type.Kind))
	return ""
}

func hyphenToCamel(text string) string {
	if !strings.Contains(text, "-") {
		return title(text)
	}
	parts := strings.Split(text, "-")
	var newParts []string
	for _, part := range parts {
		newParts = append(newParts, title(part))
	}
	return strings.Join(newParts, "")
}

func title(text string) string {
	if len(text) < 2 {
		return text
	}
	return strings.ToUpper(text[0:1]) + text[1:]
}

func (m *InterfaceMember) MemberName() string {

	if m.Kind == kind.IndexSignature {
		var parameterStrings []string
		for _, parameter := range m.Parameters {
			parameterStrings = append(parameterStrings, parameter.Definition())
		}

		return fmt.Sprintf("[%s]", strings.Join(parameterStrings, ", "))
	}

	if m.Name == nil {
		fmt.Printf("%+v\n", m)
		scripts.NoErr(fmt.Errorf("no name"))
	}

	text := m.Name.UpperText()
	if m.QuestionToken != nil {
		return fmt.Sprintf("%s?", text)
	}
	return text
}

type InterfaceMemberType struct {
	Node
	TypeName *TypeName              `json:"typeName,omitempty"`
	Types    []*InterfaceMemberType `json:"types,omitempty"`
	Type     *InterfaceMemberType   `json:"type,omitempty"`
	Literal  *Node                  `json:"literal,omitempty"`
	Members  []*InterfaceMember     `json:"members"`
}

func (mt *InterfaceMemberType) TypeLiteral() string {
	var result []string
	for _, member := range mt.Members {
		result = append(result, member.MemberName()+": "+member.Type.KindString())
	}
	return fmt.Sprintf("{ %s }", strings.Join(result, ", "))
}

func (mt *InterfaceMemberType) KindString() string {
	switch mt.Kind {
	case kind.VoidKeyword:
		return "void"
	case kind.BooleanKeyword:
		return "bool"
	case kind.NumberKeyword:
		return "float64"
	case kind.StringKeyword:
		return "string"
	case kind.ArrayType:
		return "ArrayType"
	case kind.LiteralType:
		return mt.Literal.Text
	case kind.TypeReference:
		return mt.TypeName.TypeName()
	case kind.AnyKeyword:
		return "any"
	case kind.TypeLiteral:
		return mt.TypeLiteral()
	case kind.FunctionType:
		if mt.Type.Kind != kind.VoidKeyword {
			return fmt.Sprintf("func() %s", mt.Type.KindString())
		}
		return "func()"
	case kind.UnionType:
		var unionKinds []string
		for _, memberType := range mt.Types {
			unionKinds = append(unionKinds, memberType.KindString())
		}
		return strings.Join(unionKinds, " | ")
	case kind.ParenthesizedType:
		return fmt.Sprintf("(%s)", mt.Type.KindString())
	default:
		//scripts.NoErr(fmt.Errorf("unhandled kind %d", mt.Kind))

		return fmt.Sprintf("Kind:%d", mt.Kind)
	}
}

type TypeName struct {
	Node
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func (n *TypeName) TypeName() string {
	if n.EscapedText != "" {
		return n.EscapedText
	}
	return n.Left.EscapedText + "." + n.Right.EscapedText
}

type Node struct {
	Pos                int    `json:"pos"`
	End                int    `json:"end"`
	Kind               Kind   `json:"kind"`
	ModifierFlagsCache int    `json:"modifierFlagsCache,omitempty"`
	TransformFlags     int    `json:"transformFlags,omitempty"`
	EscapedText        string `json:"escapedText,omitempty"`
	Flags              int    `json:"flags,omitempty"`
	Text               string `json:"text,omitempty"`
}

type Kind int

func (k Kind) ToType() string {
	switch k {
	case kind.VoidKeyword:
		return "void"
	case kind.BooleanKeyword:
		return "bool"
	case kind.NumberKeyword:
		return "float64"
	case kind.StringKeyword:
		return "string"
	default:
		scripts.NoErr(fmt.Errorf("Kind.ToType() missing case for %d", k))
		return ""
	}
}

func (n *Node) LowerText() string {
	text := n.EscapedText
	if text == "" {
		text = n.Text
	}
	return text
}

func (n *Node) UpperText() string {
	return hyphenToCamel(n.LowerText())
}

/*
ClassDeclaration  256
*/

type ClassDeclaration struct {
	Node
	Modifiers []*Node        `json:"modifiers,omitempty"`
	Name      *Node          `json:"name"`
	Members   []*ClassMember `json:"members"`
}

type ClassMember struct {
	Node
	Initializer *Initializer `json:"initializer"`
	Modifiers   []*Node      `json:"modifiers,omitempty"`
	Name        *Node        `json:"name"`
}
type Initializer struct {
	Node
	HasExtendedUnicodeEscape bool `json:"hasExtendedUnicodeEscape"`
}
