package lexer

import (
	"errors"
	"strconv"
	"strings"
)

// Tokenize the given segment
func Tokenize(current string) (identified []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			identified = nil
		}
	}()

	var tokens []string

	fields := strings.Fields(current)

	tokens, err = splitBrackets(fields)
	if err != nil {
		panic(err)
	}

	identified, err = mapper(tokens)
	if err != nil {
		panic(err)
	}

	return
}

func splitBrackets(fields []string) (temp []string, err error) {
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

	return
}

func mapper(tokens []string) (table []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			table = nil
		}
	}()

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
			panic(errors.New("invalid input " + item))
		}
	}

	return
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
	}
	return false
}
