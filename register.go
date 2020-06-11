package fire

func RegisterToken(token Token)  {
    tokenMap[token.TokenType()] = token
}

func RegisterI18nDataKey(dataKey DataKey, lang map[Lang]string)  {
    if i18nDataKeyMap[dataKey] == nil {
        i18nDataKeyMap[dataKey] = make(map[Lang]string)
    }
    for k, v := range lang {
        i18nDataKeyMap[dataKey][k] = v
    }
}

func RegisterI18nDataKeyMap(m map[DataKey]map[Lang]string)  {
    for k, v := range m {
        RegisterI18nDataKey(k, v)
    }
}

func RegisterMsgFormat(dataKey DataKey, msg map[Lang]MsgFormat)  {
    if _, ok := msgMap[dataKey]; !ok {
        msgMap[dataKey] = make(map[Lang]MsgFormat)
    }
    for k, v := range msg {
        msgMap[dataKey][k] = v
    }
}