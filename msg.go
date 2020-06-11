package fire

var msgMap = make(map[DataKey]map[Lang]MsgFormat)

type MsgFormat string

func getMsgFormat(dataKey DataKey, lang Lang) MsgFormat {
    if msgMap[dataKey] != nil {
        return msgMap[dataKey][lang]
    }
    return ""
}
