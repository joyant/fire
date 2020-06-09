package fire

import (
    "fmt"
    "regexp"
)

var ipv4Compile *regexp.Regexp

func init()  {
    var err error
    ipv4Compile, err = regexp.Compile(`^(\d{1,3}\.){3}\d{1,3}$`)
    if err != nil {
        panic(err)
    }
}

type ipv4Token struct {}

func (t *ipv4Token) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return ipv4Compile.MatchString(v), nil, nil
}

func (t *ipv4Token) TokenType() TokenType {
    return "ipv4"
}

func (t *ipv4Token) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是ipv4格式"
    } else if lg == LangEN {
        return "${0} must be ipv4 format"
    }
    return ""
}

func (t *ipv4Token) ParseLiteral(literal string) error {
    return nil
}

func (t *ipv4Token) Clone() Token {
    c := *t
    return &c
}

