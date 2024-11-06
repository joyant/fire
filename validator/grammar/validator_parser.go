// Code generated from ./validator/grammar/VALIDATOR.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // VALIDATOR
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VALIDATORParser struct {
	*antlr.BaseParser
}

var VALIDATORParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func validatorParserInit() {
	staticData := &VALIDATORParserStaticData
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
		"validator", "expr", "andExpr", "conditionalExpr", "comparingExpr",
		"addingExpr", "multiplyingExpr", "factor", "value", "constant", "func",
		"funcArgs", "funcName", "selectorHead", "selectorBody", "bracketSelector",
		"exprSelector", "fieldExpr", "msg", "sprintf", "sprintfArgs", "compareOperator",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 182, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 1, 0, 1, 0, 3, 0, 48, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 3, 1, 55, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 74, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 85, 8, 5,
		10, 5, 12, 5, 88, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 102, 8, 6, 10, 6, 12, 6, 105, 9, 6, 1, 7,
		1, 7, 3, 7, 109, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		118, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 125, 8, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 5, 11, 132, 8, 11, 10, 11, 12, 11, 135, 9, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 143, 8, 13, 1, 14, 1,
		14, 1, 14, 3, 14, 148, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16,
		5, 16, 156, 8, 16, 10, 16, 12, 16, 159, 9, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 170, 8, 19, 10, 19, 12, 19,
		173, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 0,
		2, 10, 12, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 0, 2, 1, 0, 29, 31, 1, 0, 4, 9, 179, 0, 44, 1,
		0, 0, 0, 2, 51, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 69,
		1, 0, 0, 0, 10, 75, 1, 0, 0, 0, 12, 89, 1, 0, 0, 0, 14, 108, 1, 0, 0, 0,
		16, 117, 1, 0, 0, 0, 18, 119, 1, 0, 0, 0, 20, 121, 1, 0, 0, 0, 22, 128,
		1, 0, 0, 0, 24, 136, 1, 0, 0, 0, 26, 142, 1, 0, 0, 0, 28, 147, 1, 0, 0,
		0, 30, 149, 1, 0, 0, 0, 32, 153, 1, 0, 0, 0, 34, 160, 1, 0, 0, 0, 36, 162,
		1, 0, 0, 0, 38, 165, 1, 0, 0, 0, 40, 176, 1, 0, 0, 0, 42, 179, 1, 0, 0,
		0, 44, 47, 3, 2, 1, 0, 45, 46, 5, 16, 0, 0, 46, 48, 3, 36, 18, 0, 47, 45,
		1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 5, 0, 0, 1,
		50, 1, 1, 0, 0, 0, 51, 54, 3, 4, 2, 0, 52, 53, 5, 23, 0, 0, 53, 55, 3,
		2, 1, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 3, 1, 0, 0, 0, 56,
		59, 3, 6, 3, 0, 57, 58, 5, 22, 0, 0, 58, 60, 3, 4, 2, 0, 59, 57, 1, 0,
		0, 0, 59, 60, 1, 0, 0, 0, 60, 5, 1, 0, 0, 0, 61, 67, 3, 8, 4, 0, 62, 63,
		5, 28, 0, 0, 63, 64, 3, 2, 1, 0, 64, 65, 5, 15, 0, 0, 65, 66, 3, 6, 3,
		0, 66, 68, 1, 0, 0, 0, 67, 62, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 7, 1,
		0, 0, 0, 69, 73, 3, 10, 5, 0, 70, 71, 3, 42, 21, 0, 71, 72, 3, 8, 4, 0,
		72, 74, 1, 0, 0, 0, 73, 70, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 9, 1, 0,
		0, 0, 75, 76, 6, 5, -1, 0, 76, 77, 3, 12, 6, 0, 77, 86, 1, 0, 0, 0, 78,
		79, 10, 2, 0, 0, 79, 80, 5, 10, 0, 0, 80, 85, 3, 12, 6, 0, 81, 82, 10,
		1, 0, 0, 82, 83, 5, 11, 0, 0, 83, 85, 3, 12, 6, 0, 84, 78, 1, 0, 0, 0,
		84, 81, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1,
		0, 0, 0, 87, 11, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 6, 6, -1, 0, 90,
		91, 3, 14, 7, 0, 91, 103, 1, 0, 0, 0, 92, 93, 10, 3, 0, 0, 93, 94, 5, 12,
		0, 0, 94, 102, 3, 14, 7, 0, 95, 96, 10, 2, 0, 0, 96, 97, 5, 13, 0, 0, 97,
		102, 3, 14, 7, 0, 98, 99, 10, 1, 0, 0, 99, 100, 5, 14, 0, 0, 100, 102,
		3, 14, 7, 0, 101, 92, 1, 0, 0, 0, 101, 95, 1, 0, 0, 0, 101, 98, 1, 0, 0,
		0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		13, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 109, 3, 16, 8, 0, 107, 109,
		3, 20, 10, 0, 108, 106, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 15, 1, 0,
		0, 0, 110, 118, 3, 32, 16, 0, 111, 118, 3, 18, 9, 0, 112, 118, 5, 20, 0,
		0, 113, 114, 5, 24, 0, 0, 114, 115, 3, 2, 1, 0, 115, 116, 5, 25, 0, 0,
		116, 118, 1, 0, 0, 0, 117, 110, 1, 0, 0, 0, 117, 111, 1, 0, 0, 0, 117,
		112, 1, 0, 0, 0, 117, 113, 1, 0, 0, 0, 118, 17, 1, 0, 0, 0, 119, 120, 7,
		0, 0, 0, 120, 19, 1, 0, 0, 0, 121, 122, 3, 24, 12, 0, 122, 124, 5, 24,
		0, 0, 123, 125, 3, 22, 11, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0,
		0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 25, 0, 0, 127, 21, 1, 0, 0, 0, 128,
		133, 3, 2, 1, 0, 129, 130, 5, 17, 0, 0, 130, 132, 3, 2, 1, 0, 131, 129,
		1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 23, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 21, 0, 0,
		137, 25, 1, 0, 0, 0, 138, 139, 5, 19, 0, 0, 139, 140, 5, 1, 0, 0, 140,
		143, 3, 34, 17, 0, 141, 143, 5, 18, 0, 0, 142, 138, 1, 0, 0, 0, 142, 141,
		1, 0, 0, 0, 143, 27, 1, 0, 0, 0, 144, 148, 3, 30, 15, 0, 145, 146, 5, 1,
		0, 0, 146, 148, 3, 34, 17, 0, 147, 144, 1, 0, 0, 0, 147, 145, 1, 0, 0,
		0, 148, 29, 1, 0, 0, 0, 149, 150, 5, 26, 0, 0, 150, 151, 3, 2, 1, 0, 151,
		152, 5, 27, 0, 0, 152, 31, 1, 0, 0, 0, 153, 157, 3, 26, 13, 0, 154, 156,
		3, 28, 14, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1,
		0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 33, 1, 0, 0, 0, 159, 157, 1, 0, 0,
		0, 160, 161, 5, 21, 0, 0, 161, 35, 1, 0, 0, 0, 162, 163, 5, 2, 0, 0, 163,
		164, 3, 38, 19, 0, 164, 37, 1, 0, 0, 0, 165, 166, 5, 3, 0, 0, 166, 167,
		5, 24, 0, 0, 167, 171, 5, 29, 0, 0, 168, 170, 3, 40, 20, 0, 169, 168, 1,
		0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0,
		0, 172, 174, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 175, 5, 25, 0, 0, 175,
		39, 1, 0, 0, 0, 176, 177, 5, 17, 0, 0, 177, 178, 3, 2, 1, 0, 178, 41, 1,
		0, 0, 0, 179, 180, 7, 1, 0, 0, 180, 43, 1, 0, 0, 0, 17, 47, 54, 59, 67,
		73, 84, 86, 101, 103, 108, 117, 124, 133, 142, 147, 157, 171,
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

// VALIDATORParserInit initializes any static state used to implement VALIDATORParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVALIDATORParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VALIDATORParserInit() {
	staticData := &VALIDATORParserStaticData
	staticData.once.Do(validatorParserInit)
}

// NewVALIDATORParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVALIDATORParser(input antlr.TokenStream) *VALIDATORParser {
	VALIDATORParserInit()
	this := new(VALIDATORParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VALIDATORParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VALIDATOR.g4"

	return this
}

// VALIDATORParser tokens.
const (
	VALIDATORParserEOF               = antlr.TokenEOF
	VALIDATORParserT__0              = 1
	VALIDATORParserT__1              = 2
	VALIDATORParserT__2              = 3
	VALIDATORParserGT                = 4
	VALIDATORParserEGT               = 5
	VALIDATORParserLT                = 6
	VALIDATORParserELT               = 7
	VALIDATORParserEQ                = 8
	VALIDATORParserNOTEQ             = 9
	VALIDATORParserPLUS              = 10
	VALIDATORParserMINUS             = 11
	VALIDATORParserASTERISK          = 12
	VALIDATORParserSLASH             = 13
	VALIDATORParserMODULO            = 14
	VALIDATORParserCOLON             = 15
	VALIDATORParserSEMICOLON         = 16
	VALIDATORParserCOMMA             = 17
	VALIDATORParserSELF              = 18
	VALIDATORParserTHIS              = 19
	VALIDATORParserNIL               = 20
	VALIDATORParserVAR               = 21
	VALIDATORParserAND               = 22
	VALIDATORParserOR                = 23
	VALIDATORParserLEFT_PARENTHESIS  = 24
	VALIDATORParserRIGHT_PARENTHESIS = 25
	VALIDATORParserLEFT_BRACKET      = 26
	VALIDATORParserRIGHT_BRACKET     = 27
	VALIDATORParserQUESTION_MARK     = 28
	VALIDATORParserSTRING            = 29
	VALIDATORParserINTEGER_VALUE     = 30
	VALIDATORParserDOUBLE_VALUE      = 31
	VALIDATORParserWS                = 32
)

