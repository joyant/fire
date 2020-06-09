package fire

import (
    "fmt"
    "regexp"
)

type regexpToken struct {
    literal string
    literalValue *regexp.Regexp
}

func (t *regexpToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return t.literalValue.MatchString(v), []string{t.literal}, nil
}

func (t *regexpToken) TokenType() TokenType {
    return "regexp"
}

func (t *regexpToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须符合正则表达式${1}"
    } else if lg == LangEN {
        return "${0} must match regexp ${1}"
    }
    return ""
}

func (t *regexpToken) ParseLiteral(literal string) error {
    reg, err := regexp.Compile(literal)
    if err != nil {
        return fmt.Errorf("compile regexp %s error %s", literal, err.Error())
    }
    t.literalValue = reg
    t.literal = literal
    return nil
}

func (t *regexpToken) Clone() Token {
    c := *t
    c.literalValue = nil
    return &c
}
