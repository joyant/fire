package fire

import (
    "fmt"
    "strings"
)

type inToken struct {
    literal string
    literalValue []string
}

func (t *inToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    literalValue = t.literalValue
    for _, e := range t.literalValue {
        if v == e {
            qualified = true
            break
        }
    }
    return
}

func (t *inToken) TokenType() TokenType {
    return "in"
}

func (t *inToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return MsgFormat(fmt.Sprintf("${0}必须在[%s]内", strings.Join(t.literalValue, ",")))
    } else if lg == LangEN {
        return MsgFormat(fmt.Sprintf("${0} must in [%s]", strings.Join(t.literalValue, ",")))
    }
    return ""
}

func (t *inToken) ParseLiteral(literal string) error {
    slice := strings.Split(literal, ",")
    var s []string
    for _, v := range slice {
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

func (t *inToken) Clone() Token {
    c := *t
    c.literalValue = make([]string, 0)
    return &c
}
