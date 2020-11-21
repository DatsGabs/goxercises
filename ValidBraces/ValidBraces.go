package main

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

// ValidBraces detect if braces are valid
func ValidBraces(str string) bool {
	stack := make([]string, 0)
	for _, bracket := range str {
		bracket := string(bracket)
		top := len(stack) - 1
		if len(stack) > 0 && m[stack[top]] == bracket {
			stack = stack[:top]
		} else {
			stack = append(stack, bracket)
		}
	}
	return len(stack) == 0
}

func main() {
	ValidBraces("{}{}{{(())")
}
