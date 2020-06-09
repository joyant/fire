package fire

type Lang string

var DefaultLang = LangZH

const (
    LangZH Lang = "zh" // 汉语
    LangEN Lang = "en" // 英语
    LangDE Lang = "de" // 德语
    LangFR Lang = "fr" // 法语
    LangRU Lang = "ru" // 俄语
    LangES Lang = "es" // 西班牙语
    LangJA Lang = "ja" // 日语
    LangAR Lang = "ar" // 阿拉伯语
    LangKR Lang = "kr" // 韩语
    LangPT Lang = "pt" // 葡萄牙语
)

func getLang(lgs ...Lang) Lang {
    if len(lgs) > 0 {
        return lgs[0]
    }
    return DefaultLang
}

