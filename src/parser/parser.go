package parser

// Parse : parse the tokens recursively
func Parse(tokens []interface{}) (result []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	/*
	* - find inner most unit to Parse
	* - parse recursively
	 */

	/*
	* - if this token unit contains nothing else,
	* - it is the inner most unit, process it and return result list
	* - otherwise recursively call Parse and eval
	 */

	result, err = _parse(tokens)

	return
}

func _parse(token []interface{}) (result []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	op := token[1]
	args := token[2:(len(token) - 2)]

	result, err = eval(op, args)
	return
}

func eval(op interface{}, args []interface{}) (result []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	return
}
