package lexer

import (
	"strconv"
	"strings"
)

// Tokenize the given segment
func Tokenize(current string) []interface{} {
	fields := strings.Fields(current)
	tokens := splitBrackets(fields)
	identified := mapper(tokens)
	return identified
}

func splitBrackets(fields []string) []string {
	temp := []string{}
	for _, item := range fields {
		item = strings.TrimSpace(item)

		if item[0] == '(' {
			temp = append(temp, "(", item[1:])
		} else if item[len(item)-1] == ')' {
			temp = append(temp, item[:len(item)-1], ")")
		} else {
			temp = append(temp, item)
		}
	}

	return temp
}

func mapper(tokens []string) []interface{} {
	var table []interface{}

	for _, item := range tokens {
		if item == "" {
			continue
		}

		if item == "(" || item == ")" {
			table = append(table, item)
			continue
		}

		if val, ok := Optable[item]; ok {
			table = append(table, val)
			continue
		}

		switch {
		case chkInt(item):
			val, _ := strconv.Atoi(item)
			table = append(table, val)
		case chkFloat(item):
			val, _ := strconv.ParseFloat(item, 64)
			table = append(table, val)
		default:
			panic("invalid input " + item)
		}
	}

	return table
}

func chkInt(item string) bool {
	if _, err := strconv.Atoi(item); err == nil {
		return true
	}
	return false
}

func chkFloat(item string) bool {
	if _, err := strconv.ParseFloat(item, 64); err == nil {
		return true
	} else {
		return false
	}
}
