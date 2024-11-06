package parser

import (
    "errors"
    "fmt"
    "github.com/joyant/fire/v2/validator/grammar"
    "reflect"
    "regexp"
    "strconv"
    "strings"
)

var trim = strings.TrimSpace

const (
    typeErrorTemplate = "type error: the evaluation of '%v' should yield a %s"
)

type Fn func(args ...any) error

var registerFunctions = make(map[string]Fn)

func RegisterFunction(functionName string, fn Fn) {
    registerFunctions[functionName] = fn
}

type listener struct {
    *grammar.BaseVALIDATORListener
    s    any
    self string
    err  error
}

func newListener(s any, self string) *listener {
    l := &listener{
        s:    s,
        self: self,
    }
    return l
}

func (l *listener) Validate() (err error) {
    return l.err
}

func (l *listener) EnterValidator(ctx *grammar.ValidatorContext) {
    children := ctx.GetChildren()
    var (
        expr *grammar.ExprContext
        msg  *grammar.MsgContext
    )
    for _, child := range children {
        if c, ok := child.(*grammar.ExprContext); ok {
            expr = c
        } else if m, ok := child.(*grammar.MsgContext); ok {
            msg = m
        }
    }
    result, err := l.evalExpr(expr)
    if err == nil && reflect.ValueOf(result).Kind() == reflect.Bool && result.(bool) == false {
        err = fmt.Errorf("expr %s's result is %t", expr.GetText(), false)
    }
    if err != nil {
        if msg != nil {
            if m, e := l.evalMsg(msg); e != nil {
                err = e
            } else if len(m) > 0 {
                err = errors.New(m)
            }
        }
    }
    l.err = err
}

func (l *listener) evalExpr(ctx *grammar.ExprContext) (result any, err error) {
    var (
        andExpr *grammar.AndExprContext
        expr    *grammar.ExprContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.ExprContext); ok {
            expr = c
        } else if a, ok := child.(*grammar.AndExprContext); ok {
            andExpr = a
        }
    }
    andResult, err := l.evalAndExpr(andExpr)
    if err != nil {
        return nil, err
    }
    if expr == nil {
        return andResult, nil
    }
    exprResult, err := l.evalExpr(expr)
    if err != nil {
        return nil, err
    }
    var (
        a bool
        b bool
    )
    switch v := andResult.(type) {
    case bool:
        a = v
    default:
        return nil, fmt.Errorf(typeErrorTemplate, andExpr.GetText(), "boolean")
    }
    switch v := exprResult.(type) {
    case bool:
        b = v
    default:
        return nil, fmt.Errorf(typeErrorTemplate, expr.GetText(), "boolean")
    }
    if !(a || b) {
        return false, fmt.Errorf("either the expression '%s' or '%s' must evaluate to true",
            andExpr.GetText(), expr.GetText())
    }
    return true, nil
}

func (l *listener) evalMsg(ctx *grammar.MsgContext) (result string, err error) {
    sprintfExpr := ctx.Sprintf()
    format := strings.Trim(sprintfExpr.STRING().GetText(), "'")
    sprintArgs := sprintfExpr.AllSprintfArgs()
    var args []any
    for _, a := range sprintArgs {
        exprResult, err := l.evalExpr(a.Expr().(*grammar.ExprContext))
        if err != nil {
            return "", err
        }
        args = append(args, exprResult)
    }
    return fmt.Sprintf(format, args...), nil
}

func (l *listener) evalAndExpr(ctx *grammar.AndExprContext) (result any, err error) {
    var (
        conditionalExpr *grammar.ConditionalExprContext
        andExpr         *grammar.AndExprContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.AndExprContext); ok {
            andExpr = c
        } else if a, ok := child.(*grammar.ConditionalExprContext); ok {
            conditionalExpr = a
        }
    }
    conditionalResult, err := l.evalConditionalExpr(conditionalExpr)
    if err != nil {
        return nil, err
    }
    if andExpr == nil {
        return conditionalResult, nil
    }
    andResult, err := l.evalAndExpr(andExpr)
    if err != nil {
        return nil, err
    }
    var (
        a bool
        b bool
    )
    switch v := conditionalResult.(type) {
    case bool:
        a = v
    default:
        return nil, fmt.Errorf(typeErrorTemplate, conditionalExpr.GetText(), "boolean")
    }
    switch v := andResult.(type) {
    case bool:
        b = v
    default:
        return nil, fmt.Errorf(typeErrorTemplate, andExpr.GetText(), "boolean")
    }
    if !(a && b) {
        return false, fmt.Errorf("the results of both expressions '%s' and '%s' must be true",
            conditionalExpr.GetText(), andExpr.GetText())
    }
    return true, nil
}

