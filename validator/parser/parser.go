package parser

import (
    "errors"
    "github.com/antlr4-go/antlr/v4"
    "github.com/joyant/fire/v2/validator/grammar"
    "strings"
    "sync"
)

var (
    // for testing
    getParserFunc = getParser
    errorHandle   = &errorListener{}
    walker        = antlr.ParseTreeWalkerDefault

    lexerPool  sync.Pool
    parserPool sync.Pool
)

// Parse parses
func Parse(b string, s any, self string) (err error) {
    defer func() {
        if r := recover(); r != nil {
           switch x := r.(type) {
           case string:
               err = errors.New(x)
           case error:
               err = x
           default:
               err = errors.New("unknown panic when parse")
           }
        }
    }()

    b = strings.ReplaceAll(b, `\"`, `"`)
    input := antlr.NewInputStream(b)

    lexer := getLexer(input)
    defer putLexer(lexer)

    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := getParserFunc(tokens)
    ctx := parser.Validator()

    // create listener
    listener := newListener(s, self)

    walker.Walk(listener, ctx)

    err = listener.Validate()
    if err == nil {
        putParser(parser)
    }

    return
}

func getLexer(input *antlr.InputStream) *grammar.VALIDATORLexer {
    lexer := lexerPool.Get()
    if lexer == nil {
        lexer := grammar.NewVALIDATORLexer(input)
        lexer.RemoveErrorListeners()
        lexer.AddErrorListener(errorHandle)
        return lexer
    }
    l := lexer.(*grammar.VALIDATORLexer)
    l.SetInputStream(input)
    return l
}

func putLexer(l *grammar.VALIDATORLexer) {
    lexerPool.Put(l)
}

// getParser picks a cached parser from the pool
func getParser(tokenStream *antlr.CommonTokenStream) *grammar.VALIDATORParser {
    parser := parserPool.Get()
    if parser == nil {
        parser := grammar.NewVALIDATORParser(tokenStream)
        parser.BuildParseTrees = true
        parser.RemoveErrorListeners()
        parser.AddErrorListener(errorHandle)
        return parser
    }
    p := parser.(*grammar.VALIDATORParser)
    p.SetInputStream(tokenStream)
    return p
}

func putParser(p *grammar.VALIDATORParser) {
    parserPool.Put(p)
}
