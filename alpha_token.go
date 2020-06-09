package fire

import "fmt"

type alphaToken struct {}

func (t *alphaToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return isAlpha(v), nil, nil
}

func (t *alphaToken) TokenType() TokenType {
    return "alpha"
}

func (t *alphaToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是英文字母"
    } else if lg == LangEN {
        return "${0} must be alpha"
    }
    return ""
}

func (t *alphaToken) ParseLiteral(literal string) error {
    return nil
}

func (t *alphaToken) Clone() Token {
    c := *t
    return &c
}


