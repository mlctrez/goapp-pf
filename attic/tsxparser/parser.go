package tsxparser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type TsxProps struct {
	Name  string
	Props []*Property
	prop  string
}

func (p *TsxProps) line(line string) {
	subMatch := leadingSpaceRegex.FindStringSubmatch(line)
	// this is a var
	if len(subMatch) == 4 {
		// top level prop
		if len(subMatch[1]) == 2 {
			var prop *Property
			p.prop = subMatch[2]
			// single line prop
			if strings.HasSuffix(line, ";") {
				val := strings.Trim(subMatch[3], "; ")
				prop = &Property{
					Name:  p.prop,
					Value: []*ValueMap{{Name: "single", Values: []string{val}}},
				}
				p.prop = ""
			}
			// multi line prop
			if strings.HasSuffix(line, "{") {
				prop = &Property{Name: subMatch[2], Value: []*ValueMap{}}
			}
			p.Props = append(p.Props, prop)
			return
		}
		if len(subMatch[1]) == 4 {
			// multi line prop
			var currentProp *Property
			for _, prop := range p.Props {
				if prop.Name == p.prop {
					currentProp = prop
				}
			}
			if currentProp == nil {
				panic("sub map but could not find parent " + line)
			}
			currentProp.current = subMatch[2]
			currentProp.line(subMatch[3])

		}
	} else {
		if line == "  };" {
			return
		}
		var currentProp *Property
		for _, prop := range p.Props {
			if prop.Name == p.prop {
				currentProp = prop
			}
		}
		if currentProp == nil {
			panic("sub map but could not find parent " + line)
		}
		currentProp.line(line)

	}

}

type Property struct {
	Name    string
	Value   []*ValueMap
	current string
}

func (p *Property) line(line string) {
	if p.current == "" {
		panic("property line without setting current")
	}
	var existingValue *ValueMap
	for _, valueMap := range p.Value {
		if valueMap.Name == p.current {
			existingValue = valueMap
		}
	}
	if existingValue == nil {
		existingValue = &ValueMap{Name: p.current}
		p.Value = append(p.Value, existingValue)
	}
	existingValue.line(line)

}

type ValueMap struct {
	Name   string
	Values []string
}

func (m *ValueMap) line(line string) {
	parts := strings.Split(strings.Trim(line, "|;"), "|")
	unQuoted := make([]string, len(parts))
	for i, part := range parts {
		unQuoted[i] = strings.Trim(part, "' ")
	}
	for _, s := range unQuoted {
		if s != "" {
			m.Values = append(m.Values, s)
		}
	}

}

func (p *TsxProps) exportLine(el string) {
	p.Name = strings.Split(el, " ")[2]
}

var leadingSpaceRegex = regexp.MustCompile(`(\s+)'?(\w+)'?\?:(.*)`)

func Parse(path string) (*TsxProps, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() { _ = open.Close() }()

	scanner := bufio.NewScanner(open)

	txsProps := &TsxProps{}

	for scanner.Scan() {
		st := scanner.Text()
		if strings.HasPrefix(st, `export interface`) {
			txsProps.exportLine(st)
			for scanner.Scan() {
				st = scanner.Text()
				if st == "}" {
					break
				}
				lt := strings.TrimSpace(st)
				if strings.HasPrefix(lt, "/*") && strings.HasSuffix(lt, "*/") {
					continue
				}
				txsProps.line(st)
			}
		}
	}

	return txsProps, nil
}
