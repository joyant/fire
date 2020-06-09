package fire

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "unicode/utf8"
)

type lengthToken struct {
    literal string
    literalValue int
}

func (t *lengthToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    typ := reflect.TypeOf(value).String()
    if typ == "string" {
        return utf8.RuneCountInString(value.(string)) == t.literalValue, []string{t.literal}, nil
    } else if typ == "[]string" {
        return len(value.([]string)) == t.literalValue, []string{t.literal}, nil
    } else {
        return false, nil, fmt.Errorf("%s's value must be string or []string", t.TokenType())
    }
}

func (t *lengthToken) TokenType() TokenType {
    return "length"
}

func (t *lengthToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}长度在必须是${1}"
    } else if lg == LangEN {
        return "${0}'s length must be ${1}"
    }
    return ""
}

func (t *lengthToken) ParseLiteral(literal string) error {
    literal = strings.TrimSpace(literal)
    i, err := strconv.Atoi(literal)
    if err != nil {
        return fmt.Errorf("%s's literal %s must be integer", t.TokenType(), literal)
    }
    t.literalValue = i
    t.literal = literal
    return nil
}

func (t *lengthToken) Clone() Token {
    c := *t
    return &c
}
