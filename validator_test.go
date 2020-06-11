package fire

import (
    "testing"
)

func TestBytes2Int(t *testing.T)  {
    i, err := bytes2Int([]byte{'1', '2', '3'})
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if i != 123 {
        t.Errorf("expected %d got %d", 123, i)
    }
    i2, err := bytes2Int([]byte{'1', '9', '0', '0', '0', '0'})
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if i != 123 {
        t.Errorf("expected %d got %d", 123, i)
    }
    if i2 != 190000 {
        t.Errorf("expected %d got %d", 190000, i2)
    }
    _, err = bytes2Int([]byte{'1', '9', 'a', '0', '0', '0'})
    if err == nil {
        t.Errorf("expected err got nil")
    }
}

func TestReplaceMsg(t *testing.T)  {
    s := replaceMsg("name", MsgFormat("${0}长度必须是${1}"), []string{"5"}, LangZH)
    if s != "name长度必须是5" {
        t.Errorf("expected %s got %s", "名称长度必须是5", s)
    }
    s = replaceMsg("名称", MsgFormat("${0}长度必须是${1}"), []string{"5"}, LangZH)
    if s != "名称长度必须是5" {
        t.Errorf("expected %s got %s", "名称长度必须是5", s)
    }
    s = replaceMsg("年龄", MsgFormat("${0}大小必须在${1}~${2}之间"), []string{"1", "130"}, LangZH)
    if s != "年龄大小必须在1~130之间" {
        t.Errorf("expected %s got %s", "年龄大小必须在1~130之间", s)
    }
    s = replaceMsg("年龄", MsgFormat("${0}大小必须等于#{age2}"), nil, LangZH)
    if s != "年龄大小必须等于age2" {
        t.Errorf("expected %s got %s", "年龄大小必须等于age2", s)
    }
    RegisterI18nDataKey("age2", map[Lang]string{LangZH:"年龄2"})
    s = replaceMsg("年龄", MsgFormat("${0}大小必须等于#{age2}"), nil, LangZH)
    if s != "年龄大小必须等于年龄2" {
        t.Errorf("expected %s got %s", "年龄大小必须等于年龄2", s)
    }
}

func TestNewLength(t *testing.T) {
    rule := Rule{
        "name":"length:5",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangZH: "姓名",
    })
    v, err := New(rule, LangZH)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"tom",
    })
    if qualified {
        t.Errorf("expected false got true")
    }
    if err == nil {
        t.Errorf("expected err got nil")
    }
    rule["name"] = "length:3"
    v, err = New(rule, LangZH)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err = v.Validate(map[string]interface{}{
        "name":"jim",
    })
    if !qualified {
        t.Errorf("expected true got false")
    }
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
}

