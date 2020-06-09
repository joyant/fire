package fire

import (
    "fmt"
    "strings"
)

type notInToken struct {
    literal string
    literalValue []string
}

func (t *notInToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    qualified = true
    literalValue = t.literalValue
    for _, e := range t.literalValue {
        if v == e {
            qualified = false
            break
        }
    }
    return
}

func (t *notInToken) TokenType() TokenType {
    return "notIn"
}

func (t *notInToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return MsgFormat(fmt.Sprintf("${0}不能在[%s]内", strings.Join(t.literalValue, ",")))
    } else if lg == LangEN {
        return MsgFormat(fmt.Sprintf("${0} must not in [%s]", strings.Join(t.literalValue, ",")))
    }
    return ""
}

func (t *notInToken) ParseLiteral(literal string) error {
    var s []string
    for _, v := range strings.Split(literal, ",") {
        v = strings.TrimSpace(v)
        if v != "" {
            s = append(s, v)
        }
    }
    if len(s) == 0 {
        return fmt.Errorf("%s can not be nil", t.TokenType())
    }
    t.literalValue = s
    t.literal = literal
    return nil
}

func (t *notInToken) Clone() Token {
    c := *t
    c.literalValue = make([]string, 0)
    return &c
}
