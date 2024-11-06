package parser

import (
    "github.com/andeya/ameda"
    "reflect"
    "unsafe"
)

func getValue(s any, fieldName string) any {
    v, is := s.(reflect.Value)
    if !is {
        v = reflect.ValueOf(s)
    }
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    v = ameda.DereferenceInterfaceValue(v)
    v = ameda.DereferenceValue(v)
    field := v.FieldByName(fieldName)
    if field.IsValid() {
        if field.CanInterface() {
            return field.Interface()
        } else {
            return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
        }
    }
    return nil
}

func convertIntToType(i int64, kind reflect.Kind) any {
    switch kind {
    case reflect.Int:
        return int(i)
    case reflect.Int8:
        return int8(i)
    case reflect.Int16:
        return int16(i)
    case reflect.Int32:
        return int32(i)
    case reflect.Int64:
        return i
    case reflect.Uint:
        return uint(i)
    case reflect.Uint8:
        return uint8(i)
    case reflect.Uint16:
        return uint16(i)
    case reflect.Uint32:
        return uint32(i)
    case reflect.Uint64:
        return uint64(i)
    default:
        panic("unhandled default case")
    }
}
