package line

import "strings"

func Parts(maxLen int, input string) (result []string) {

	if len(input) < maxLen {
		return []string{input}
	}

	parts := strings.Split(input, " ")
	var currentPart string
	for _, part := range parts {
		if currentPart == "" {
			currentPart = part
		} else {
			currentPart += " " + part
		}
		if len(currentPart) > maxLen {
			result = append(result, currentPart)
			currentPart = ""
		}
	}
	if currentPart != "" {
		result = append(result, currentPart)
	}

	return
}