// VALIDATORParser rules.
const (
	VALIDATORParserRULE_validator       = 0
	VALIDATORParserRULE_expr            = 1
	VALIDATORParserRULE_andExpr         = 2
	VALIDATORParserRULE_conditionalExpr = 3
	VALIDATORParserRULE_comparingExpr   = 4
	VALIDATORParserRULE_addingExpr      = 5
	VALIDATORParserRULE_multiplyingExpr = 6
	VALIDATORParserRULE_factor          = 7
	VALIDATORParserRULE_value           = 8
	VALIDATORParserRULE_constant        = 9
	VALIDATORParserRULE_func            = 10
	VALIDATORParserRULE_funcArgs        = 11
	VALIDATORParserRULE_funcName        = 12
	VALIDATORParserRULE_selectorHead    = 13
	VALIDATORParserRULE_selectorBody    = 14
	VALIDATORParserRULE_bracketSelector = 15
	VALIDATORParserRULE_exprSelector    = 16
	VALIDATORParserRULE_fieldExpr       = 17
	VALIDATORParserRULE_msg             = 18
	VALIDATORParserRULE_sprintf         = 19
	VALIDATORParserRULE_sprintfArgs     = 20
	VALIDATORParserRULE_compareOperator = 21
)

// IValidatorContext is an interface to support dynamic dispatch.
type IValidatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Msg() IMsgContext

	// IsValidatorContext differentiates from other interfaces.
	IsValidatorContext()
}

type ValidatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidatorContext() *ValidatorContext {
	var p = new(ValidatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_validator
	return p
}

func InitEmptyValidatorContext(p *ValidatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_validator
}

func (*ValidatorContext) IsValidatorContext() {}

func NewValidatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidatorContext {
	var p = new(ValidatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_validator

	return p
}

func (s *ValidatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidatorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValidatorContext) EOF() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserEOF, 0)
}

func (s *ValidatorContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserSEMICOLON, 0)
}

func (s *ValidatorContext) Msg() IMsgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMsgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMsgContext)
}

func (s *ValidatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterValidator(s)
	}
}

func (s *ValidatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitValidator(s)
	}
}

