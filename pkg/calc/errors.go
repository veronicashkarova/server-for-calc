package calc

import "errors"

var (
	ErrInvalidExpression = errors.New("неправильное выражение")
	ErrEmptyExpression   = errors.New("пустое выражение")
)