package lexer

import "fmt"

// define table of supported operations

func _addOp(lhs, rhs float64) float64 {
	return lhs + rhs
}

func _subOp(lhs, rhs float64) float64 {
	return lhs - rhs
}

func _mulOp(lhs, rhs float64) float64 {
	return lhs * rhs
}

func _divOp(lhs, rhs float64) float64 {
	return lhs / rhs
}

func _modOp(lhs, rhs int64) int64 {
	return lhs % rhs
}

func _printOp(val interface{}) {
	fmt.Println(val)
}

// Optable : map of supported operations
var Optable = map[string]interface{}{
	"+":     _addOp,
	"-":     _subOp,
	"*":     _mulOp,
	"/":     _divOp,
	"%":     _modOp,
	"print": _printOp,
}
