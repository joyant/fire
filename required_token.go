package fire

import (
    "strings"
)

type requiredToken struct {
    literal string
    literalValue string
}

func (t *requiredToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    switch value.(type) {
    case string:
        v := value.(string)
        return strings.TrimSpace(v) != "", nil, nil
    case []string:
        v := value.([]string)
        return len(v) > 0, nil, nil
    default:
        return false, nil, nil
    }
}

func (t *requiredToken) TokenType() TokenType {
    return "required"
}

func (t *requiredToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}是必填项"
    } else if lg == LangEN {
        return "${0} is required"
    }
    return ""
}

func (t *requiredToken) ParseLiteral(literal string) error {
    t.literal = literal
    t.literalValue = literal
    return nil
}

func (t *requiredToken) Clone() Token {
    c := *t
    return &c
}

type requireToken struct {
    literal string
    literalValue string
}

func (t *requireToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    switch value.(type) {
    case string:
        v := value.(string)
        return strings.TrimSpace(v) != "", nil, nil
    case []string:
        v := value.([]string)
        return len(v) > 0, nil, nil
    default:
        return false, nil, nil
    }
}

func (t *requireToken) TokenType() TokenType {
    return "require"
}

func (t *requireToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}是必填项"
    } else if lg == LangEN {
        return "${0} is required"
    }
    return ""
}

func (t *requireToken) ParseLiteral(literal string) error {
    t.literal = literal
    t.literalValue = literal
    return nil
}

func (t *requireToken) Clone() Token {
    c := *t
    return &c
}
