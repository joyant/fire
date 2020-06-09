package fire

func RegisterToken(token Token)  {
    tokenMap[token.TokenType()] = token
}

func RegisterI18nDataKey(dataKey DataKey, lang map[Lang]string)  {
    i18nDataKeyMap[dataKey] = lang
}

func RegisterI18nDataKeyMap(m map[DataKey]map[Lang]string)  {
    for k, v := range m {
        i18nDataKeyMap[k] = v
    }
}