func (s *ValidatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitValidator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Validator() (localctx IValidatorContext) {
	localctx = NewValidatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VALIDATORParserRULE_validator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Expr()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VALIDATORParserSEMICOLON {
		{
			p.SetState(45)
			p.Match(VALIDATORParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Msg()
		}

	}
	{
		p.SetState(49)
		p.Match(VALIDATORParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AndExpr() IAndExprContext
	OR() antlr.TerminalNode
	Expr() IExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserOR, 0)
}

func (s *ExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VALIDATORParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.AndExpr()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VALIDATORParserOR {
		{
			p.SetState(52)
			p.Match(VALIDATORParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Expr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext
	AND() antlr.TerminalNode
	AndExpr() IAndExprContext

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_andExpr
	return p
}

func InitEmptyAndExprContext(p *AndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_andExpr
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserAND, 0)
}

func (s *AndExprContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) AndExpr() (localctx IAndExprContext) {
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VALIDATORParserRULE_andExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.ConditionalExpr()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VALIDATORParserAND {
		{
			p.SetState(57)
			p.Match(VALIDATORParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.AndExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalExprContext is an interface to support dynamic dispatch.
type IConditionalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComparingExpr() IComparingExprContext
	QUESTION_MARK() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	ConditionalExpr() IConditionalExprContext

	// IsConditionalExprContext differentiates from other interfaces.
	IsConditionalExprContext()
}

type ConditionalExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalExprContext() *ConditionalExprContext {
	var p = new(ConditionalExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_conditionalExpr
	return p
}

func InitEmptyConditionalExprContext(p *ConditionalExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_conditionalExpr
}

func (*ConditionalExprContext) IsConditionalExprContext() {}

func NewConditionalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExprContext {
	var p = new(ConditionalExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_conditionalExpr

	return p
}

func (s *ConditionalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExprContext) ComparingExpr() IComparingExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparingExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparingExprContext)
}

func (s *ConditionalExprContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserQUESTION_MARK, 0)
}

func (s *ConditionalExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionalExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserCOLON, 0)
}

func (s *ConditionalExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitConditionalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) ConditionalExpr() (localctx IConditionalExprContext) {
	localctx = NewConditionalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VALIDATORParserRULE_conditionalExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.ComparingExpr()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VALIDATORParserQUESTION_MARK {
		{
			p.SetState(62)
			p.Match(VALIDATORParserQUESTION_MARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Expr()
		}
		{
			p.SetState(64)
			p.Match(VALIDATORParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.ConditionalExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparingExprContext is an interface to support dynamic dispatch.
type IComparingExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddingExpr() IAddingExprContext
	CompareOperator() ICompareOperatorContext
	ComparingExpr() IComparingExprContext

	// IsComparingExprContext differentiates from other interfaces.
	IsComparingExprContext()
}

type ComparingExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparingExprContext() *ComparingExprContext {
	var p = new(ComparingExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_comparingExpr
	return p
}

func InitEmptyComparingExprContext(p *ComparingExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_comparingExpr
}

func (*ComparingExprContext) IsComparingExprContext() {}

func NewComparingExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparingExprContext {
	var p = new(ComparingExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_comparingExpr

	return p
}

func (s *ComparingExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparingExprContext) AddingExpr() IAddingExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddingExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddingExprContext)
}

func (s *ComparingExprContext) CompareOperator() ICompareOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareOperatorContext)
}

func (s *ComparingExprContext) ComparingExpr() IComparingExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparingExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparingExprContext)
}

func (s *ComparingExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparingExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparingExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterComparingExpr(s)
	}
}

func (s *ComparingExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitComparingExpr(s)
	}
}

func (s *ComparingExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitComparingExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) ComparingExpr() (localctx IComparingExprContext) {
	localctx = NewComparingExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VALIDATORParserRULE_comparingExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.addingExpr(0)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0 {
		{
			p.SetState(70)
			p.CompareOperator()
		}
		{
			p.SetState(71)
			p.ComparingExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddingExprContext is an interface to support dynamic dispatch.
type IAddingExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MultiplyingExpr() IMultiplyingExprContext
	AddingExpr() IAddingExprContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsAddingExprContext differentiates from other interfaces.
	IsAddingExprContext()
}

type AddingExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddingExprContext() *AddingExprContext {
	var p = new(AddingExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_addingExpr
	return p
}

func InitEmptyAddingExprContext(p *AddingExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_addingExpr
}

func (*AddingExprContext) IsAddingExprContext() {}

func NewAddingExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddingExprContext {
	var p = new(AddingExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_addingExpr

	return p
}

func (s *AddingExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AddingExprContext) MultiplyingExpr() IMultiplyingExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplyingExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExprContext)
}

func (s *AddingExprContext) AddingExpr() IAddingExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddingExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddingExprContext)
}

func (s *AddingExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserPLUS, 0)
}

func (s *AddingExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserMINUS, 0)
}

func (s *AddingExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddingExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddingExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterAddingExpr(s)
	}
}

func (s *AddingExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitAddingExpr(s)
	}
}

func (s *AddingExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitAddingExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) AddingExpr() (localctx IAddingExprContext) {
	return p.addingExpr(0)
}

func (p *VALIDATORParser) addingExpr(_p int) (localctx IAddingExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAddingExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAddingExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, VALIDATORParserRULE_addingExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.multiplyingExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAddingExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VALIDATORParserRULE_addingExpr)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(VALIDATORParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.multiplyingExpr(0)
				}

			case 2:
				localctx = NewAddingExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VALIDATORParserRULE_addingExpr)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					p.Match(VALIDATORParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.multiplyingExpr(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplyingExprContext is an interface to support dynamic dispatch.
type IMultiplyingExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	MultiplyingExpr() IMultiplyingExprContext
	ASTERISK() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	MODULO() antlr.TerminalNode

	// IsMultiplyingExprContext differentiates from other interfaces.
	IsMultiplyingExprContext()
}

type MultiplyingExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExprContext() *MultiplyingExprContext {
	var p = new(MultiplyingExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_multiplyingExpr
	return p
}

func InitEmptyMultiplyingExprContext(p *MultiplyingExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_multiplyingExpr
}

func (*MultiplyingExprContext) IsMultiplyingExprContext() {}

func NewMultiplyingExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExprContext {
	var p = new(MultiplyingExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_multiplyingExpr

	return p
}

func (s *MultiplyingExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExprContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *MultiplyingExprContext) MultiplyingExpr() IMultiplyingExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplyingExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExprContext)
}

func (s *MultiplyingExprContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserASTERISK, 0)
}

