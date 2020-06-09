package fire

import (
    "fmt"
    "regexp"
)

var urlCompile *regexp.Regexp

func init()  {
    var err error
    urlCompile, err = regexp.Compile(`^(http[s]?://)?([\w-_\.]+\.)?[\w-_]+\.[a-z]{1,15}.*`)
    if err != nil {
        panic(err)
    }
}

type urlToken struct {}

func (t *urlToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return urlCompile.MatchString(v), nil, nil
}

func (t *urlToken) TokenType() TokenType {
    return "url"
}

func (t *urlToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是网址格式"
    } else if lg == LangEN {
        return "${0} must be url format"
    }
    return ""
}

func (t *urlToken) ParseLiteral(literal string) error {
    return nil
}

func (t *urlToken) Clone() Token {
    c := *t
    return &c
}

