package fire

import (
    "fmt"
    "regexp"
)

var emailCompile *regexp.Regexp

func init()  {
    var err error
    emailCompile, err = regexp.Compile(`^[\w-_]{1,30}@[\w]{1,10}\.\w{1,10}$`)
    if err != nil {
        panic(err)
    }
}

type emailToken struct {}

func (t *emailToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return emailCompile.MatchString(v), nil, nil
}

func (t *emailToken) TokenType() TokenType {
    return "email"
}

func (t *emailToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是邮箱格式"
    } else if lg == LangEN {
        return "${0} must be email format"
    }
    return ""
}

func (t *emailToken) ParseLiteral(literal string) error {
    return nil
}

func (t *emailToken) Clone() Token {
    c := *t
    return &c
}