func TestNewRequired(t *testing.T) {
    rule := Rule{
        "name":"required",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":[]string{},
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name is required"
        if err.Error() != expected {
            t.Errorf("expected %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
    qualified, err = v.Validate(map[string]interface{}{
        "name":"",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name is required"
        if err.Error() != expected {
            t.Errorf("expected %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestNewRequired2(t *testing.T) {
    rule := Rule{
        "name":"required",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"abc",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestInteger(t *testing.T)  {
    rule := Rule{
        "age":"integer",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"11",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestInteger2(t *testing.T)  {
    rule := Rule{
        "age":"integer",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"abc",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "age must be integer"
        if err.Error() != expected {
            t.Errorf("expected %s get %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestBetween(t *testing.T)  {
    rule := Rule{
        "age":"between:1, 130",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"abc",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "age must between 1 and 130"
        if err.Error() != expected {
            t.Errorf("expected %s get %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestBetween2(t *testing.T)  {
    rule := Rule{
        "age":"between:1, 130",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"123",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestLengthMin(t *testing.T) {
    rule := Rule{
        "address":"lengthMin:5",
    }
    RegisterI18nDataKey("address", map[Lang]string{
        LangEN: "address",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "address":"street A number B",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestLengthMin2(t *testing.T) {
    rule := Rule{
        "address":"lengthMin:5",
    }
    RegisterI18nDataKey("address", map[Lang]string{
        LangEN: "address",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "address":"A",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "address must be less than or equal to 5"
        if err.Error() != expected {
            t.Errorf("expected %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestLengthBetween(t *testing.T) {
    rule := Rule{
        "name":"lengthBetween:5, 10",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"tom",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name's length must between 5 and 10"
        if err.Error() != expected {
            t.Errorf("expected %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestLengthBetween2(t *testing.T) {
    rule := Rule{
        "name":"lengthBetween:5, 10",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"tommy",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestMin(t *testing.T) {
    rule := Rule{
        "age":"min:5",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"6",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestMin2(t *testing.T) {
    rule := Rule{
        "age":"min:5",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"4",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "age must be greater than or equal to 5"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestMax(t *testing.T) {
    rule := Rule{
        "age":"max:5",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"1",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestMax2(t *testing.T) {
    rule := Rule{
        "age":"max:5",
    }
    RegisterI18nDataKey("age", map[Lang]string{
        LangEN: "age",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "age":"6",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "age must be less than or equal to 5"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestBool(t *testing.T) {
    rule := Rule{
        "installed":"bool",
    }
    RegisterI18nDataKey("installed", map[Lang]string{
        LangEN: "installed",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "installed":"6",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "installed's value must be boolean type"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestBool2(t *testing.T) {
    rule := Rule{
        "installed":"bool",
    }
    RegisterI18nDataKey("installed", map[Lang]string{
        LangEN: "installed",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "installed":"true",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestIn(t *testing.T) {
    rule := Rule{
        "name":"in:tom, jim, jerry",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"6",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name must in [tom,jim,jerry]"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestIn2(t *testing.T) {
    rule := Rule{
        "name":"in:tom, jim, jerry",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"jerry",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestNotIn(t *testing.T) {
    rule := Rule{
        "name":"notIn:tom, jim, jerry",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"tom",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name must not in [tom,jim,jerry]"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestNotIn2(t *testing.T) {
    rule := Rule{
        "name":"notIn:tom, jim, jerry",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"white",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestEmail(t *testing.T) {
    rule := Rule{
        "mail":"email",
    }
    RegisterI18nDataKey("mail", map[Lang]string{
        LangEN: "mail",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "mail":"tom",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "mail must be email format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestEmail2(t *testing.T) {
    rule := Rule{
        "mail":"email",
    }
    RegisterI18nDataKey("mail", map[Lang]string{
        LangEN: "mail",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"white_head@163.com",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestIpv4(t *testing.T) {
    rule := Rule{
        "site_ip":"ipv4",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"192.168.0.",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "site_ip must be ipv4 format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestIpv42(t *testing.T) {
    rule := Rule{
        "site_ip":"ipv4",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"192.168.0.1",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestURL(t *testing.T) {
    rule := Rule{
        "my_url":"url",
    }
    RegisterI18nDataKey("my_url", map[Lang]string{
        LangEN: "my_url",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "my_url":"https//www.baidu.com",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "my_url must be url format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestURL2(t *testing.T) {
    rule := Rule{
        "my_url":"url",
    }
    RegisterI18nDataKey("my_url", map[Lang]string{
        LangEN: "my_url",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "my_url":"https://www.baidu.cn",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestNumeric(t *testing.T) {
    rule := Rule{
        "money":"numeric",
    }
    RegisterI18nDataKey("money", map[Lang]string{
        LangEN: "money",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "money":"123.567890a",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "money must be numeric"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestNumeric2(t *testing.T) {
    rule := Rule{
        "money":"numeric",
    }
    RegisterI18nDataKey("money", map[Lang]string{
        LangEN: "money",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "money":"123.4564390239",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestNumeric3(t *testing.T) {
    rule := Rule{
        "money":"numeric",
    }
    RegisterI18nDataKey("money", map[Lang]string{
        LangEN: "money",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "money":"123345678",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestNumeric4(t *testing.T) {
    rule := Rule{
        "money":"numeric",
    }
    RegisterI18nDataKey("money", map[Lang]string{
        LangEN: "money",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "money":"123.5678 90",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "money must be numeric"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestAlpha(t *testing.T) {
    rule := Rule{
        "first_name":"alpha",
    }
    RegisterI18nDataKey("first_name", map[Lang]string{
        LangEN: "first_name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "first_name":"Smith2",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "first_name must be alpha"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestAlpha2(t *testing.T) {
    rule := Rule{
        "first_name":"alpha",
    }
    RegisterI18nDataKey("first_name", map[Lang]string{
        LangEN: "first_name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "first_name":"Smith",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestAlphaNumeric(t *testing.T) {
    rule := Rule{
        "account":"alphaNum",
    }
    RegisterI18nDataKey("account", map[Lang]string{
        LangEN: "account",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "account":"afklsd32321*&",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "account must be alpha or numeric"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestAlphaNumeric2(t *testing.T) {
    rule := Rule{
        "account":"alphaNum",
    }
    RegisterI18nDataKey("account", map[Lang]string{
        LangEN: "account",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "account":"abc123",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestRegexp(t *testing.T) {
    rule := Rule{
        "password":"regexp:^[a-zA-Z][\\w_-]{5,20}$",
    }
    RegisterI18nDataKey("password", map[Lang]string{
        LangEN: "password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "password":"123456",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "password must match regexp ^[a-zA-Z][\\w_-]{5,20}$"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestRegexp2(t *testing.T) {
    rule := Rule{
        "password":"regexp:^[a-zA-Z][\\w_-]{5,20}$",
    }
    RegisterI18nDataKey("password", map[Lang]string{
        LangEN: "password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "password":"abc1234-56def",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestDate(t *testing.T) {
    rule := Rule{
        "birthday":"date",
    }
    RegisterI18nDataKey("birthday", map[Lang]string{
        LangEN: "birthday",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "birthday":"2006",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "birthday must be date format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestDate2(t *testing.T) {
    rule := Rule{
        "birthday":"date",
    }
    RegisterI18nDataKey("birthday", map[Lang]string{
        LangEN: "birthday",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "birthday":"2006-09-09 09:09:09",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestDate3(t *testing.T) {
    rule := Rule{
        "birthday":"date:2006-01-02",
    }
    RegisterI18nDataKey("birthday", map[Lang]string{
        LangEN: "birthday",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "birthday":"2006",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "birthday must be date format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestDate4(t *testing.T) {
    rule := Rule{
        "birthday":"date:2006",
    }
    RegisterI18nDataKey("birthday", map[Lang]string{
        LangEN: "birthday",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "birthday":"2010",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestContains(t *testing.T) {
    rule := Rule{
        "name":"contains:mi",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"tommy",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name must contains mi"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestContains2(t *testing.T) {
    rule := Rule{
        "name":"contains:my",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"tommy",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestContains3(t *testing.T) {
    rule := Rule{
        "name":"contains:mi/i",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"TOMMY",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "name must contains mi"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestContains4(t *testing.T) {
    rule := Rule{
        "name":"contains:my/i",
    }
    RegisterI18nDataKey("name", map[Lang]string{
        LangEN: "name",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "name":"TOMMY",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestEquals(t *testing.T) {
    rule := Rule{
        "password":"equals:password2",
    }
    RegisterI18nDataKey("password", map[Lang]string{
        LangEN: "password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "password":"test",
        "password2":"test1",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "password must equals to password2"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestEquals2(t *testing.T) {
    rule := Rule{
        "password":"equals:password2",
    }
    RegisterI18nDataKey("password", map[Lang]string{
        LangEN: "password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "password":"test",
        "password2":"test",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestEquals3(t *testing.T) {
    rule := Rule{
        "password":"equals:password2",
    }
    RegisterI18nDataKey("password", map[Lang]string{
        LangZH:"密码",
    })
    RegisterI18nDataKey("password2", map[Lang]string{
        LangZH:"第二次输入密码",
    })
    v, err := New(rule, LangZH)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "password":"test",
        "password2":"test1",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "密码必须和第二次输入密码相同"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestDifferent(t *testing.T) {
    rule := Rule{
        "new_password":"different:old_password",
    }
    RegisterI18nDataKey("new_password", map[Lang]string{
        LangEN: "new_password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "new_password":"test",
        "old_password":"test",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "new_password must different from old_password"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestDifferent2(t *testing.T) {
    rule := Rule{
        "new_password":"different:old_password",
    }
    RegisterI18nDataKey("new_password", map[Lang]string{
        LangEN: "new_password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "new_password":"test",
        "old_password":"test2",
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestDifferent3(t *testing.T) {
    rule := Rule{
        "new_password":"different:old_password",
    }
    RegisterI18nDataKey("new_password", map[Lang]string{
        LangEN: "new_password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "new_password":[]string{"test"},
        "old_password":[]string{"test2"},
    })
    if err != nil {
        t.Errorf("expected nil got err %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestDifferent4(t *testing.T) {
    rule := Rule{
        "new_password":"different:old_password",
    }
    RegisterI18nDataKey("new_password", map[Lang]string{
        LangEN: "new_password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "new_password":[]string{"test"},
        "old_password":[]string{"test"},
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "new_password must different from old_password"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestDifferent5(t *testing.T) {
    rule := Rule{
        "new_password":"different:old_password",
    }
    RegisterI18nDataKey("new_password", map[Lang]string{
        LangEN: "new_password",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "new_password":[]string{"test"},
        "old_password":"test",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "type of different's both sides must be string or []string"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestIpv6(t *testing.T) {
    rule := Rule{
        "site_ip":"ipv6",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"192.168.0.",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "site_ip must be ipv6 format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestIpv62(t *testing.T) {
    rule := Rule{
        "site_ip":"ipv6",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"0:0:123:0:8:800:123C:132A",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestIp(t *testing.T) {
    rule := Rule{
        "site_ip":"ip",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"192.168.0.",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "site_ip must be ipv4 or ipv6 format"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestIp2(t *testing.T) {
    rule := Rule{
        "site_ip":"ip",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"0:0:123:0:8:800:123C:132A",
    })
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    if !qualified {
        t.Errorf("expected true got false")
    }
}

func TestIp3(t *testing.T) {
    rule := Rule{
        "site_ip":"ip",
    }
    RegisterI18nDataKey("site_ip", map[Lang]string{
        LangEN: "site_ip",
    })
    RegisterMsgFormat("site_ip", map[Lang]MsgFormat{
        LangEN:"${0} is not ip",
    })
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{
        "site_ip":"192.168.0.",
    })
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "site_ip is not ip"
        if expected != err.Error() {
            t.Errorf("expeceted %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestAlias(t *testing.T) {
    rule := Rule{
        "name":"required|alias:class_name",
    }
    v, err := New(rule, LangEN)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    }
    qualified, err := v.Validate(map[string]interface{}{})
    if err == nil {
        t.Errorf("expected err got nil")
    } else {
        expected := "class_name is required"
        if expected != err.Error() {
            t.Errorf("expected %s got %s", expected, err.Error())
        }
    }
    if qualified {
        t.Errorf("expected false got true")
    }
}

func TestCamel2Underline(t *testing.T)  {
    if s := camel2Underline("name"); s != "name" {
        t.Errorf("expected %s got %s", "name", s)
    }
    if s := camel2Underline("FirstName"); s != "first_name" {
        t.Errorf("expected %s got %s", "first_name", s)
    }
    if s := camel2Underline("First_Name"); s != "first_name" {
        t.Errorf("expected %s got %s", "first_name", s)
    }
}

func TestAll(t *testing.T)  {
    RegisterI18nDataKeyMap(map[DataKey]map[Lang]string{
        "name":             {LangZH: "名字"},
        "age":              {LangZH: "年龄"},
        "mail":             {LangZH: "邮箱"},
        "gender":           {LangZH: "性别"},
        "address":          {LangZH: "地址"},
        "score":            {LangZH: "分数"},
        "adult":            {LangZH: "是否成年人"},
        "country":          {LangZH: "国家"},
        "province":         {LangZH: "省份"},
        "blog":             {LangZH: "博客"},
        "deposit":          {LangZH: "存款"},
        "firstName":        {LangZH: "首名称"},
        "username":         {LangZH: "用户名"},
        "birthday":         {LangZH: "生日"},
        "favorite":         {LangZH: "爱好"},
        "password":         {LangZH: "密码"},
        "confirm_password": {LangZH: "第二次输入密码"},
        "account":          {LangZH: "账号"},
        "blog_ip4":         {LangZH: "博客IP4"},
        "blog_ip6":         {LangZH: "博客IP6"},
        "favorite_site":    {LangZH: "最爱网站"},
    })
    v, err := New(Rule{
        "name":             "required|lengthBetween:5,10",
        "age":              "integer|between:1,130",
        "mail":             "require|email",
        "gender":           "require|int|length:1",
        "address":          "require|lengthMin:1|lengthMax:20",
        "score":            "require|min:0|max:100",
        "adult":            "require|bool",
        "country":          "require|in:China,England",
        "province":         "require|notIn:BeiJing,ShangHai",
        "blog":             "url",
        "deposit":          "require|numeric",
        "firstName":        "require|alpha",
        "username":         "require|alphaNum",
        "birthday":         "require|date:2006-01-02",
        "favorite":         "require|contains:ball/i",
        "password":         "require|regexp:^\\w{6,10}$",
        "confirm_password": "require|equals:password",
        "account":          "require|different:username",
        "blog_ip4":         "ipv4",
        "blog_ip6":         "require|ipv6",
        "favorite_site":    "required",
        "ignore_key":        "",
    }, LangZH)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    } else {
        pass, err := v.Validate(map[string]interface{}{
            "name":             "tommy",
            "age":              "19",
            "mail":             "123456@qq.com",
            "gender":           "1",
            "address":          "xinghua street no. 1",
            "score":            "95",
            "adult":            "false",
            "country":          "China",
            "province":         "xian",
            "blog":             "https://www.my-blog.com",
            "deposit":          "100",
            "firstName":        "Tom",
            "username":         "tom123456",
            "birthday":         "2010-01-02",
            "favorite":         "FOOTBALL",
            "password":         "123456abc",
            "confirm_password": "123456abc",
            "account":          "tom123456@",
            "blog_ip4":         "192.168.0.1",
            "blog_ip6":         "0:0:123:0:8:800:123C:132A",
            "favorite_site":    "192.168.0.2",
            "ignore_key":       "abc",
        })
        if err != nil {
            t.Errorf("expected nil got %v", err)
        }
        if !pass {
            t.Errorf("expected true got false")
        }
    }
}

func TestAll2(t *testing.T)  {
    RegisterI18nDataKeyMap(map[DataKey]map[Lang]string{
        "name":             {LangZH: "名字"},
        "age":              {LangZH: "年龄"},
        "mail":             {LangZH: "邮箱"},
        "gender":           {LangZH: "性别"},
        "address":          {LangZH: "地址"},
        "score":            {LangZH: "分数"},
        "adult":            {LangZH: "是否成年人"},
        "country":          {LangZH: "国家"},
        "province":         {LangZH: "省份"},
        "blog":             {LangZH: "博客"},
        "deposit":          {LangZH: "存款"},
        "firstName":        {LangZH: "首名称"},
        "username":         {LangZH: "用户名"},
        "birthday":         {LangZH: "生日"},
        "favorite":         {LangZH: "爱好"},
        "password":         {LangZH: "密码"},
        "confirm_password": {LangZH: "第二次输入密码"},
        "account":          {LangZH: "账号"},
        "blog_ip4":         {LangZH: "博客IP4"},
        "blog_ip6":         {LangZH: "博客IP6"},
        "favorite_site":    {LangZH: "最爱网站"},
    })
    v, err := New(Rule{
        "name":             "required|lengthBetween:5,10",
        "age":              "integer|between:1,130",
        "mail":             "require|email",
        "gender":           "require|int|length:1",
        "address":          "require|lengthMin:1|lengthMax:20",
        "score":            "require|min:0|max:100",
        "adult":            "require|bool",
        "country":          "require|in:China,England",
        "province":         "require|notIn:BeiJing,ShangHai",
        "blog":             "url",
        "deposit":          "require|numeric",
        "FirstName":        "require|alpha",
        "username":         "require|alphaNum",
        "birthday":         "require|date:2006-01-02",
        "favorite":         "require|contains:ball/i",
        "password":         "require|regexp:^\\w{6,10}$",
        "confirm_password": "require|equals:password",
        "account":          "require|different:username",
        "blog_ip4":         "ipv4",
        "blog_ip6":         "require|ipv6",
        "favorite_site":    "required",
        "ignore_key":        "",
    }, LangZH)
    if err != nil {
        t.Errorf("expected nil got %v", err)
    } else {
        type MyData struct {
            Name string
            Age int
            Mail string
            Gender int
            Address string
            Score float64
            Adult bool
            Country string
            Province string
            Blog string
            Deposit float64
            FirstName string
            Username string
            Birthday string
            Favorite string
            Password string
            Confirm_Password string
            Account string
            BlogIPV4 string `fire:"blog_ip4"`
            BlogIPV6 string `fire:"blog_ip6"`
            FavoriteSite string
            IgnoreKey string
        }
        myData := MyData{
            Name:             "tommy",
            Age:              19,
            Mail:             "123456@qq.com",
            Gender:           1,
            Address:          "xinghua street no. 1",
            Score:            95,
            Adult:            false,
            Country:          "China",
            Province:         "xian",
            Blog:             "https://www.my-blog.com",
            Deposit:          100,
            FirstName:        "Tom",
            Username:         "tom123456",
            Birthday:         "2010-01-02",
            Favorite:         "FOOTBALL",
            Password:         "123456abc",
            Confirm_Password: "123456abc",
            Account:          "tom123456@",
            BlogIPV4:         "192.168.0.1",
            BlogIPV6:         "0:0:123:0:8:800:123C:132A",
            FavoriteSite:     "192.168.0.2",
            IgnoreKey:        "abc",
        }
        pass, err := v.Validate(&myData)
        if err != nil {
            t.Errorf("expected nil got %v", err)
        }
        if !pass {
            t.Errorf("expected true got false")
        }
        pass, err = v.Validate(myData)
        if err != nil {
            t.Errorf("expected nil got %v", err)
        }
        if !pass {
            t.Errorf("expected true got false")
        }
    }
}

func BenchmarkNew(b *testing.B) {
    RegisterI18nDataKeyMap(map[DataKey]map[Lang]string{
        "name":             {LangZH: "名字"},
        "age":              {LangZH: "年龄"},
        "mail":             {LangZH: "邮箱"},
        "gender":           {LangZH: "性别"},
        "address":          {LangZH: "地址"},
        "score":            {LangZH: "分数"},
        "adult":            {LangZH: "是否成年人"},
        "country":          {LangZH: "国家"},
        "province":         {LangZH: "省份"},
        "blog":             {LangZH: "博客"},
        "deposit":          {LangZH: "存款"},
        "firstName":        {LangZH: "首名称"},
        "username":         {LangZH: "用户名"},
        "birthday":         {LangZH: "生日"},
        "favorite":         {LangZH: "爱好"},
        "password":         {LangZH: "密码"},
        "confirm_password": {LangZH: "第二次输入密码"},
        "account":          {LangZH: "账号"},
        "blog_ip4":         {LangZH: "博客IP4"},
        "blog_ip6":         {LangZH: "博客IP6"},
        "favorite_site":    {LangZH: "最爱网站"},
    })
    v, err := New(Rule{
        "name":             "required|lengthBetween:5,10",
        "age":              "integer|between:1,130",
        "mail":             "require|email",
        "gender":           "require|int|length:1",
        "address":          "require|lengthMin:1|lengthMax:20",
        "score":            "require|min:0|max:100",
        "adult":            "require|bool",
        "country":          "require|in:China,England",
        "province":         "require|notIn:BeiJing,ShangHai",
        "blog":             "url",
        "deposit":          "require|numeric",
        "firstName":        "require|alpha",
        "username":         "require|alphaNum",
        "birthday":         "require|date:2006-01-02",
        "favorite":         "require|contains:ball/i",
        "password":         "require|regexp:^\\w{6,10}$",
        "confirm_password": "require|equals:password",
        "account":          "require|different:username",
        "blog_ip4":         "ipv4",
        "blog_ip6":         "require|ipv6",
        "favorite_site":    "required",
        "ignore_key":        "",
    }, LangZH)
    if err != nil {
        b.Errorf("expected nil got %v", err)
    } else {
        for i := 0; i < b.N; i++ {
            pass, err := v.Validate(map[string]interface{}{
                "name":             "tommy",
                "age":              "19",
                "mail":             "123456@qq.com",
                "gender":           "1",
                "address":          "xinghua street no. 1",
                "score":            "95",
                "adult":            "false",
                "country":          "China",
                "province":         "xian",
                "blog":             "https://www.my-blog.com",
                "deposit":          "100",
                "firstName":        "Tom",
                "username":         "tom123456",
                "birthday":         "2010-01-02",
                "favorite":         "FOOTBALL",
                "password":         "123456abc",
                "confirm_password": "123456abc",
                "account":          "tom123456@",
                "blog_ip4":         "192.168.0.1",
                "blog_ip6":         "0:0:123:0:8:800:123C:132A",
                "favorite_site":    "192.168.0.2",
                "ignore_key":       "abc",
            })
            if err != nil {
                b.Errorf("expected nil got %v", err)
            }
            if !pass {
                b.Errorf("expected true got false")
            }
        }
    }
}