package apierr

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[size:]
}

func FromFieldError(fe validator.FieldError) ProblemDetailError {
	field := lowerFirst(fe.Field())
	var (
		code   = fe.Tag()
		detail string
		params map[string]string
	)
	switch fe.Tag() {
	case "required":
		detail = "required"
	case "email":
		detail = "invalid email address"
	case "min":
		params = map[string]string{"min": fe.Param()}
		detail = fmt.Sprintf("must be at least %s characters", fe.Param())
	case "max":
		params = map[string]string{"max": fe.Param()}
		detail = fmt.Sprintf("must be at most %s characters", fe.Param())
	case "eqfield":
		other := lowerFirst(fe.Param())
		params = map[string]string{"otherField": other}
		detail = fmt.Sprintf("%s must match %s", field, other)
	default:
		detail = fmt.Sprintf("%s is invalid", field)
	}
	return ProblemDetailError{Detail: detail, Pointer: "/" + field, Code: code, Params: params}
}