func (s *MultiplyingExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserSLASH, 0)
}

func (s *MultiplyingExprContext) MODULO() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserMODULO, 0)
}

func (s *MultiplyingExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterMultiplyingExpr(s)
	}
}

func (s *MultiplyingExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitMultiplyingExpr(s)
	}
}

func (s *MultiplyingExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitMultiplyingExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) MultiplyingExpr() (localctx IMultiplyingExprContext) {
	return p.multiplyingExpr(0)
}

func (p *VALIDATORParser) multiplyingExpr(_p int) (localctx IMultiplyingExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMultiplyingExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplyingExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, VALIDATORParserRULE_multiplyingExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Factor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyingExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VALIDATORParserRULE_multiplyingExpr)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(93)
					p.Match(VALIDATORParserASTERISK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)
					p.Factor()
				}

			case 2:
				localctx = NewMultiplyingExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VALIDATORParserRULE_multiplyingExpr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.Match(VALIDATORParserSLASH)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)
					p.Factor()
				}

			case 3:
				localctx = NewMultiplyingExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VALIDATORParserRULE_multiplyingExpr)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(99)
					p.Match(VALIDATORParserMODULO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(100)
					p.Factor()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	Func_() IFuncContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FactorContext) Func_() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VALIDATORParserRULE_factor)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VALIDATORParserSELF, VALIDATORParserTHIS, VALIDATORParserNIL, VALIDATORParserLEFT_PARENTHESIS, VALIDATORParserSTRING, VALIDATORParserINTEGER_VALUE, VALIDATORParserDOUBLE_VALUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Value()
		}

	case VALIDATORParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Func_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprSelector() IExprSelectorContext
	Constant() IConstantContext
	NIL() antlr.TerminalNode
	LEFT_PARENTHESIS() antlr.TerminalNode
	Expr() IExprContext
	RIGHT_PARENTHESIS() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) ExprSelector() IExprSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprSelectorContext)
}

func (s *ValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) NIL() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserNIL, 0)
}

func (s *ValueContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserLEFT_PARENTHESIS, 0)
}

func (s *ValueContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserRIGHT_PARENTHESIS, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VALIDATORParserRULE_value)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VALIDATORParserSELF, VALIDATORParserTHIS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.ExprSelector()
		}

	case VALIDATORParserSTRING, VALIDATORParserINTEGER_VALUE, VALIDATORParserDOUBLE_VALUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Constant()
		}

	case VALIDATORParserNIL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(VALIDATORParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VALIDATORParserLEFT_PARENTHESIS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Match(VALIDATORParserLEFT_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Expr()
		}
		{
			p.SetState(115)
			p.Match(VALIDATORParserRIGHT_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	INTEGER_VALUE() antlr.TerminalNode
	DOUBLE_VALUE() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) STRING() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserSTRING, 0)
}

func (s *ConstantContext) INTEGER_VALUE() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserINTEGER_VALUE, 0)
}

func (s *ConstantContext) DOUBLE_VALUE() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserDOUBLE_VALUE, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VALIDATORParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncName() IFuncNameContext
	LEFT_PARENTHESIS() antlr.TerminalNode
	RIGHT_PARENTHESIS() antlr.TerminalNode
	FuncArgs() IFuncArgsContext

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) FuncName() IFuncNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FuncContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserLEFT_PARENTHESIS, 0)
}

func (s *FuncContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserRIGHT_PARENTHESIS, 0)
}

func (s *FuncContext) FuncArgs() IFuncArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgsContext)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VALIDATORParserRULE_func)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.FuncName()
	}
	{
		p.SetState(122)
		p.Match(VALIDATORParserLEFT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3778805760) != 0 {
		{
			p.SetState(123)
			p.FuncArgs()
		}

	}
	{
		p.SetState(126)
		p.Match(VALIDATORParserRIGHT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgsContext is an interface to support dynamic dispatch.
type IFuncArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncArgsContext differentiates from other interfaces.
	IsFuncArgsContext()
}

type FuncArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgsContext() *FuncArgsContext {
	var p = new(FuncArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_funcArgs
	return p
}

func InitEmptyFuncArgsContext(p *FuncArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_funcArgs
}

func (*FuncArgsContext) IsFuncArgsContext() {}

func NewFuncArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgsContext {
	var p = new(FuncArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_funcArgs

	return p
}

func (s *FuncArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VALIDATORParserCOMMA)
}

func (s *FuncArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VALIDATORParserCOMMA, i)
}

func (s *FuncArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterFuncArgs(s)
	}
}

func (s *FuncArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitFuncArgs(s)
	}
}

