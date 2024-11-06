package validator

import "github.com/joyant/fire/v2/validator/parser"

type Validator interface {
    Validate(value any) error
}

type validator struct {
    tagName string
}

func NewValidator(tagName string) Validator {
    return &validator{
        tagName: tagName,
    }
}

func (v *validator) Validate(value any) error {
    return parser.ParseStruct(value, v.tagName)
}
