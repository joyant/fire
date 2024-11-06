package parser

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type hair struct {
    color  string `vd:"$ == 'red'"`
    length int    `vd:"$ > 5"`
}

type person struct {
    hair  *hair
    grade byte   `vd:"(this.age > 10 ? 5:4) + 1 >= 6"`
    age   int    `vd:"$ >= 18; msg:sprintf('age %d is underage', $)"`
    name  string `vd:"$ == 'tom'"`
}

type contact struct {
    Email string `fire:"regexp('[a-zA-Z0-9_]{1,10}@[a-zA-Z0-9]{1,5}\\.(com|cn|edu)')"`
}

type people struct {
    Gender   string
    Age      int
    CanMarry bool `fire:"this.Age >= (this.Gender == 'male' ? 22:20)"`
}

type name struct {
    Gender string `fire:"($ == 'male') || ($ == 'female')"`
}

func TestParseStruct(t *testing.T) {
    p := person{
        hair: &hair{
            color:  "red",
            length: 10,
        },
        grade: 5,
        age:   21,
        name:  "tom",
    }
    err := ParseStruct(&p, "vd")
    assert.NoError(t, err)
}

func TestParseStruct2(t *testing.T) {
    p := person{
        hair: &hair{
            color:  "red",
            length: 10,
        },
        grade: 5,
        age:   17,
        name:  "tom",
    }
    err := ParseStruct(&p, "vd")
    assert.Error(t, err)
}

func TestParseStruct3(t *testing.T) {
    c := contact{
        Email: "tom@xyz.com",
    }
    err := ParseStruct(&c, "fire")
    assert.NoError(t, err)
}

func TestParseStruct4(t *testing.T) {
    p := &people{
        Gender:   "male",
        Age:      22,
        CanMarry: true,
    }
    err := ParseStruct(p, "fire")
    assert.NoError(t, err)
}

func TestParseStruct5(t *testing.T) {
    p := &name{
        Gender: "male2",
    }
    err := ParseStruct(p, "fire")
    assert.NoError(t, err)
}
