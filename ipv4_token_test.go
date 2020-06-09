package fire

import "testing"

func TestIpv4Compile(t *testing.T)  {
    if !ipv4Compile.MatchString("192.168.0.1") {
        t.Errorf("expected true got false")
    }
    if ipv4Compile.MatchString("1923.168.0.1") {
        t.Errorf("expected false got true")
    }
    if ipv4Compile.MatchString("192.168.0.") {
        t.Errorf("expected false got true")
    }
    if ipv4Compile.MatchString("192.168.0.a") {
        t.Errorf("expected false got true")
    }
}
