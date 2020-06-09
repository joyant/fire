package fire

import (
    "fmt"
    "strconv"
)

type boolToken struct {}

func (t *boolToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    _, e := strconv.ParseBool(v)
    qualified = e == nil
    return
}

func (t *boolToken) TokenType() TokenType {
    return "bool"
}

func (t *boolToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是布尔类型"
    } else if lg == LangEN {
        return "${0}'s value must be boolean type"
    }
    return ""
}

func (t *boolToken) ParseLiteral(literal string) error {
    return nil
}

func (t *boolToken) Clone() Token {
    c := *t
    return &c
}

type booleanToken struct {}

func (t *booleanToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    _, e := strconv.ParseBool(v)
    qualified = e == nil
    return
}

func (t *booleanToken) TokenType() TokenType {
    return "boolean"
}

func (t *booleanToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是布尔类型"
    } else if lg == LangEN {
        return "${0}'s value must be boolean type"
    }
    return ""
}

func (t *booleanToken) ParseLiteral(literal string) error {
    return nil
}

func (t *booleanToken) Clone() Token {
    c := *t
    return &c
}
