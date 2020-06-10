package fire

import (
    "fmt"
)

type integerToken struct {}

func (t *integerToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return isInteger(v), nil, nil
}

func (t *integerToken) TokenType() TokenType {
    return "integer"
}

func (t *integerToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是整数"
    } else if lg == LangEN {
        return "${0} must be integer"
    }
    return ""
}

func (t *integerToken) ParseLiteral(literal string) error {
    return nil
}

func (t *integerToken) Clone() Token {
    c := *t
    return &c
}

type intToken struct {}

func (t *intToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return isInteger(v), nil, nil
}

func (t *intToken) TokenType() TokenType {
    return "int"
}

func (t *intToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是整数"
    } else if lg == LangEN {
        return "${0} must be integer"
    }
    return ""
}

func (t *intToken) ParseLiteral(literal string) error {
    return nil
}

func (t *intToken) Clone() Token {
    c := *t
    return &c
}