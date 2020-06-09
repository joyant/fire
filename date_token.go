package fire

import (
    "fmt"
    "time"
)

type dateToken struct {
    literal string
}

func (t *dateToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    _, e := time.Parse(t.literal, v)
    return e == nil, []string{t.literal}, nil
}

func (t *dateToken) TokenType() TokenType {
    return "date"
}

func (t *dateToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是日期格式"
    } else if lg == LangEN {
        return "${0} must be date format"
    }
    return ""
}

func (t *dateToken) ParseLiteral(literal string) error {
    if literal == "" {
        literal = "2006-01-02 15:04:05"
    }
    t.literal = literal
    return nil
}

func (t *dateToken) Clone() Token {
    c := *t
    return &c
}
