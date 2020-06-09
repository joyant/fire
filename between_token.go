package fire

import (
    "fmt"
    "strconv"
    "strings"
)

type betweenToken struct {
    literal string
    literalSlice [2]string
    literalValue [2]float64
}

func (t *betweenToken) Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error) {
    if value == nil {
        qualified = true
        return
    }
    v, ok := value.(string)
    if !ok {
        err = fmt.Errorf("%s's value must be string", t.TokenType())
        return
    }
    vv, err := strconv.ParseFloat(v, 64)
    if err != nil {
        return false, t.literalSlice[:], nil
    }
    return vv >= t.literalValue[0] && vv <= t.literalValue[1], t.literalSlice[:], nil
}

func (t *betweenToken) TokenType() TokenType {
    return "between"
}

func (t *betweenToken) I18nMsgFormat(lg Lang) MsgFormat {
    if lg == LangZH {
        return "${0}大小必须在${1}和${2}之间"
    } else if lg == LangEN {
        return "${0} must between ${1} and ${2}"
    }
    return ""
}

func (t *betweenToken) ParseLiteral(literal string) error {
    t.literal = literal
    var pure []string
    slice := strings.Split(literal, ",")
    for k, v := range slice {
        v = strings.TrimSpace(v)
        if v != "" {
            a, err := strconv.ParseFloat(v, 64)
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

func (t *betweenToken) Clone() Token {
    c := *t
    c.literalSlice = [2]string{}
    c.literalValue = [2]float64{}
    return &c
}


