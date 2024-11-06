package parser

import (
    "errors"
    "reflect"
)

func toFloat(value any) (f float64, err error) {
    val := reflect.ValueOf(value)
    switch val.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return float64(val.Int()), nil
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return float64(val.Uint()), nil
    case reflect.Float32, reflect.Float64:
        return val.Float(), nil
    default:
        return 0, errors.New("NaN")
    }
}
