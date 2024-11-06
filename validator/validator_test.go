package validator

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type contact struct {
    Name         string `fire:"len($) > 0 && len($) < 20"`
    Gender       string `fire:"$ == 'male' || $ == 'female'"`
    Age          int    `fire:"$ > 0 && $ <= 120"`
    Marriageable bool   `fire:"this.Age >= (this.Gender == 'male' ? 22:20)"`
    EnName       string `fire:"regexp('[a-zA-Z]{1,30}')"`
    Score        []int  `fire:"len($) == 6"`
}

func TestNewValidator(t *testing.T) {
    c := &contact{
        Name:         "Tom",
        Gender:       "male",
        Age:          23,
        Marriageable: true,
        EnName:       "tom",
        Score:        []int{99, 100, 90, 100, 88, 90},
    }
    vm := NewValidator("fire")
    err := vm.Validate(c)
    assert.NoError(t, err)
}
