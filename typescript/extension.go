package typescript

import (
	"fmt"
	"strings"

	"github.com/mlctrez/goapp-pf/internal/line"
	"github.com/mlctrez/goapp-pf/internal/util"
	"github.com/mlctrez/goapp-pf/scripts"
	"github.com/mlctrez/goapp-pf/typescript/kind"
)

func deHyphenate(text string) string {
	if strings.Contains(text, "-") {
		parts := strings.Split(text, "-")
		for i, part := range parts {
			parts[i] = util.Title(part)
		}
		text = strings.Join(parts, "")
	}
	return text
}

func (n *Name) TextString() string {

	var textString string
	if n.EscapedText != "" {
		textString = n.EscapedText
	} else {
		textString = n.Text
	}
	textString = deHyphenate(textString)

	return textString
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
				txt := "<unknown>"
				if member.Name != nil {
					txt = member.Name.TextString()
				} else {
					if member.Parameters != nil {
						var txtParts []string
						for _, parameter := range member.Parameters {
							txtPart := fmt.Sprintf("%s: %s",
								parameter.Name.TextString(), parameter.Type.GoType())
							txtParts = append(txtParts, txtPart)
						}
						txt = fmt.Sprintf("[%s]", strings.Join(txtParts, ", "))
					} else {
						scripts.NoErr(fmt.Errorf("cannot build member name"), member)
					}
				}
				switch memType.Kind {
				case kind.AnyKeyword:
					mv := fmt.Sprintf("%q:%s", txt, "any")
					memberValues = append(memberValues, mv)
				case kind.BooleanKeyword:
					mv := fmt.Sprintf("%q:%s", txt, "bool")
					memberValues = append(memberValues, mv)
				case kind.StringKeyword:
					mv := fmt.Sprintf("%s:%s", txt, "string")
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
							scripts.NoErr(fmt.Errorf("memType.Types.Kind %v", m.Kind.StringValue()), member)
						}
					}
					memberValues = append(memberValues, fmt.Sprintf("%s:{ %s }", txt, strings.Join(unionParts, ", ")))
				case kind.TypeReference:
					mv := fmt.Sprintf("%q:%s", txt, member.Type.TypeName.TextString())
					memberValues = append(memberValues, mv)
				case kind.LiteralType:
					switch memType.Literal.Kind {
					case kind.StringLiteral:
						memberValues = append(memberValues, fmt.Sprintf("'%s'", memType.Literal.Text))
					default:
						scripts.NoErr(fmt.Errorf("memType.Literal.Kind %v",
							memType.Literal.Kind.StringValue()), member)
					}
				default:
					scripts.NoErr(fmt.Errorf("member.Type.Kind %v", memType.Kind.StringValue()), member)
				}
			}
			return fmt.Sprintf("map[string]string /* { %s } */", strings.Join(memberValues, ", "))
		}
		scripts.NoErr(fmt.Errorf("syntax.TypeLiteral unhandled"), t)
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
			case kind.ParenthesizedType:
				return fmt.Sprintf("[]( %s )", t.ElementType.Type.GoType())
			case kind.FunctionType:
				elType := t.ElementType.Type
				return "[]" + functionSignature(elType.Type, elType.Parameters)
			default:
				scripts.NoErr(fmt.Errorf("t.ElementType.Kind %s",
					t.ElementType.Kind.StringValue()), t)
				return ""
			}
		}
		scripts.NoErr(fmt.Errorf("t.ElementType ==nil"), t)
		return ""

	case kind.ObjectKeyword:
		return "any /* TODO: how to pass additional properties */"
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
					scripts.NoErr(fmt.Errorf("syntax.ArrayType - nill p.ElementType"), p)
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
				case kind.ParenthesizedType:

					up := fmt.Sprintf("[]( %s )", p.ElementType.Type.GoType())
					unionParts = append(unionParts, up)
				default:
					scripts.NoErr(fmt.Errorf("unhandled union array type %v", p.ElementType.Kind.StringValue()), t)
				}
			case kind.TypeReference:
				unionParts = append(unionParts, p.TypeName.TextString())
			case kind.ParenthesizedType:
				unionParts = append(unionParts, p.Type.GoType())
			case kind.AnyKeyword:
				unionParts = append(unionParts, "any")

			default:
				scripts.NoErr(fmt.Errorf("unhandled union Types kind %s", p.Kind.StringValue()), t)
			}
		}
		// TODO: peel of each of these types and write out a separate constant declaration?
		return "any /* " + strings.Join(unionParts, " | ") + " */"
	default:
		scripts.NoErr(fmt.Errorf("%s", t.Kind.StringValue()), t)
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
	var textString string

	if t.EscapedText != "" {
		textString = t.EscapedText
	} else if t.Left != nil && t.Right != nil {
		textString = fmt.Sprintf("%s.%s", t.Left.EscapedText, t.Right.EscapedText)
	} else {
		scripts.NoErr(fmt.Errorf("textstring from TypeName failed"), t)
	}

	return fmt.Sprintf("any // %s ", textString)

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

	declaration := util.Title(m.Name.TextString())
	comment := fmt.Sprintf("%s - %s", declaration, strings.Join(lines, " "))

	parts := line.Parts(100, comment)

	return "\t// " + strings.Join(parts, "\n\t// ")

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
