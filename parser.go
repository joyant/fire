package fire

import (
    "errors"
    "strings"
)

type parser interface {
    Parse(chaoticRule string) ([]Token, error)
}

type fireParser struct {}

func (f *fireParser)Parse(chaoticRule string) (tokens []Token, err error) {
    if len(chaoticRule) == 0 {
        return
    }
    var (
        token Token
        lastIndex int
    )
    for i, l := 0, len(chaoticRule); i < l; i++ {
        e := chaoticRule[i]
        if e == '|' {
            if i == 0 {
                err = errors.New("rule can not start with | ")
                return
            }
            token, err = f.parse(chaoticRule[lastIndex:i])
            if err != nil {
                return
            }
            tokens = append(tokens, token)
            lastIndex = i + 1
        } else if i == l-1 {
            token, err = f.parse(chaoticRule[lastIndex:])
            if err != nil {
                return
            }
            tokens = append(tokens, token)
        }
    }
    return
}

func (f *fireParser)parse(s string) (Token, error) {
    s = strings.TrimSpace(s)
    if s == "" {
        return nil, errors.New("token can not be white space")
    }
    slice := strings.Split(s, ":")
    for k, v := range slice {
        slice[k] = strings.TrimSpace(v)
    }
    if t, ok := tokenMap[TokenType(slice[0])]; !ok {
        return nil, errors.New("token " + slice[0] + " not exists")
    } else {
        var literal string
        if len(slice) > 1 {
            literal = slice[1]
        }
        c := t.Clone()
        return c, c.ParseLiteral(literal)
    }
}
