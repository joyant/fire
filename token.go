package fire

type TokenType string

type Token interface {
    Evaluate(value DataValue, data Data) (qualified bool, literalValue []string, err error)
    TokenType() TokenType
    I18nMsgFormat(Lang) MsgFormat
    ParseLiteral(literal string) error
    Clone() Token
}
