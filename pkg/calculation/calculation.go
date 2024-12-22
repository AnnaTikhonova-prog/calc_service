package calculation

import (
	"errors"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	return parseExpression(&expression)
}

func parseExpression(expr *string) (float64, error) {
	result, err := parseTerm(expr)
	if err != nil {
		return 0, err
	}
	for len(*expr) > 0 {
		op := (*expr)[0]
		if op != '+' && op != '-' {
			break
		}
		*expr = (*expr)[1:]

		term, err := parseTerm(expr)
		if err != nil {
			return 0, err
		}

		if op == '+' {
			result += term
		} else {
			result -= term
		}
	}
	return result, nil
}

func parseTerm(expr *string) (float64, error) {
	result, err := parseFactor(expr)
	if err != nil {
		return 0, err
	}

	for len(*expr) > 0 {
		op := (*expr)[0]
		if op != '*' && op != '/' {
			break
		}
		*expr = (*expr)[1:]

		factor, err := parseFactor(expr)
		if err != nil {
			return 0, err
		}

		if op == '*' {
			result *= factor
		} else {
			if factor == 0 {
				return 0, errors.New("деление на ноль")
			}
			result /= factor
		}
	}

	return result, nil
}

func parseFactor(expr *string) (float64, error) {
	if len(*expr) == 0 {
		return 0, errors.New("преждевременный конец выражения...")
	}

	if (*expr)[0] == '(' {
		*expr = (*expr)[1:]
		result, err := parseExpression(expr)
		if err != nil {
			return 0, err
		}

		if len(*expr) == 0 || (*expr)[0] != ')' {
			return 0, errors.New("несовпадающие скобки")
		}
		*expr = (*expr)[1:]
		return result, nil
	}

	i := 0
	for i < len(*expr) && (('0' <= (*expr)[i] && (*expr)[i] <= '9') || (*expr)[i] == '.') {
		i++
	}
	numberStr := (*expr)[:i]
	*expr = (*expr)[i:]

	value, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, errors.New("невалидное число")
	}

	return value, nil
}
