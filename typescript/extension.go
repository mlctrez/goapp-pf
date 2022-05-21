package typescript

import (
	"fmt"
	"strings"

	"github.com/mlctrez/goapp-pf/internal/util"
	"github.com/mlctrez/goapp-pf/scripts"
	"github.com/mlctrez/goapp-pf/typescript/kind"
)

func (n *Name) TextString() string {
	if n.EscapedText != "" {
		return n.EscapedText
	}
	return n.Text
}

func (t *Type) GoType() string {
	switch t.Kind {
	case kind.StringKeyword:
		return "string"
	case kind.BooleanKeyword:
		return "bool"
	case kind.AnyKeyword:
		return "any"
	case kind.NumberKeyword:
		return "int"
	case kind.LiteralType:
		return fmt.Sprintf("%q", t.Literal.Text)
	case kind.TypeOperator:
		return "any /* keyof " + t.Type.GoType() + " */"
	case kind.TypeReference:
		return t.TypeName.TextString()
	case kind.TypeQuery:
		return fmt.Sprintf("typeof %s", t.ExprName.EscapedText)
	case kind.TypeLiteral:
		if t.Literal != nil && t.Literal.Text != "" {
			return fmt.Sprintf("%q", t.Literal.Text)
		}
		if t.Members != nil {
			var memberValues []string
			for _, member := range t.Members {
				memType := member.Type
				txt := member.Name.TextString()
				switch memType.Kind {
				case kind.AnyKeyword:
					mv := fmt.Sprintf("%q:%s", txt, "any")
					memberValues = append(memberValues, mv)
				case kind.BooleanKeyword:
					mv := fmt.Sprintf("%q:%s", txt, "bool")
					memberValues = append(memberValues, mv)
				case kind.StringKeyword:
					mv := fmt.Sprintf("%q:%s", txt, "string")
					memberValues = append(memberValues, mv)
				case kind.FunctionType:
					mv := fmt.Sprintf("%q:%s", txt, functionSignature(memType.Type, memType.Parameters))
					memberValues = append(memberValues, mv)
				case kind.UnionType:
					var unionParts []string
					for _, m := range memType.Types {
						switch m.Kind {
						case kind.LiteralType:
							unionParts = append(unionParts, m.Literal.Text)
						case kind.NumberKeyword:
							unionParts = append(unionParts, "int")
						case kind.StringKeyword:
							unionParts = append(unionParts, "string")
						default:
							fmt.Println(scripts.JsonIndent(member))
							scripts.NoErr(fmt.Errorf("memType.Types.Kind %v", m.Kind.StringValue()))
						}
					}
					memberValues = append(memberValues, fmt.Sprintf("%s:{ %s }", txt, strings.Join(unionParts, ", ")))
				case kind.TypeReference:
					mv := fmt.Sprintf("%q:%s", txt, member.Type.TypeName.TextString())
					memberValues = append(memberValues, mv)
				default:
					fmt.Println(scripts.JsonIndent(member))
					scripts.NoErr(fmt.Errorf("member.Type.Kind %v", memType.Kind.StringValue()))
				}
			}
			return fmt.Sprintf("map[string]string // { %s }", strings.Join(memberValues, ", "))
		}
		fmt.Println(scripts.JsonIndent(t))
		scripts.NoErr(fmt.Errorf("syntax.TypeLiteral unhandled"))
		return ""
	case kind.ArrayType:
		if t.ElementType != nil {
			switch t.ElementType.Kind {
			case kind.TypeReference:
				return "[]" + t.ElementType.TypeName.TextString()
			case kind.StringKeyword:
				return "[]string"
			case kind.NumberKeyword:
				return "[]int"
			case kind.AnyKeyword:
				return "[]any"
			case kind.FunctionType, kind.ParenthesizedType:
				elType := t.ElementType.Type
				return "[]" + functionSignature(elType.Type, elType.Parameters)
			default:
				fmt.Println(scripts.JsonIndent(t))
				scripts.NoErr(fmt.Errorf("t.ElementType.Kind %d", t.ElementType.Kind))
				return ""
			}
		}
		fmt.Println(scripts.JsonIndent(t))
		scripts.NoErr(fmt.Errorf("t.ElementType ==nil"))
		return ""

	case kind.ParenthesizedType:
		return fmt.Sprintf("(%s)", t.Type.GoType())
	case kind.FunctionType:
		return functionSignature(t.Type, t.Parameters)
	case kind.UnionType:
		var unionParts []string
		for _, p := range t.Types {
			switch p.Kind {
			case kind.LiteralType:
				unionParts = append(unionParts, fmt.Sprintf("%q", p.Literal.Text))
			case kind.UndefinedKeyword:
				unionParts = append(unionParts, "undefined")
			case kind.StringKeyword:
				unionParts = append(unionParts, "string")
			case kind.BooleanKeyword:
				unionParts = append(unionParts, "bool")
			case kind.NumberKeyword:
				unionParts = append(unionParts, "int")
			case kind.TypeQuery:
				typeOf := fmt.Sprintf("typeof %s", p.ExprName.EscapedText)
				unionParts = append(unionParts, typeOf)
			case kind.ArrayType:
				if p.ElementType == nil {
					fmt.Println(scripts.JsonIndent(p))
					scripts.NoErr(fmt.Errorf("syntax.ArrayType - nill p.ElementType"))
				}
				switch p.ElementType.Kind {
				case kind.TypeReference:
					part := "[]" + p.ElementType.TypeName.TextString()
					unionParts = append(unionParts, part)
				case kind.StringKeyword:
					unionParts = append(unionParts, "[]string")
				case kind.NumberKeyword:
					unionParts = append(unionParts, "[]int")
				case kind.AnyKeyword:
					unionParts = append(unionParts, "[]any")
				default:
					fmt.Println(scripts.JsonIndent(t))
					scripts.NoErr(fmt.Errorf("unhandled union array type %v", p.ElementType.Kind.StringValue()))
				}
			case kind.TypeReference:
				unionParts = append(unionParts, p.TypeName.TextString())
			case kind.ParenthesizedType:
				unionParts = append(unionParts, p.Type.GoType())
			case kind.AnyKeyword:
				unionParts = append(unionParts, "any")

			default:
				fmt.Println(scripts.JsonIndent(t))
				scripts.NoErr(fmt.Errorf("unhandled union Types kind %s", p.Kind.StringValue()))
			}
		}
		// TODO: peel of each of these types and write out a separate constant declaration?
		return "any /* " + strings.Join(unionParts, " | ") + " */"
	default:
		fmt.Printf(scripts.JsonIndent(t))
		scripts.NoErr(fmt.Errorf(fmt.Sprintf("Kind:%s", t.Kind.StringValue())))
		return ""
	}
}