func (s *FuncArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitFuncArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) FuncArgs() (localctx IFuncArgsContext) {
	localctx = NewFuncArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VALIDATORParserRULE_funcArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Expr()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VALIDATORParserCOMMA {
		{
			p.SetState(129)
			p.Match(VALIDATORParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Expr()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_funcName
	return p
}

func InitEmptyFuncNameContext(p *FuncNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_funcName
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameContext) VAR() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserVAR, 0)
}

func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (s *FuncNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitFuncName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VALIDATORParserRULE_funcName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(VALIDATORParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorHeadContext is an interface to support dynamic dispatch.
type ISelectorHeadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THIS() antlr.TerminalNode
	FieldExpr() IFieldExprContext
	SELF() antlr.TerminalNode

	// IsSelectorHeadContext differentiates from other interfaces.
	IsSelectorHeadContext()
}

type SelectorHeadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorHeadContext() *SelectorHeadContext {
	var p = new(SelectorHeadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_selectorHead
	return p
}

func InitEmptySelectorHeadContext(p *SelectorHeadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_selectorHead
}

func (*SelectorHeadContext) IsSelectorHeadContext() {}

func NewSelectorHeadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorHeadContext {
	var p = new(SelectorHeadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_selectorHead

	return p
}

func (s *SelectorHeadContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorHeadContext) THIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserTHIS, 0)
}

func (s *SelectorHeadContext) FieldExpr() IFieldExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldExprContext)
}

func (s *SelectorHeadContext) SELF() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserSELF, 0)
}

func (s *SelectorHeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorHeadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorHeadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterSelectorHead(s)
	}
}

func (s *SelectorHeadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitSelectorHead(s)
	}
}

func (s *SelectorHeadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitSelectorHead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) SelectorHead() (localctx ISelectorHeadContext) {
	localctx = NewSelectorHeadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VALIDATORParserRULE_selectorHead)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VALIDATORParserTHIS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Match(VALIDATORParserTHIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(139)
			p.Match(VALIDATORParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.FieldExpr()
		}

	case VALIDATORParserSELF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(VALIDATORParserSELF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorBodyContext is an interface to support dynamic dispatch.
type ISelectorBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BracketSelector() IBracketSelectorContext
	FieldExpr() IFieldExprContext

	// IsSelectorBodyContext differentiates from other interfaces.
	IsSelectorBodyContext()
}

type SelectorBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorBodyContext() *SelectorBodyContext {
	var p = new(SelectorBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_selectorBody
	return p
}

func InitEmptySelectorBodyContext(p *SelectorBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_selectorBody
}

func (*SelectorBodyContext) IsSelectorBodyContext() {}

func NewSelectorBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorBodyContext {
	var p = new(SelectorBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_selectorBody

	return p
}

func (s *SelectorBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorBodyContext) BracketSelector() IBracketSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBracketSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBracketSelectorContext)
}

func (s *SelectorBodyContext) FieldExpr() IFieldExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldExprContext)
}

func (s *SelectorBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterSelectorBody(s)
	}
}

func (s *SelectorBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitSelectorBody(s)
	}
}

func (s *SelectorBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitSelectorBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) SelectorBody() (localctx ISelectorBodyContext) {
	localctx = NewSelectorBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VALIDATORParserRULE_selectorBody)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VALIDATORParserLEFT_BRACKET:
		{
			p.SetState(144)
			p.BracketSelector()
		}

	case VALIDATORParserT__0:
		{
			p.SetState(145)
			p.Match(VALIDATORParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.FieldExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBracketSelectorContext is an interface to support dynamic dispatch.
type IBracketSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	Expr() IExprContext
	RIGHT_BRACKET() antlr.TerminalNode

	// IsBracketSelectorContext differentiates from other interfaces.
	IsBracketSelectorContext()
}

type BracketSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketSelectorContext() *BracketSelectorContext {
	var p = new(BracketSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_bracketSelector
	return p
}

func InitEmptyBracketSelectorContext(p *BracketSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_bracketSelector
}

func (*BracketSelectorContext) IsBracketSelectorContext() {}

func NewBracketSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketSelectorContext {
	var p = new(BracketSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_bracketSelector

	return p
}

func (s *BracketSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketSelectorContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserLEFT_BRACKET, 0)
}

func (s *BracketSelectorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BracketSelectorContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserRIGHT_BRACKET, 0)
}

func (s *BracketSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterBracketSelector(s)
	}
}

