// Code generated from ./validator/grammar/VALIDATOR.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // VALIDATOR
import "github.com/antlr4-go/antlr/v4"

// BaseVALIDATORListener is a complete listener for a parse tree produced by VALIDATORParser.
type BaseVALIDATORListener struct{}

var _ VALIDATORListener = &BaseVALIDATORListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVALIDATORListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVALIDATORListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVALIDATORListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVALIDATORListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterValidator is called when production validator is entered.
func (s *BaseVALIDATORListener) EnterValidator(ctx *ValidatorContext) {}

// ExitValidator is called when production validator is exited.
func (s *BaseVALIDATORListener) ExitValidator(ctx *ValidatorContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseVALIDATORListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseVALIDATORListener) ExitExpr(ctx *ExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseVALIDATORListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseVALIDATORListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterConditionalExpr is called when production conditionalExpr is entered.
func (s *BaseVALIDATORListener) EnterConditionalExpr(ctx *ConditionalExprContext) {}

// ExitConditionalExpr is called when production conditionalExpr is exited.
func (s *BaseVALIDATORListener) ExitConditionalExpr(ctx *ConditionalExprContext) {}

// EnterComparingExpr is called when production comparingExpr is entered.
func (s *BaseVALIDATORListener) EnterComparingExpr(ctx *ComparingExprContext) {}

// ExitComparingExpr is called when production comparingExpr is exited.
func (s *BaseVALIDATORListener) ExitComparingExpr(ctx *ComparingExprContext) {}

// EnterAddingExpr is called when production addingExpr is entered.
func (s *BaseVALIDATORListener) EnterAddingExpr(ctx *AddingExprContext) {}

// ExitAddingExpr is called when production addingExpr is exited.
func (s *BaseVALIDATORListener) ExitAddingExpr(ctx *AddingExprContext) {}

// EnterMultiplyingExpr is called when production multiplyingExpr is entered.
func (s *BaseVALIDATORListener) EnterMultiplyingExpr(ctx *MultiplyingExprContext) {}

// ExitMultiplyingExpr is called when production multiplyingExpr is exited.
func (s *BaseVALIDATORListener) ExitMultiplyingExpr(ctx *MultiplyingExprContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseVALIDATORListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseVALIDATORListener) ExitFactor(ctx *FactorContext) {}

// EnterValue is called when production value is entered.
func (s *BaseVALIDATORListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseVALIDATORListener) ExitValue(ctx *ValueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseVALIDATORListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseVALIDATORListener) ExitConstant(ctx *ConstantContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseVALIDATORListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseVALIDATORListener) ExitFunc(ctx *FuncContext) {}

// EnterFuncArgs is called when production funcArgs is entered.
func (s *BaseVALIDATORListener) EnterFuncArgs(ctx *FuncArgsContext) {}

// ExitFuncArgs is called when production funcArgs is exited.
func (s *BaseVALIDATORListener) ExitFuncArgs(ctx *FuncArgsContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BaseVALIDATORListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BaseVALIDATORListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterSelectorHead is called when production selectorHead is entered.
func (s *BaseVALIDATORListener) EnterSelectorHead(ctx *SelectorHeadContext) {}

// ExitSelectorHead is called when production selectorHead is exited.
func (s *BaseVALIDATORListener) ExitSelectorHead(ctx *SelectorHeadContext) {}

// EnterSelectorBody is called when production selectorBody is entered.
func (s *BaseVALIDATORListener) EnterSelectorBody(ctx *SelectorBodyContext) {}

// ExitSelectorBody is called when production selectorBody is exited.
func (s *BaseVALIDATORListener) ExitSelectorBody(ctx *SelectorBodyContext) {}

// EnterBracketSelector is called when production bracketSelector is entered.
func (s *BaseVALIDATORListener) EnterBracketSelector(ctx *BracketSelectorContext) {}

// ExitBracketSelector is called when production bracketSelector is exited.
func (s *BaseVALIDATORListener) ExitBracketSelector(ctx *BracketSelectorContext) {}

// EnterExprSelector is called when production exprSelector is entered.
func (s *BaseVALIDATORListener) EnterExprSelector(ctx *ExprSelectorContext) {}

// ExitExprSelector is called when production exprSelector is exited.
func (s *BaseVALIDATORListener) ExitExprSelector(ctx *ExprSelectorContext) {}

// EnterFieldExpr is called when production fieldExpr is entered.
func (s *BaseVALIDATORListener) EnterFieldExpr(ctx *FieldExprContext) {}

// ExitFieldExpr is called when production fieldExpr is exited.
func (s *BaseVALIDATORListener) ExitFieldExpr(ctx *FieldExprContext) {}

// EnterMsg is called when production msg is entered.
func (s *BaseVALIDATORListener) EnterMsg(ctx *MsgContext) {}

// ExitMsg is called when production msg is exited.
func (s *BaseVALIDATORListener) ExitMsg(ctx *MsgContext) {}

// EnterSprintf is called when production sprintf is entered.
func (s *BaseVALIDATORListener) EnterSprintf(ctx *SprintfContext) {}

// ExitSprintf is called when production sprintf is exited.
func (s *BaseVALIDATORListener) ExitSprintf(ctx *SprintfContext) {}

// EnterSprintfArgs is called when production sprintfArgs is entered.
func (s *BaseVALIDATORListener) EnterSprintfArgs(ctx *SprintfArgsContext) {}

// ExitSprintfArgs is called when production sprintfArgs is exited.
func (s *BaseVALIDATORListener) ExitSprintfArgs(ctx *SprintfArgsContext) {}

// EnterCompareOperator is called when production compareOperator is entered.
func (s *BaseVALIDATORListener) EnterCompareOperator(ctx *CompareOperatorContext) {}

// ExitCompareOperator is called when production compareOperator is exited.
func (s *BaseVALIDATORListener) ExitCompareOperator(ctx *CompareOperatorContext) {}
