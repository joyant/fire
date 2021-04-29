package fire

import (
    "bytes"
    "errors"
    "fmt"
    "reflect"
    "strings"
)

type Validator interface {
    Validate(data interface{}) (qualified bool, err error)
    Rule() Rule
}

func New(rule Rule, lgs ...Lang) (Validator, error) {
    f := &fireValidator{
        parser: &fireParser{},
        rule:   rule,
        lang:   getLang(lgs...),
        token:  make(map[DataKey][]Token),
    }
    if err := f.parseRule(); err != nil {
        return nil, err
    }
    return f, nil
}

type fireValidator struct {
    parser parser
    rule   Rule
    token  map[DataKey][]Token
    lang   Lang
}

func (f *fireValidator)Rule() Rule {
    return f.rule
}

func (f *fireValidator)parseRule() error {
    for k, v := range f.rule {
        tokens, err := f.parser.Parse(v)
        if err != nil {
            return err
        }
        f.token[DataKey(k)] = tokens
    }
    return nil
}

func (f *fireValidator)Validate(data interface{}) (qualified bool, err error) {
    switch reflect.TypeOf(data).Kind() {
    case reflect.Map:
        switch data.(type) {
        case Data:
            return f.validateMap(data.(Data))
        case map[string]interface{}:
            return f.validateMap(data.(map[string]interface{}))
        default:
            return false, errors.New("data need fire.Data type")
        }
    case reflect.Struct, reflect.Ptr:
        return f.validateStruct(data)
    default:
        return false, errors.New("data need fire.Data or structure type")
    }
}

func (f *fireValidator)validateMap(data Data) (qualified bool, err error) {
    for k := range f.rule {
        tokens := f.token[DataKey(k)]
        var key = k
        //如果tokens中有多个别名(alias)，只取第一个
        for _, token := range tokens {
            if a, ok := token.(alias); ok && a.Alias() != "" {
                key = a.Alias()
                break
            }
        }
        var msg MsgFormat
        for _, token := range tokens {
            qualified, literalValue, err := token.Evaluate(DataValue(data[k]), data)
            if err != nil {
                return false, err
            }
            if !qualified {
                msg = getMsgFormat(DataKey(k), f.lang)
                if msg == "" {
                    msg = token.I18nMsgFormat(f.lang)
                }
                if msg == "" {
                    return false, fmt.Errorf("dataKey %s's i18nMsgFormat is not exists", k)
                }
                i18nV := getI18nDataValue(DataKey(key), f.lang)
                if i18nV == "" {
                    i18nV = key
                }
                hint := replaceMsg(i18nV, msg, literalValue, f.lang)
                return false, errors.New(hint)
            }
        }
    }
    return true, nil
}

func (f *fireValidator)tryName(name string) string {
    for _, v := range []string{name, camel2Underline(name), strings.ToLower(name)} {
        if _, ok := f.rule[v]; ok {
            return v
        }
    }
    return ""
}

func (f *fireValidator)makeData(s interface{}) Data {
    t := reflect.TypeOf(s)
    v := reflect.ValueOf(s)
    k := t.Kind()
    var l int
    if k == reflect.Ptr {
        l = t.Elem().NumField()
    } else {
        l = t.NumField()
    }
    data := Data{}
    for i := 0; i < l; i++ {
        var field reflect.StructField
        var iv interface{}
        if k == reflect.Ptr {
            field = t.Elem().Field(i)
            iv = getValue(v.Elem().Field(i).Interface())
        } else {
            field = t.Field(i)
            iv = getValue(v.Field(i).Interface())
        }
        if tag, ok := field.Tag.Lookup(Tag); ok {
            data[tag] = iv
            continue
        }
        name := f.tryName(field.Name)
        if name != "" {
            data[name] = iv
        }
    }
    return data
}

func (f *fireValidator)validateStruct(s interface{}) (qualified bool, err error) {
    return f.validateMap(f.makeData(s))
}

func getI18nDataValue(key DataKey, lang Lang) string {
    ls := i18nDataKeyMap[key]
    if len(ls) == 0 {
        return ""
    }
    if v := ls[lang]; v == "" {
        return string(key)
    } else {
        return v
    }
}

func getValue(v interface{}) interface{} {
    switch v.(type) {
    case string, []string:
        return v
    default:
        if reflect.TypeOf(v).Kind() == reflect.Slice && reflect.ValueOf(v).Len() == 0 {
            return "" // 空数组以空字符串对待
        }
        return fmt.Sprintf("%v", v)
    }
}

func camel2Underline(s string) string {
    buff := bytes.Buffer{}
    for k, v := range s {
        if v >= 'A' && v <= 'Z' {
            if k > 0 && s[k-1] != '_' {
                buff.WriteByte('_')
            }
            buff.WriteByte(byte('a'+v-'A'))
            continue
        }
        buff.WriteByte(byte(v))
    }
    return buff.String()
}

//replaceMsg将含有占位符的字符串替换为有明确意义的字符串
//比如验证字符创长度的BetweenToken，它的msgFormat是"${0}'s length must between ${1} and ${2}"
//${0}表示dataKey的占位符，${1}表示第一个规则字面量占位符，在规则"between:5,10"中，${1}就是5的占位符
//${2}是10的占位符，依次类推;如果i18nKey的值是name，那么返回的提示是"name's length must between 5 and 10"
//如果消息用到了其他的key，比如equals这个token，我们在使用的时候，可能这样写，"password":"equals:confirm_password"
//confirm_password就是"其他key"，用到其他key的占位符要这样写:#{key},在这里我们写成#{confirm_password},
//我们设定一个msgFormat:"${0}必须和#{confirm_password}相同"，当发生错误时候，占位符就会被替换，在中文环境下，
//替换后的字符可能是"密码必须和第二次输入密码相同"
func replaceMsg(i18nKey string, msg MsgFormat, literalValue []string, lang Lang) string {
    buff := bytes.Buffer{}
    var j int
    for i, l := 0, len(msg); i < l; {
        if msg[i] == '$' && i+1<l && msg[i+1] == '{' {
            for j = i+2; j < l && msg[j] != '}'; j++ {}
            if j >= l {
                j = l-1
            }
            if msg[j] == '}' && j>i+2 {
                index, err := bytes2Int([]byte(msg[i+2:j]))
                if err != nil {
                    return err.Error()
                }
                if index == 0 {
                    buff.WriteString(i18nKey)
                } else {
                    if len(literalValue) > index-1 {
                        buff.WriteString(literalValue[index-1])
                    }
                }
                i = j+1
                continue
            }
        } else if msg[i] == '#' && i+1<l && msg[i+1] == '{' {
            for j = i+2; j < l && msg[j] != '}'; j++ {}
            if j >= l {
                j = l-1
            }
            if msg[j] == '}' && j>i+2 {
                i18n := msg[i+2:j]
                i18nV := getI18nDataValue(DataKey(i18n), lang)
                if i18nV != "" {
                    buff.WriteString(i18nV)
                } else {
                    buff.WriteString(string(i18n))
                }
                i = j+1
                continue
            }
        }
        buff.WriteByte(msg[i])
        i++
    }
    return buff.String()
}

//bytes2Int将一个字节数组转为int
//如输入['1', '2', '3']会返回123，bytes2Int并不 检查数组的合法性，
//默认地，它认为传递的参数是经过检查的
func bytes2Int(bs []byte) (int, error) {
    i := 0
    for _, b := range bs {
        if !(b >= '0' && b <= '9') {
            return 0, fmt.Errorf("need numeric got %c", b)
        }
        i = i*10 + int(b - '0')
    }
    return i, nil
}
