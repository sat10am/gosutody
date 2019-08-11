package calc

const plusOperator = "+"
const minusOperator = "-"
const multiplyOperator = "*"
const divideOperator = "/"

func Calc(operand1 int, operand2 int, operator string) int {
	switch operator {
	case plusOperator:
		return operand1 + operand2
	case minusOperator:
		return operand1 - operand2
	case multiplyOperator:
		return operand1 * operand2
	case divideOperator:
		return operand1 / operand2
	}

	return -1
}
