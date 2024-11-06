# Fire: Ignite Your Go Validation

Fire is a blazingly fast, intuitive, and powerful form validation library for Go. Inspired by the transformative power of fire, this library helps you refine your data, burning away invalid inputs and leaving only the pure, validated information your application needs.

## 🔥 Features

- **Intuitive as Go Itself**: Fire's validation rules are as simple and expressive as Go code. No additional syntax or special rules to learn. If you can write Go, you can write Fire validation rules.

- **Powerful Validation Capabilities**: Fire can handle a wide range of data types and validation scenarios. From basic string length and pattern checks to complex cross-field and list validations, Fire has got you covered.

## 🚀 Quick Start

Get up and running with Fire in no time with this simple example:

```go
type Contact struct {
    Name         string `fire:"len($) > 0 && len($) < 20"`
    Gender       string `fire:"$ == 'male' || $ == 'female'"`
    Age          int    `fire:"$ > 0 && $ <= 120"`
    Marriageable bool   `fire:"this.Age >= (this.Gender == 'male' ? 22:20)"`
    EnName       string `fire:"regexp('[a-zA-Z]{1,30}')"`
    Score        []int  `fire:"len($) == 6"`
}

func main() {
    c := &Contact{
        Name:         "Tom",
        Gender:       "male",
        Age:          23,
        Marriageable: true,
        EnName:       "tom",
        Score:        []int{99, 100, 90, 100, 88, 90},
    }
    vm := NewValidator("fire")
    if err := vm.Validate(c); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Contact validated successfully!")
}
```

## 📘 Syntax Guide

Fire's syntax is as intuitive as Go itself. Here's a quick guide to the key concepts:

- **$**: The placeholder for the field value being validated.
- **len($)**: Checks the length of the value.
- **$ == 'value'**: Checks if the value equals a specific value.
- **this.Field**: Accesses another field in the same struct.
- **regexp(pattern)**: Checks if the value matches a regular expression pattern.

With Fire, form validation becomes a breeze. Ignite your Go validation today with Fire!
