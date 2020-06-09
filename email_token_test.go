package fire

import "testing"

func TestEmailTokenEmail(t *testing.T) {
    if !emailCompile.MatchString("tom@163.com") {
        t.Errorf("expected true got false")
    }
    if emailCompile.MatchString("1234") {
        t.Errorf("expected false got true")
    }
    if emailCompile.MatchString("abc") {
        t.Errorf("expected false got true")
    }
    if emailCompile.MatchString("@qq.cn") {
        t.Errorf("expected false got true")
    }
    if emailCompile.MatchString("1234567890123456789012345678901234567890@163.com") {
        t.Errorf("expected false got true")
    }
    if emailCompile.MatchString("jerry@qq.1234567890abc") {
        t.Errorf("expected false got true")
    }
    if !emailCompile.MatchString("tom-_@qq.com") {
        t.Errorf("expected true got false")
    }
}
