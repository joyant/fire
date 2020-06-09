package fire

import "testing"

func TestURLCompile(t *testing.T)  {
    if !urlCompile.MatchString("http://www.baidu.com") {
        t.Errorf("expected true got false")
    }
    if !urlCompile.MatchString("http://www.baidu.com/fldksajf9432904023kldsfkl") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("https://www.baidu.com/") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("http://baidu.com") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("https://baidu.com") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("https://baidu.com/") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("https://baidu.com/?key=name") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("www.baidu.com") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("www.baidu.com/lkdfsjlkdjsafljsdafl/fsdljfklfklsda/afjdsla") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("baidu.com") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("baidu.com/fs/5678e4kj32.423423?fksd=fiwoef&iowef=23e23") {
        t.Errorf("expected true get false")
    }
    if !urlCompile.MatchString("https://www.baidu.cn/") {
        t.Errorf("expected true get false")
    }
    if urlCompile.MatchString("http:www.baidu.com") {
        t.Errorf("expected false get true")
    }
    if urlCompile.MatchString("http//www.baidu.com") {
        t.Errorf("expected false get true")
    }
    if urlCompile.MatchString("https//www.baidu.com") {
        t.Errorf("expected false get true")
    }
    if urlCompile.MatchString("httpss://www.baidu.com") {
        t.Errorf("expected false get true")
    }
}