func (s *BracketSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitBracketSelector(s)
	}
}

func (s *BracketSelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitBracketSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) BracketSelector() (localctx IBracketSelectorContext) {
	localctx = NewBracketSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VALIDATORParserRULE_bracketSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(VALIDATORParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Expr()
	}
	{
		p.SetState(151)
		p.Match(VALIDATORParserRIGHT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprSelectorContext is an interface to support dynamic dispatch.
type IExprSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SelectorHead() ISelectorHeadContext
	AllSelectorBody() []ISelectorBodyContext
	SelectorBody(i int) ISelectorBodyContext

	// IsExprSelectorContext differentiates from other interfaces.
	IsExprSelectorContext()
}

type ExprSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprSelectorContext() *ExprSelectorContext {
	var p = new(ExprSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_exprSelector
	return p
}

func InitEmptyExprSelectorContext(p *ExprSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_exprSelector
}

func (*ExprSelectorContext) IsExprSelectorContext() {}

func NewExprSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprSelectorContext {
	var p = new(ExprSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_exprSelector

	return p
}

func (s *ExprSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprSelectorContext) SelectorHead() ISelectorHeadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorHeadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorHeadContext)
}

func (s *ExprSelectorContext) AllSelectorBody() []ISelectorBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorBodyContext); ok {
			len++
		}
	}

	tst := make([]ISelectorBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorBodyContext); ok {
			tst[i] = t.(ISelectorBodyContext)
			i++
		}
	}

	return tst
}

func (s *ExprSelectorContext) SelectorBody(i int) ISelectorBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorBodyContext)
}

func (s *ExprSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterExprSelector(s)
	}
}

func (s *ExprSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitExprSelector(s)
	}
}

func (s *ExprSelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitExprSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) ExprSelector() (localctx IExprSelectorContext) {
	localctx = NewExprSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VALIDATORParserRULE_exprSelector)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.SelectorHead()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(154)
				p.SelectorBody()
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldExprContext is an interface to support dynamic dispatch.
type IFieldExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode

	// IsFieldExprContext differentiates from other interfaces.
	IsFieldExprContext()
}

type FieldExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldExprContext() *FieldExprContext {
	var p = new(FieldExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_fieldExpr
	return p
}

func InitEmptyFieldExprContext(p *FieldExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_fieldExpr
}

func (*FieldExprContext) IsFieldExprContext() {}

func NewFieldExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldExprContext {
	var p = new(FieldExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_fieldExpr

	return p
}

func (s *FieldExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldExprContext) VAR() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserVAR, 0)
}

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterFieldExpr(s)
	}
}

func (s *FieldExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitFieldExpr(s)
	}
}

func (s *FieldExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitFieldExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) FieldExpr() (localctx IFieldExprContext) {
	localctx = NewFieldExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, VALIDATORParserRULE_fieldExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(VALIDATORParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMsgContext is an interface to support dynamic dispatch.
type IMsgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sprintf() ISprintfContext

	// IsMsgContext differentiates from other interfaces.
	IsMsgContext()
}

type MsgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgContext() *MsgContext {
	var p = new(MsgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_msg
	return p
}

func InitEmptyMsgContext(p *MsgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_msg
}

func (*MsgContext) IsMsgContext() {}

func NewMsgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgContext {
	var p = new(MsgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_msg

	return p
}

func (s *MsgContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgContext) Sprintf() ISprintfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISprintfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISprintfContext)
}

func (s *MsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterMsg(s)
	}
}

func (s *MsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitMsg(s)
	}
}

func (s *MsgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitMsg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Msg() (localctx IMsgContext) {
	localctx = NewMsgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VALIDATORParserRULE_msg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(VALIDATORParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Sprintf()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISprintfContext is an interface to support dynamic dispatch.
type ISprintfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_PARENTHESIS() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RIGHT_PARENTHESIS() antlr.TerminalNode
	AllSprintfArgs() []ISprintfArgsContext
	SprintfArgs(i int) ISprintfArgsContext

	// IsSprintfContext differentiates from other interfaces.
	IsSprintfContext()
}

type SprintfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySprintfContext() *SprintfContext {
	var p = new(SprintfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_sprintf
	return p
}

func InitEmptySprintfContext(p *SprintfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_sprintf
}

func (*SprintfContext) IsSprintfContext() {}

func NewSprintfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SprintfContext {
	var p = new(SprintfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_sprintf

	return p
}

func (s *SprintfContext) GetParser() antlr.Parser { return s.parser }

func (s *SprintfContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserLEFT_PARENTHESIS, 0)
}

func (s *SprintfContext) STRING() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserSTRING, 0)
}

func (s *SprintfContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserRIGHT_PARENTHESIS, 0)
}