func functionSignature(typ *Type, parameters []*Parameter) string {
	var params []string
	for _, p := range parameters {
		params = append(params, fmt.Sprintf("%s %s", p.Name.TextString(), p.Type.GoType()))
	}
	funcSig := fmt.Sprintf("func(%s)", strings.Join(params, ", "))
	if typ.Kind != kind.VoidKeyword {
		return fmt.Sprintf("%s %s", funcSig, typ.GoType())
	}
	return funcSig
}

func (t *TypeName) TextString() string {
	// TODO: more valid checking here on types

	if t == nil {
		panic("TypeName nil")
	}
	if t.EscapedText != "" {
		return t.EscapedText
	}
	if t.Left != nil && t.Right != nil {
		return fmt.Sprintf("%s.%s", t.Left.EscapedText, t.Right.EscapedText)
	}
	return "TypeName:TextString fixme"
}

func (m *Member) Docstring() string {

	lines := JsDocs(m.JsDoc).CommentLines()

	for i, line := range lines {
		lt := strings.TrimSpace(line)
		lt = strings.ReplaceAll(lt, "\r", " ")
		lt = strings.ReplaceAll(lt, "\n", " ")
		if !strings.HasSuffix(lt, ".") {
			lt += "."
		}
		lines[i] = lt
	}
	if m.QuestionToken != nil {
		lines = append(lines, "Optional.")
	}

	cs := fmt.Sprintf("\t// %s - ", util.Title(m.Name.TextString()))

	allLines := strings.Join(lines, " ")
	if len(allLines+cs) > 120 {
		parts := strings.Split(allLines, ". ")
		for i, part := range parts {
			if !strings.HasSuffix(part, ".") {
				parts[i] = part + "."
			}
		}
		sep := "\n\t// "
		return cs + parts[0:1][0] + sep + strings.Join(parts[1:], sep)
	}

	return cs + allLines
}

type JsDocs []*JsDoc

func (m JsDocs) CommentLines() []string {
	var lines []string
	if m != nil {
		for _, j := range m {
			switch ct := j.Comment.(type) {
			case string:
				lines = append(lines, ct)
			case []string:
				lines = append(lines, ct...)
			}
		}
	}
	return lines
}
