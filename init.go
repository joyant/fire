package fire

var i18nDataKeyMap = make(map[DataKey]map[Lang]string)

var tokenMap = make(map[TokenType]Token)

func init()  {
    RegisterToken(&requiredToken{})
    RegisterToken(&requireToken{})
    RegisterToken(&lengthToken{})
    RegisterToken(&integerToken{})
    RegisterToken(&intToken{})
    RegisterToken(&betweenToken{})
    RegisterToken(&lengthMinToken{})
    RegisterToken(&lengthMaxToken{})
    RegisterToken(&lengthBetweenToken{})
    RegisterToken(&minToken{})
    RegisterToken(&maxToken{})
    RegisterToken(&boolToken{})
    RegisterToken(&booleanToken{})
    RegisterToken(&inToken{})
    RegisterToken(&notInToken{})
    RegisterToken(&emailToken{})
    RegisterToken(&urlToken{})
    RegisterToken(&numericToken{})
    RegisterToken(&alphaToken{})
    RegisterToken(&alphaNumericToken{})
    RegisterToken(&regexpToken{})
    RegisterToken(&dateToken{})
    RegisterToken(&containsToken{})
    RegisterToken(&equalsToken{})
    RegisterToken(&differentToken{})
    RegisterToken(&ipv4Token{})
    RegisterToken(&ipv6Token{})
    RegisterToken(&ipToken{})
}