func (l *listener) evalConditionalExpr(ctx *grammar.ConditionalExprContext) (result any, err error) {
    var (
        comparingExpr   *grammar.ComparingExprContext
        expr            *grammar.ExprContext
        conditionalExpr *grammar.ConditionalExprContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.ComparingExprContext); ok {
            comparingExpr = c
        } else if a, ok := child.(*grammar.ConditionalExprContext); ok {
            conditionalExpr = a
        } else if b, ok := child.(*grammar.ExprContext); ok {
            expr = b
        }
    }
    comparingResult, err := l.evalComparingExpr(comparingExpr)
    if err != nil {
        //return nil, err
    }
    if expr == nil {
        return comparingResult, nil
    }
    exprResult, err := l.evalExpr(expr)
    if err != nil {
        return nil, err
    }
    conditionalResult, err := l.evalConditionalExpr(conditionalExpr)
    if err != nil {
        return nil, err
    }
    var yes bool
    switch v := comparingResult.(type) {
    case bool:
        yes = v
    default:
        return nil, fmt.Errorf(typeErrorTemplate, comparingExpr.GetText(), "boolean")
    }
    if yes {
        return exprResult, nil
    } else {
        return conditionalResult, nil
    }
}

func (l *listener) evalComparingExpr(ctx *grammar.ComparingExprContext) (result any, err error) {
    var (
        addingExpr      *grammar.AddingExprContext
        compareOperator *grammar.CompareOperatorContext
        comparingExpr   *grammar.ComparingExprContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.AddingExprContext); ok {
            addingExpr = c
        } else if a, ok := child.(*grammar.CompareOperatorContext); ok {
            compareOperator = a
        } else if b, ok := child.(*grammar.ComparingExprContext); ok {
            comparingExpr = b
        }
    }
    addingResult, err := l.evalAddingExpr(addingExpr)
    if err != nil {
        return nil, err
    }
    if compareOperator == nil {
        return addingResult, nil
    }
    compareResult, err := l.evalComparingExpr(comparingExpr)
    if err != nil {
        return nil, err
    }
    // string comparing
    if reflect.ValueOf(addingResult).Kind() == reflect.String && (reflect.ValueOf(compareResult).Kind() == reflect.String) {
        if compareOperator.EQ() != nil {
            if addingResult.(string) != compareResult.(string) {
                return false, fmt.Errorf("field: %s '%s' must eq '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
            }
        } else if compareOperator.NOTEQ() != nil {
            if addingResult.(string) == compareResult.(string) {
                return false, fmt.Errorf("field: %s '%s' must not eq '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
            }
        } else {
            return false, fmt.Errorf("field: %s two string comparing only support eq and not eq", l.self)
        }
        return true, nil
    }
    // number comparing
    a, err := toFloat(addingResult)
    if err != nil {
        return nil, fmt.Errorf(typeErrorTemplate, addingExpr.GetText(), "int or float")
    }
    b, err := toFloat(compareResult)
    if err != nil {
        return nil, fmt.Errorf(typeErrorTemplate, comparingExpr.GetText(), "int or float")
    }
    if compareOperator.GT() != nil {
        if !(a > b) {
            return false, fmt.Errorf("field: %s '%s' must gt '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
        }
    } else if compareOperator.EGT() != nil {
        if !(a >= b) {
            return false, fmt.Errorf("field: %s '%s' must egt '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
        }
    } else if compareOperator.LT() != nil {
        if !(a < b) {
            return false, fmt.Errorf("field: %s '%s' must lt '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
        }
    } else if compareOperator.ELT() != nil {
        if !(a <= b) {
            return false, fmt.Errorf("field: %s '%s' must egt '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
        }
    } else if compareOperator.EQ() != nil {
        if !(a == b) {
            return false, fmt.Errorf("field: %s '%s' must eq '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
        }
    } else if compareOperator.NOTEQ() != nil {
        if !(a != b) {
            return false, fmt.Errorf("field: %s '%s' must not eq '%s'", l.self, addingExpr.GetText(), comparingExpr.GetText())
        }
    } else {
        return false, nil
    }
    return true, nil
}

func (l *listener) evalAddingExpr(ctx *grammar.AddingExprContext) (result any, err error) {
    var (
        addingExpr      *grammar.AddingExprContext
        multiplyingExpr *grammar.MultiplyingExprContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.AddingExprContext); ok {
            addingExpr = c
        } else if a, ok := child.(*grammar.MultiplyingExprContext); ok {
            multiplyingExpr = a
        }
    }
    multiplyingResult, err := l.evalMultiplyingExpr(multiplyingExpr)
    if err != nil {
        return nil, err
    }
    if addingExpr == nil {
        return multiplyingResult, nil
    }
    addingResult, err := l.evalAddingExpr(addingExpr)
    if err != nil {
        return nil, err
    }
    a, err := toFloat(multiplyingResult)
    if err != nil {
        return nil, fmt.Errorf(typeErrorTemplate, multiplyingExpr.GetText(), "int or float")
    }
    b, err := toFloat(addingResult)
    if err != nil {
        return nil, fmt.Errorf(typeErrorTemplate, addingExpr.GetText(), "int or float")
    }
    if ctx.PLUS() != nil {
        return a + b, nil
    } else {
        return b - a, nil
    }
}

func (l *listener) evalMultiplyingExpr(ctx *grammar.MultiplyingExprContext) (result any, err error) {
    var (
        factorExpr      *grammar.FactorContext
        multiplyingExpr *grammar.MultiplyingExprContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.MultiplyingExprContext); ok {
            multiplyingExpr = c
        } else if a, ok := child.(*grammar.FactorContext); ok {
            factorExpr = a
        }
    }
    factorResult, err := l.evalFactor(factorExpr)
    if err != nil {
        return nil, err
    }
    if multiplyingExpr == nil {
        return factorResult, nil
    }
    multiplyingResult, err := l.evalMultiplyingExpr(multiplyingExpr)
    if err != nil {
        return nil, err
    }
    a, err := toFloat(factorResult)
    if err != nil {
        return nil, fmt.Errorf(typeErrorTemplate, factorExpr.GetText(), "int or float")
    }
    b, err := toFloat(multiplyingResult)
    if err != nil {
        return nil, fmt.Errorf(typeErrorTemplate, multiplyingExpr.GetText(), "int or float")
    }
    if ctx.ASTERISK() != nil {
        return a * b, nil
    } else if ctx.SLASH() != nil {
        return b / a, nil
    } else {
        return int(b) % int(a), nil
    }
}

func (l *listener) evalFactor(ctx *grammar.FactorContext) (result any, err error) {
    var (
        valueExpr *grammar.ValueContext
        funcExpr  *grammar.FuncContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.ValueContext); ok {
            valueExpr = c
        } else if a, ok := child.(*grammar.FuncContext); ok {
            funcExpr = a
        }
    }
    if valueExpr != nil {
        return l.evalValue(valueExpr)
    }
    return l.evalFunc(funcExpr)
}

func (l *listener) evalValue(ctx *grammar.ValueContext) (result any, err error) {
    var (
        exprSelector *grammar.ExprSelectorContext
        constantExpr *grammar.ConstantContext
    )
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.ConstantContext); ok {
            constantExpr = c
        } else if a, ok := child.(*grammar.ExprSelectorContext); ok {
            exprSelector = a
        }
    }
    if exprSelector != nil {
        return l.evalExprSelector(exprSelector)
    }
    if constantExpr != nil {
        return l.evalConstant(constantExpr)
    }
    if ctx.NIL() != nil {
        return nil, nil
    }
    if ctx.Expr() != nil {
        return l.evalExpr(ctx.Expr().(*grammar.ExprContext))
    }
    return nil, nil
}

func (l *listener) evalFunc(ctx *grammar.FuncContext) (result any, err error) {
    funcName := ctx.FuncName()
    var exprs []*grammar.ExprContext
    funcArgs := ctx.FuncArgs()
    if funcArgs == nil {
        return nil, fmt.Errorf("expr %s's function len need args", funcName.GetText())
    }
    for _, child := range funcArgs.GetChildren() {
        if c, ok := child.(*grammar.ExprContext); ok {
            exprs = append(exprs, c)
        }
    }
    if funcName.GetText() == "len" {
        if len(exprs) != 1 {
            return nil, fmt.Errorf("expr %s's function len need only one arg", funcName.GetText())
        }
        funcArgResult, err0 := l.evalExpr(exprs[0])
        if err0 != nil {
            return nil, err0
        }
        v := reflect.ValueOf(funcArgResult)
        switch v.Kind() {
        case reflect.Slice, reflect.String, reflect.Array, reflect.Map:
            return v.Len(), nil
        default:
            return nil, fmt.Errorf("expr %s's result type is not [slice|string|array|map]", funcName.GetText())
        }
    } else if funcName.GetText() == "regexp" {
        if len(exprs) != 1 {
            return nil, fmt.Errorf("expr %s's function len need only one arg", funcName.GetText())
        }
        funcArgResult, err0 := l.evalExpr(exprs[0])
        if err0 != nil {
            return nil, err0
        }
        v := reflect.ValueOf(funcArgResult)
        switch v.Kind() {
        case reflect.String:
            selfValue := getValue(l.s, l.self)
            if v, ok := selfValue.(string); ok {
                return regexp.Match(funcArgResult.(string), []byte(v))
            } else {
                return nil, fmt.Errorf("field %s's value is not string under regexp", l.self)
            }
        default:
            return nil, fmt.Errorf("expr %s's result type is not [string]", funcName.GetText())
        }
    } else {
        function := registerFunctions[funcName.GetText()]
        if function == nil {
            return nil, fmt.Errorf("function %s does not exist", funcName.GetText())
        }
        var values []any
        for _, expr := range exprs {
            exprResult, err0 := l.evalExpr(expr)
            if err0 != nil {
                return nil, err0
            }
            values = append(values, exprResult)
        }
        err = function(values...)
        if err != nil {
            return nil, err
        } else {
            return true, nil
        }
    }
}

func (l *listener) evalExprSelector(ctx *grammar.ExprSelectorContext) (result any, err error) {
    head := ctx.SelectorHead()
    var last any
    if head.SELF() != nil {
        last = getValue(l.s, l.self)
    } else {
        last = getValue(l.s, head.FieldExpr().VAR().GetText())
    }
    for _, child := range ctx.GetChildren() {
        if c, ok := child.(*grammar.SelectorBodyContext); ok {
            if c.BracketSelector() != nil {
                exprResult, err0 := l.evalExpr(c.BracketSelector().Expr().(*grammar.ExprContext))
                if err0 != nil {
                    return nil, err0
                }
                switch v := exprResult.(type) {
                case string:
                    v2 := reflect.ValueOf(last)
                    if v2.Kind() == reflect.Map {
                        var k string
                        for _, key := range v2.MapKeys() {
                            x := key.String()
                            _ = x
                            if key.String() == v {
                                k = key.String()
                            }
                        }
                        if k == "" {
                            return nil, fmt.Errorf("expr %s's key %s does not exist", c.BracketSelector().Expr().GetText(), v)
                        }
                        last = v2.MapIndex(reflect.ValueOf(k)).Interface()
                    } else {
                        return nil, fmt.Errorf("expr %s's type must be map", c.BracketSelector().Expr().GetText())
                    }
                case int, int64:
                    index := int(reflect.ValueOf(v).Int())
                    v2 := reflect.ValueOf(last)
                    switch v2.Kind() {
                    case reflect.Map:
                        keyType := v2.Type().Key()
                        key := reflect.ValueOf(convertIntToType(int64(index), keyType.Kind()))
                        mapValue := v2.MapIndex(key)
                        if !mapValue.IsValid() {
                            return nil, fmt.Errorf("expr %s's key %d does not exist", c.BracketSelector().Expr().GetText(), v)
                        }
                        last = mapValue.Interface()
                    case reflect.Slice:
                        if v2.Len() <= index || index < 0 {
                            return nil, fmt.Errorf("expr %s's index %d does not exist", c.BracketSelector().Expr().GetText(), index)
                        }
                        last = v2.Index(index).Interface()
                    default:
                        return nil, fmt.Errorf("expr %s's type must be [map|slice]", c.BracketSelector().Expr().GetText())
                    }
                case float64:
                    v2 := reflect.ValueOf(last)
                    switch v2.Kind() {
                    case reflect.Map:
                        mapValue := v2.MapIndex(reflect.ValueOf(v))
                        if !mapValue.IsValid() {
                            return nil, fmt.Errorf("expr %s's key %f does not exist", c.BracketSelector().Expr().GetText(), v)
                        }
                        last = mapValue.Interface()
                    default:
                        return nil, fmt.Errorf("expr %s's type must be [map]", c.BracketSelector().Expr().GetText())
                    }
                default:
                    return nil, fmt.Errorf("expr %s's type must be [int|map|slice]", c.BracketSelector().Expr().GetText())
                }
            } else if c.FieldExpr() != nil {
                fieldName := c.FieldExpr().VAR().GetText()
                last = getValue(last, fieldName)
            }
        }
    }
    return last, nil
}

func (l *listener) evalConstant(ctx *grammar.ConstantContext) (result any, err error) {
    if ctx.STRING() != nil {
        return strings.Trim(ctx.STRING().GetText(), "'"), nil
    }
    if ctx.INTEGER_VALUE() != nil {
        return strconv.ParseInt(ctx.INTEGER_VALUE().GetText(), 0, 64)
    }
    if ctx.DOUBLE_VALUE() != nil {
        return strconv.ParseFloat(ctx.DOUBLE_VALUE().GetText(), 64)
    }
    return nil, nil
}
