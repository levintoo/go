package forms

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors[field] = append(f.Errors[field], "This field is required")
		}
	}
}

func (f *Form) MaxLength(field string, length int) {
	value := f.Get(field)
	if utf8.RuneCountInString(value) > length {
		f.Errors[field] = append(f.Errors[field], "This field must be less than %d characters")
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
