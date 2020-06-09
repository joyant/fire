package fire

import (
    "fmt"
    "strings"
)

type ipToken struct {}

func (t *ipToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    if strings.Contains(v, ":") {
        return ipv6Compile.MatchString(v), nil, nil
    }
    return ipv4Compile.MatchString(v), nil, nil
}

func (t *ipToken) TokenType() TokenType {
    return "ip"
}

func (t *ipToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是ipv4或ipv6格式"
    } else if lg == LangEN {
        return "${0} must be ipv4 or ipv6 format"
    }
    return ""
}

func (t *ipToken) ParseLiteral(literal string) error {
    return nil
}

func (t *ipToken) Clone() Token {
    c := *t
    return &c
}

