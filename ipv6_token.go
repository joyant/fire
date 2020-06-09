package fire

import (
    "fmt"
    "regexp"
)

var ipv6Compile *regexp.Regexp

func init()  {
    var err error
    ipv6Compile, err = regexp.Compile(`^([\w]{1,4}:){7}[\w]{1,4}$`)
    if err != nil {
        panic(err)
    }
}

type ipv6Token struct {}

func (t *ipv6Token) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return ipv6Compile.MatchString(v), nil, nil
}

func (t *ipv6Token) TokenType() TokenType {
    return "ipv6"
}

func (t *ipv6Token) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是ipv6格式"
    } else if lg == LangEN {
        return "${0} must be ipv6 format"
    }
    return ""
}

func (t *ipv6Token) ParseLiteral(literal string) error {
    return nil
}

func (t *ipv6Token) Clone() Token {
    c := *t
    return &c
}
