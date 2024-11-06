package parser

import (
    "reflect"
)

func ParseStruct(s any, tagName string) error {
    v, isReflectValue := s.(reflect.Value)
    if !isReflectValue {
        v = reflect.ValueOf(s)
    }
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    t := v.Type()
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fieldValue := v.Field(i)
        if fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Struct {
            fieldValue = fieldValue.Elem()
        }
        if fieldValue.Kind() == reflect.Struct {
            err := ParseStruct(v.Field(i), tagName)
            if err != nil {
                return err
            }
        }
        tag := field.Tag.Get(tagName)
        // tagName does not exist
        if len(tag) == 0 {
            continue
        }
        err := Parse(tag, s, field.Name)
        if err != nil {
            return err
        }
    }
    return nil
}
