package fire

import (
    "fmt"
    "strings"
)

type aliasToken struct {
    literal string
}

func (t *aliasToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    qualified = true
    return
}

func (t *aliasToken) TokenType() TokenType {
    return "alias"
}

func (t *aliasToken) I18nMsgFormat(lg Lang) MsgFormat {
    return ""
}

func (t *aliasToken) ParseLiteral(literal string) error {
    literal = strings.TrimSpace(literal)
    if literal == "" {
        return fmt.Errorf("%s's alias ca not be whitespace", t.TokenType())
    }
    t.literal = literal
    return nil
}

func (t *aliasToken) Clone() Token {
    c := *t
    return &c
}

func (t *aliasToken) Alias() string {
    return t.literal
}
