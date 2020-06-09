package fire

import (
    "testing"
)

func TestIsNumber(t *testing.T)  {
    if !isNumeric("123") {
        t.Errorf("expected true got false")
    }
    if isNumeric("12.3.4") {
        t.Errorf("expected false got true")
    }
    if !isNumeric("12.56789") {
        t.Errorf("expeceted true got false")
    }
    if isNumeric("567d") {
        t.Errorf("expected false got true")
    }
    if isNumeric("汉字") {
        t.Errorf("expeced false got true")
    }
}

func TestIsInteger(t *testing.T)  {
    if isInteger("abc") {
        t.Errorf("expected false got true")
    }
    if !isInteger("34567") {
        t.Errorf("expected true got false")
    }
    if isInteger("567.5678") {
        t.Errorf("expected false got true")
    }
}

func TestIsAlpha(t *testing.T)  {
    if !isAlpha("abc") {
        t.Errorf("expected true got false")
    }
    if isAlpha("ab c") {
        t.Errorf("expected false got true")
    }
    if isAlpha("2323") {
        t.Errorf("expected false got true")
    }
}

func TestIsAlphaNumeric(t *testing.T)  {
    if !isAlphaNumeric("abc") {
        t.Errorf("expected true got false")
    }
    if !isAlphaNumeric("abc123") {
        t.Errorf("expected true got false")
    }
    if isAlphaNumeric("ab c") {
        t.Errorf("expected false got true")
    }
    if !isAlphaNumeric("2323") {
        t.Errorf("expected false got true")
    }
}
