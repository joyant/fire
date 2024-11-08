// Code generated from ./validator/grammar/VALIDATOR.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type VALIDATORLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var VALIDATORLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func validatorlexerLexerInit() {
	staticData := &VALIDATORLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'.'", "'msg:'", "'sprintf'", "'>'", "'>='", "'<'", "'<='", "'=='",
		"'!='", "'+'", "'-'", "'*'", "'/'", "'%'", "':'", "';'", "','", "'$'",
		"'this'", "'nil'", "", "'&&'", "'||'", "'('", "')'", "'['", "']'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "GT", "EGT", "LT", "ELT", "EQ", "NOTEQ", "PLUS", "MINUS",
		"ASTERISK", "SLASH", "MODULO", "COLON", "SEMICOLON", "COMMA", "SELF",
		"THIS", "NIL", "VAR", "AND", "OR", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS",
		"LEFT_BRACKET", "RIGHT_BRACKET", "QUESTION_MARK", "STRING", "INTEGER_VALUE",
		"DOUBLE_VALUE", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "GT", "EGT", "LT", "ELT", "EQ", "NOTEQ", "PLUS",
		"MINUS", "ASTERISK", "SLASH", "MODULO", "COLON", "SEMICOLON", "COMMA",
		"SELF", "THIS", "NIL", "VAR", "AND", "OR", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS",
		"LEFT_BRACKET", "RIGHT_BRACKET", "QUESTION_MARK", "STRING", "INTEGER_VALUE",
		"DOUBLE_VALUE", "DECIMAL_INTEGER", "EXPONENT", "DIGIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 226, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 4, 20, 132, 8, 20, 11, 20, 12, 20, 133, 1, 20, 1, 20, 5, 20, 138,
		8, 20, 10, 20, 12, 20, 141, 9, 20, 3, 20, 143, 8, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 165, 8, 28, 10,
		28, 12, 28, 168, 9, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 4, 30, 175,
		8, 30, 11, 30, 12, 30, 176, 1, 30, 1, 30, 5, 30, 181, 8, 30, 10, 30, 12,
		30, 184, 9, 30, 3, 30, 186, 8, 30, 1, 30, 3, 30, 189, 8, 30, 1, 30, 1,
		30, 4, 30, 193, 8, 30, 11, 30, 12, 30, 194, 1, 30, 3, 30, 198, 8, 30, 3,
		30, 200, 8, 30, 1, 31, 1, 31, 3, 31, 204, 8, 31, 1, 31, 5, 31, 207, 8,
		31, 10, 31, 12, 31, 210, 9, 31, 1, 32, 1, 32, 3, 32, 214, 8, 32, 1, 32,
		4, 32, 217, 8, 32, 11, 32, 12, 32, 218, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 1, 34, 0, 0, 35, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 0, 65, 0, 67, 0, 69, 32, 1,
		0, 7, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 1, 0, 39, 39, 2, 0, 43, 43, 45, 45, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 238, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 1, 71, 1, 0,
		0, 0, 3, 73, 1, 0, 0, 0, 5, 78, 1, 0, 0, 0, 7, 86, 1, 0, 0, 0, 9, 88, 1,
		0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 93, 1, 0, 0, 0, 15, 96, 1, 0, 0, 0, 17,
		99, 1, 0, 0, 0, 19, 102, 1, 0, 0, 0, 21, 104, 1, 0, 0, 0, 23, 106, 1, 0,
		0, 0, 25, 108, 1, 0, 0, 0, 27, 110, 1, 0, 0, 0, 29, 112, 1, 0, 0, 0, 31,
		114, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 118, 1, 0, 0, 0, 37, 120, 1,
		0, 0, 0, 39, 125, 1, 0, 0, 0, 41, 142, 1, 0, 0, 0, 43, 144, 1, 0, 0, 0,
		45, 147, 1, 0, 0, 0, 47, 150, 1, 0, 0, 0, 49, 152, 1, 0, 0, 0, 51, 154,
		1, 0, 0, 0, 53, 156, 1, 0, 0, 0, 55, 158, 1, 0, 0, 0, 57, 160, 1, 0, 0,
		0, 59, 171, 1, 0, 0, 0, 61, 199, 1, 0, 0, 0, 63, 201, 1, 0, 0, 0, 65, 211,
		1, 0, 0, 0, 67, 220, 1, 0, 0, 0, 69, 222, 1, 0, 0, 0, 71, 72, 5, 46, 0,
		0, 72, 2, 1, 0, 0, 0, 73, 74, 5, 109, 0, 0, 74, 75, 5, 115, 0, 0, 75, 76,
		5, 103, 0, 0, 76, 77, 5, 58, 0, 0, 77, 4, 1, 0, 0, 0, 78, 79, 5, 115, 0,
		0, 79, 80, 5, 112, 0, 0, 80, 81, 5, 114, 0, 0, 81, 82, 5, 105, 0, 0, 82,
		83, 5, 110, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 102, 0, 0, 85, 6, 1,
		0, 0, 0, 86, 87, 5, 62, 0, 0, 87, 8, 1, 0, 0, 0, 88, 89, 5, 62, 0, 0, 89,
		90, 5, 61, 0, 0, 90, 10, 1, 0, 0, 0, 91, 92, 5, 60, 0, 0, 92, 12, 1, 0,
		0, 0, 93, 94, 5, 60, 0, 0, 94, 95, 5, 61, 0, 0, 95, 14, 1, 0, 0, 0, 96,
		97, 5, 61, 0, 0, 97, 98, 5, 61, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100, 5, 33,
		0, 0, 100, 101, 5, 61, 0, 0, 101, 18, 1, 0, 0, 0, 102, 103, 5, 43, 0, 0,
		103, 20, 1, 0, 0, 0, 104, 105, 5, 45, 0, 0, 105, 22, 1, 0, 0, 0, 106, 107,
		5, 42, 0, 0, 107, 24, 1, 0, 0, 0, 108, 109, 5, 47, 0, 0, 109, 26, 1, 0,
		0, 0, 110, 111, 5, 37, 0, 0, 111, 28, 1, 0, 0, 0, 112, 113, 5, 58, 0, 0,
		113, 30, 1, 0, 0, 0, 114, 115, 5, 59, 0, 0, 115, 32, 1, 0, 0, 0, 116, 117,
		5, 44, 0, 0, 117, 34, 1, 0, 0, 0, 118, 119, 5, 36, 0, 0, 119, 36, 1, 0,
		0, 0, 120, 121, 5, 116, 0, 0, 121, 122, 5, 104, 0, 0, 122, 123, 5, 105,
		0, 0, 123, 124, 5, 115, 0, 0, 124, 38, 1, 0, 0, 0, 125, 126, 5, 110, 0,
		0, 126, 127, 5, 105, 0, 0, 127, 128, 5, 108, 0, 0, 128, 40, 1, 0, 0, 0,
		129, 131, 5, 95, 0, 0, 130, 132, 7, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 143,
		1, 0, 0, 0, 135, 139, 7, 1, 0, 0, 136, 138, 7, 2, 0, 0, 137, 136, 1, 0,
		0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 129, 1, 0, 0, 0, 142,
		135, 1, 0, 0, 0, 143, 42, 1, 0, 0, 0, 144, 145, 5, 38, 0, 0, 145, 146,
		5, 38, 0, 0, 146, 44, 1, 0, 0, 0, 147, 148, 5, 124, 0, 0, 148, 149, 5,
		124, 0, 0, 149, 46, 1, 0, 0, 0, 150, 151, 5, 40, 0, 0, 151, 48, 1, 0, 0,
		0, 152, 153, 5, 41, 0, 0, 153, 50, 1, 0, 0, 0, 154, 155, 5, 91, 0, 0, 155,
		52, 1, 0, 0, 0, 156, 157, 5, 93, 0, 0, 157, 54, 1, 0, 0, 0, 158, 159, 5,
		63, 0, 0, 159, 56, 1, 0, 0, 0, 160, 166, 5, 39, 0, 0, 161, 165, 8, 3, 0,
		0, 162, 163, 5, 39, 0, 0, 163, 165, 5, 39, 0, 0, 164, 161, 1, 0, 0, 0,
		164, 162, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170,
		5, 39, 0, 0, 170, 58, 1, 0, 0, 0, 171, 172, 3, 63, 31, 0, 172, 60, 1, 0,
		0, 0, 173, 175, 3, 67, 33, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 185, 1, 0, 0, 0, 178,
		182, 5, 46, 0, 0, 179, 181, 3, 67, 33, 0, 180, 179, 1, 0, 0, 0, 181, 184,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 186, 1, 0,
		0, 0, 184, 182, 1, 0, 0, 0, 185, 178, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 188, 1, 0, 0, 0, 187, 189, 3, 65, 32, 0, 188, 187, 1, 0, 0, 0, 188,
		189, 1, 0, 0, 0, 189, 200, 1, 0, 0, 0, 190, 192, 5, 46, 0, 0, 191, 193,
		3, 67, 33, 0, 192, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 192, 1,
		0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 198, 3, 65, 32,
		0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199,
		174, 1, 0, 0, 0, 199, 190, 1, 0, 0, 0, 200, 62, 1, 0, 0, 0, 201, 208, 3,
		67, 33, 0, 202, 204, 5, 95, 0, 0, 203, 202, 1, 0, 0, 0, 203, 204, 1, 0,
		0, 0, 204, 205, 1, 0, 0, 0, 205, 207, 3, 67, 33, 0, 206, 203, 1, 0, 0,
		0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		64, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 213, 5, 69, 0, 0, 212, 214,
		7, 4, 0, 0, 213, 212, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0,
		0, 0, 215, 217, 3, 67, 33, 0, 216, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0,
		0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 66, 1, 0, 0, 0, 220,
		221, 7, 5, 0, 0, 221, 68, 1, 0, 0, 0, 222, 223, 7, 6, 0, 0, 223, 224, 1,
		0, 0, 0, 224, 225, 6, 34, 0, 0, 225, 70, 1, 0, 0, 0, 17, 0, 133, 139, 142,
		164, 166, 176, 182, 185, 188, 194, 197, 199, 203, 208, 213, 218, 1, 6,
		0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VALIDATORLexerInit initializes any static state used to implement VALIDATORLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewVALIDATORLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VALIDATORLexerInit() {
	staticData := &VALIDATORLexerLexerStaticData
	staticData.once.Do(validatorlexerLexerInit)
}

// NewVALIDATORLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewVALIDATORLexer(input antlr.CharStream) *VALIDATORLexer {
	VALIDATORLexerInit()
	l := new(VALIDATORLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &VALIDATORLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "VALIDATOR.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VALIDATORLexer tokens.
const (
	VALIDATORLexerT__0              = 1
	VALIDATORLexerT__1              = 2
	VALIDATORLexerT__2              = 3
	VALIDATORLexerGT                = 4
	VALIDATORLexerEGT               = 5
	VALIDATORLexerLT                = 6
	VALIDATORLexerELT               = 7
	VALIDATORLexerEQ                = 8
	VALIDATORLexerNOTEQ             = 9
	VALIDATORLexerPLUS              = 10
	VALIDATORLexerMINUS             = 11
	VALIDATORLexerASTERISK          = 12
	VALIDATORLexerSLASH             = 13
	VALIDATORLexerMODULO            = 14
	VALIDATORLexerCOLON             = 15
	VALIDATORLexerSEMICOLON         = 16
	VALIDATORLexerCOMMA             = 17
	VALIDATORLexerSELF              = 18
	VALIDATORLexerTHIS              = 19
	VALIDATORLexerNIL               = 20
	VALIDATORLexerVAR               = 21
	VALIDATORLexerAND               = 22
	VALIDATORLexerOR                = 23
	VALIDATORLexerLEFT_PARENTHESIS  = 24
	VALIDATORLexerRIGHT_PARENTHESIS = 25
	VALIDATORLexerLEFT_BRACKET      = 26
	VALIDATORLexerRIGHT_BRACKET     = 27
	VALIDATORLexerQUESTION_MARK     = 28
	VALIDATORLexerSTRING            = 29
	VALIDATORLexerINTEGER_VALUE     = 30
	VALIDATORLexerDOUBLE_VALUE      = 31
	VALIDATORLexerWS                = 32
)
