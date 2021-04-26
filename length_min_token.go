package fire

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "unicode/utf8"
)

type lengthMinToken struct {
    literal string
    literalValue int
}

func (t *lengthMinToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    typ := reflect.TypeOf(value).String()
    if typ == "string" {
        return utf8.RuneCountInString(value.(string)) >= t.literalValue, []string{t.literal}, nil
    } else if typ == "[]string" {
        return len(value.([]string)) >= t.literalValue, []string{t.literal}, nil
    } else {
        return false, nil, fmt.Errorf("%s's value must be string or []string", t.TokenType())
    }
}

func (t *lengthMinToken) TokenType() TokenType {
    return "lengthMin"
}

func (t *lengthMinToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}长度必须大于等于${1}"
    } else if lg == LangEN {
        return "${0} must be less than or equal to ${1}"
    }
    return ""
}

func (t *lengthMinToken) ParseLiteral(literal string) error {
    t.literal = strings.TrimSpace(literal)
    i, err := strconv.Atoi(literal)
    if err != nil {
        return err
    }
    t.literalValue = i
    return nil
}

func (t *lengthMinToken) Clone() Token {
    c := *t
    return &c
}

