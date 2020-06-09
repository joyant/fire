package fire

import (
    "fmt"
    "strconv"
    "strings"
)

type minToken struct {
    literal string
    literalValue float64
}

func (t *minToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    vv, err := strconv.ParseFloat(v, 64)
    if err != nil {
        return false, []string{v}, nil
    }
    return vv >= t.literalValue, []string{t.literal}, nil
}

func (t *minToken) TokenType() TokenType {
    return "min"
}

func (t *minToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须大于等于${1}"
    } else if lg == LangEN {
        return "${0} must be greater than or equal to ${1}"
    }
    return ""
}

func (t *minToken) ParseLiteral(literal string) error {
    literal = strings.TrimSpace(literal)
    f, err := strconv.ParseFloat(literal, 64)
    if err != nil {
        return err
    }
    t.literalValue = f
    t.literal = literal
    return nil
}

func (t *minToken) Clone() Token {
    c := *t
    return &c
}

