package fire

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
    "unicode/utf8"
)

type lengthBetweenToken struct {
    literal string
    literalSlice [2]string
    literalValue [2]int
}

func (t *lengthBetweenToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    typ := reflect.TypeOf(value).String()
    if typ == "string" {
        l := utf8.RuneCountInString(value.(string))
        return l >= t.literalValue[0] && l <= t.literalValue[1], t.literalSlice[:], nil
    } else if typ == "[]string" {
        l := len(value.([]string))
        return l >= t.literalValue[0] && l <= t.literalValue[1], t.literalSlice[:], nil
    } else {
        return false, nil, fmt.Errorf("%s's value must be string or []string", t.TokenType())
    }
}

func (t *lengthBetweenToken) TokenType() TokenType {
    return "lengthBetween"
}

func (t *lengthBetweenToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}长度必须在${1}和${2}之间"
    } else if lg == LangEN {
        return "${0}'s length must between ${1} and ${2}"
    }
    return ""
}

func (t *lengthBetweenToken) ParseLiteral(literal string) error {
    t.literal = literal
    var pure []string
    for k, v := range strings.Split(literal, ",") {
        v = strings.TrimSpace(v)
        if v != "" {
            a, err := strconv.Atoi(v)
            if err != nil {
                return fmt.Errorf("%s's literal %s is not numberic", t.TokenType(), v)
            }
            t.literalValue[k] = a
            pure = append(pure, v)
        }
    }
    if len(pure) != 2 {
        return fmt.Errorf("%s's literval must like 1,2", t.TokenType())
    }
    for k, v := range pure {
        t.literalSlice[k] = v
    }
    return nil
}

func (t *lengthBetweenToken) Clone() Token {
    c := *t
    c.literalSlice = [2]string{}
    c.literalValue = [2]int{}
    return &c
}


