package fire

import "fmt"

type alphaNumericToken struct {}

func (t *alphaNumericToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    return isAlphaNumeric(v), nil, nil
}

func (t *alphaNumericToken) TokenType() TokenType {
    return "alphaNum"
}

func (t *alphaNumericToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}必须是英文字母或数字"
    } else if lg == LangEN {
        return "${0} must be alpha or numeric"
    }
    return ""
}

func (t *alphaNumericToken) ParseLiteral(literal string) error {
    return nil
}

func (t *alphaNumericToken) Clone() Token {
    c := *t
    return &c
}

