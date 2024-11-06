package parser

import "github.com/antlr4-go/antlr/v4"

type errorListener struct {
    antlr.ErrorListener
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer,
    offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
    panic(msg)
}
