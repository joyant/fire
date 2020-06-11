package fire

import (
    "reflect"
    "testing"
)

func TestRegisterMsgFormat(t *testing.T) {
    m1 := map[Lang]MsgFormat{
        LangZH:"姓名",
        LangEN:"name",
    }
    RegisterMsgFormat(DataKey("name"), m1)
    if !reflect.DeepEqual(m1, msgMap["name"]) {
        t.Errorf("expected %v got %v", m1, msgMap["name"])
    }

    m2 := map[Lang]MsgFormat{
        LangZH:"昵称",
        LangEN:"name",
    }
    RegisterMsgFormat(DataKey("name"), m2)
    if !reflect.DeepEqual(m2, msgMap["name"]) {
        t.Errorf("expected %v got %v", m2, msgMap["name"])
    }

    m3 := map[Lang]MsgFormat{
        LangZH:"年龄",
        LangEN:"age",
    }
    RegisterMsgFormat(DataKey("age"), m3)
    expected := map[DataKey]map[Lang]MsgFormat{
        "name":m2,
        "age":m3,
    }
    if !reflect.DeepEqual(expected, msgMap) {
        t.Errorf("expected %v got %v", expected, msgMap)
    }
}

func TestRegisterI18nDataKey(t *testing.T) {
    m1 := map[Lang]string{
        LangZH:"姓名",
        LangEN:"name",
    }
    RegisterI18nDataKey("name", m1)
    expected := map[DataKey]map[Lang]string{
        "name":m1,
    }
    if !reflect.DeepEqual(expected, i18nDataKeyMap) {
        t.Errorf("expected %v got %v", expected, i18nDataKeyMap)
    }
    m2 := map[Lang]string{
        LangZH:"年龄",
        LangEN:"age",
    }
    RegisterI18nDataKey("age", m2)
    expected = map[DataKey]map[Lang]string{
        "name":m1,
        "age":m2,
    }
    if !reflect.DeepEqual(expected, i18nDataKeyMap) {
        t.Errorf("expected %v got %v", expected, i18nDataKeyMap)
    }
}

func TestRegisterI18nDataKeyMap(t *testing.T) {
    m1 := map[DataKey]map[Lang]string{
        "name":{
            LangZH:"姓名",
            LangEN:"name",
        },
        "age":{
            LangZH:"年龄",
            LangEN:"age",
        },
        "favorite":{
            LangZH:"爱好",
            LangEN:"favorite",
        },
    }
    RegisterI18nDataKeyMap(m1)
    m2 := map[DataKey]map[Lang]string{
        "name":{
            LangZH:"姓名",
            LangEN:"nickname",
        },
        "age":{
            LangZH:"年龄",
            LangEN:"age",
        },
        "gender":{
            LangZH:"性别",
            LangEN:"gender",
        },
    }
    RegisterI18nDataKeyMap(m2)
    m2["favorite"] = map[Lang]string{
        LangZH:"爱好",
        LangEN:"favorite",
    }
    if !reflect.DeepEqual(m2, i18nDataKeyMap) {
        t.Errorf("expected %v got %v", m2, i18nDataKeyMap)
    }
}
