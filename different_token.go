package fire

import (
    "fmt"
    "reflect"
)

type differentToken struct {
    literal string
}

func (t *differentToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    qualified, err = t.equals(value, data[t.literal])
    qualified = !qualified
    return
}

func (t *differentToken)equals(a, b interface{}) (bool, error) {
    at := reflect.TypeOf(a).String()
    bt := reflect.TypeOf(b).String()
    if at == "string" && bt == "string" {
        return a.(string) == b.(string), nil
    } else if at == "[]string" && bt == "[]string" {
        as := a.([]string)
        bs := b.([]string)
        if len(as) != len(bs) {
            return false, nil
        }
        for k := range as {
            if as[k] != bs[k] {
                return false, nil
            }
        }
        return true, nil
    } else {
        return false, fmt.Errorf("type of different's both sides must be string or []string")
    }
}

func (t *differentToken) TokenType() TokenType {
    return "different"
}

func (t *differentToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return MsgFormat(fmt.Sprintf("${0}必须和#{%s}不同", t.literal))
    } else if lg == LangEN {
        return MsgFormat(fmt.Sprintf("${0} must different from #{%s}", t.literal))
    }
    return ""
}

func (t *differentToken) ParseLiteral(literal string) error {
    t.literal = literal
    return nil
}

func (t *differentToken) Clone() Token {
    c := *t
    return &c
}

