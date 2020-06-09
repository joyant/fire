package fire

func isNumeric(s string) bool {
    dot := false
    for k, v := range s {
        if k == 0 || k == len(s)-1 {
            if v < '0' || v > '9' {
                return false
            }
        } else {
            if v == '.' {
                if dot {
                    return false
                }
                dot = true
            } else if v < '0' || v > '9' {
                return false
            }
        }
    }
    return true
}

func isInteger(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, v := range s {
        if v < '0' || v > '9' {
            return false
        }
    }
    return true
}

func isAlpha(s string) bool {
    for _, v := range s {
        if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) {
            return false
        }
    }
    return true
}

//isAlphaNumeric判断传入的字符串是否只包含字母和数字
//函数不认为小数点是数字，遇到小数点会返回false
func isAlphaNumeric(s string) bool {
    for _, v := range s {
        if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) {
            return false
        }
    }
    return true
}
