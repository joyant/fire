package fire

import (
    "fmt"
    "strings"
)

type containsToken struct {
    literal string
    caseSensitive bool
}

func (t *containsToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    if t.caseSensitive {
        return strings.Contains(v, t.literal), []string{t.literal}, nil
    }
    return strings.Contains(strings.ToLower(v), strings.ToLower(t.literal)), []string{t.literal}, nil
}

func (t *containsToken) TokenType() TokenType {
    return "contains"
}

func (t *containsToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须包含${1}"
    } else if lg == LangEN {
        return "${0} must contains ${1}"
    }
    return ""
}

func (t *containsToken) ParseLiteral(literal string) error {
    if strings.HasSuffix(literal, "/i") { // "Tom/i" 表示不区分大小写，不带/i默认区分大小写
        t.literal = literal[0:len(literal)-2]
        t.caseSensitive = false
    } else {
        t.literal = literal
        t.caseSensitive = true
    }
    return nil
}

func (t *containsToken) Clone() Token {
    c := *t
    return &c
}