func (s *SprintfContext) AllSprintfArgs() []ISprintfArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISprintfArgsContext); ok {
			len++
		}
	}

	tst := make([]ISprintfArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISprintfArgsContext); ok {
			tst[i] = t.(ISprintfArgsContext)
			i++
		}
	}

	return tst
}

func (s *SprintfContext) SprintfArgs(i int) ISprintfArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISprintfArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISprintfArgsContext)
}

func (s *SprintfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SprintfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SprintfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterSprintf(s)
	}
}

func (s *SprintfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitSprintf(s)
	}
}

func (s *SprintfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitSprintf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) Sprintf() (localctx ISprintfContext) {
	localctx = NewSprintfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VALIDATORParserRULE_sprintf)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(VALIDATORParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(VALIDATORParserLEFT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(VALIDATORParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VALIDATORParserCOMMA {
		{
			p.SetState(168)
			p.SprintfArgs()
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)
		p.Match(VALIDATORParserRIGHT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISprintfArgsContext is an interface to support dynamic dispatch.
type ISprintfArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Expr() IExprContext

	// IsSprintfArgsContext differentiates from other interfaces.
	IsSprintfArgsContext()
}

type SprintfArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySprintfArgsContext() *SprintfArgsContext {
	var p = new(SprintfArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_sprintfArgs
	return p
}

func InitEmptySprintfArgsContext(p *SprintfArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_sprintfArgs
}

func (*SprintfArgsContext) IsSprintfArgsContext() {}

func NewSprintfArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SprintfArgsContext {
	var p = new(SprintfArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_sprintfArgs

	return p
}

func (s *SprintfArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *SprintfArgsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserCOMMA, 0)
}

func (s *SprintfArgsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SprintfArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SprintfArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SprintfArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterSprintfArgs(s)
	}
}

func (s *SprintfArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitSprintfArgs(s)
	}
}

func (s *SprintfArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitSprintfArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) SprintfArgs() (localctx ISprintfArgsContext) {
	localctx = NewSprintfArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VALIDATORParserRULE_sprintfArgs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(VALIDATORParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompareOperatorContext is an interface to support dynamic dispatch.
type ICompareOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GT() antlr.TerminalNode
	EGT() antlr.TerminalNode
	LT() antlr.TerminalNode
	ELT() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NOTEQ() antlr.TerminalNode

	// IsCompareOperatorContext differentiates from other interfaces.
	IsCompareOperatorContext()
}

type CompareOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareOperatorContext() *CompareOperatorContext {
	var p = new(CompareOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_compareOperator
	return p
}

func InitEmptyCompareOperatorContext(p *CompareOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VALIDATORParserRULE_compareOperator
}

func (*CompareOperatorContext) IsCompareOperatorContext() {}

func NewCompareOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOperatorContext {
	var p = new(CompareOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VALIDATORParserRULE_compareOperator

	return p
}

func (s *CompareOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserGT, 0)
}

func (s *CompareOperatorContext) EGT() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserEGT, 0)
}

func (s *CompareOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserLT, 0)
}

func (s *CompareOperatorContext) ELT() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserELT, 0)
}

func (s *CompareOperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserEQ, 0)
}

func (s *CompareOperatorContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(VALIDATORParserNOTEQ, 0)
}

func (s *CompareOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.EnterCompareOperator(s)
	}
}

func (s *CompareOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VALIDATORListener); ok {
		listenerT.ExitCompareOperator(s)
	}
}

func (s *CompareOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VALIDATORVisitor:
		return t.VisitCompareOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VALIDATORParser) CompareOperator() (localctx ICompareOperatorContext) {
	localctx = NewCompareOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VALIDATORParserRULE_compareOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VALIDATORParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *AddingExprContext = nil
		if localctx != nil {
			t = localctx.(*AddingExprContext)
		}
		return p.AddingExpr_Sempred(t, predIndex)

	case 6:
		var t *MultiplyingExprContext = nil
		if localctx != nil {
			t = localctx.(*MultiplyingExprContext)
		}
		return p.MultiplyingExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VALIDATORParser) AddingExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *VALIDATORParser) MultiplyingExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
