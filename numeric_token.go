package fire

import (
    "fmt"
)

type numericToken struct {
    literal string
}

func (t *numericToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return isNumeric(v), nil, nil
}

func (t *numericToken) TokenType() TokenType {
    return "numeric"
}

func (t *numericToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是数字"
    } else if lg == LangEN {
        return "${0} must be numeric"
    }
    return ""
}

func (t *numericToken) ParseLiteral(literal string) error {
    return nil
}

func (t *numericToken) Clone() Token {
    c := *t
    return &c
}



