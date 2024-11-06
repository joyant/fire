// Code generated from ./validator/grammar/VALIDATOR.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // VALIDATOR
import "github.com/antlr4-go/antlr/v4"

// VALIDATORListener is a complete listener for a parse tree produced by VALIDATORParser.
type VALIDATORListener interface {
	antlr.ParseTreeListener

	// EnterValidator is called when entering the validator production.
	EnterValidator(c *ValidatorContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterConditionalExpr is called when entering the conditionalExpr production.
	EnterConditionalExpr(c *ConditionalExprContext)

	// EnterComparingExpr is called when entering the comparingExpr production.
	EnterComparingExpr(c *ComparingExprContext)

	// EnterAddingExpr is called when entering the addingExpr production.
	EnterAddingExpr(c *AddingExprContext)

	// EnterMultiplyingExpr is called when entering the multiplyingExpr production.
	EnterMultiplyingExpr(c *MultiplyingExprContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterFuncArgs is called when entering the funcArgs production.
	EnterFuncArgs(c *FuncArgsContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterSelectorHead is called when entering the selectorHead production.
	EnterSelectorHead(c *SelectorHeadContext)

	// EnterSelectorBody is called when entering the selectorBody production.
	EnterSelectorBody(c *SelectorBodyContext)

	// EnterBracketSelector is called when entering the bracketSelector production.
	EnterBracketSelector(c *BracketSelectorContext)

	// EnterExprSelector is called when entering the exprSelector production.
	EnterExprSelector(c *ExprSelectorContext)

	// EnterFieldExpr is called when entering the fieldExpr production.
	EnterFieldExpr(c *FieldExprContext)

	// EnterMsg is called when entering the msg production.
	EnterMsg(c *MsgContext)

	// EnterSprintf is called when entering the sprintf production.
	EnterSprintf(c *SprintfContext)

	// EnterSprintfArgs is called when entering the sprintfArgs production.
	EnterSprintfArgs(c *SprintfArgsContext)

	// EnterCompareOperator is called when entering the compareOperator production.
	EnterCompareOperator(c *CompareOperatorContext)

	// ExitValidator is called when exiting the validator production.
	ExitValidator(c *ValidatorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitConditionalExpr is called when exiting the conditionalExpr production.
	ExitConditionalExpr(c *ConditionalExprContext)

	// ExitComparingExpr is called when exiting the comparingExpr production.
	ExitComparingExpr(c *ComparingExprContext)

	// ExitAddingExpr is called when exiting the addingExpr production.
	ExitAddingExpr(c *AddingExprContext)

	// ExitMultiplyingExpr is called when exiting the multiplyingExpr production.
	ExitMultiplyingExpr(c *MultiplyingExprContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitFuncArgs is called when exiting the funcArgs production.
	ExitFuncArgs(c *FuncArgsContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitSelectorHead is called when exiting the selectorHead production.
	ExitSelectorHead(c *SelectorHeadContext)

	// ExitSelectorBody is called when exiting the selectorBody production.
	ExitSelectorBody(c *SelectorBodyContext)

	// ExitBracketSelector is called when exiting the bracketSelector production.
	ExitBracketSelector(c *BracketSelectorContext)

	// ExitExprSelector is called when exiting the exprSelector production.
	ExitExprSelector(c *ExprSelectorContext)

	// ExitFieldExpr is called when exiting the fieldExpr production.
	ExitFieldExpr(c *FieldExprContext)

	// ExitMsg is called when exiting the msg production.
	ExitMsg(c *MsgContext)

	// ExitSprintf is called when exiting the sprintf production.
	ExitSprintf(c *SprintfContext)

	// ExitSprintfArgs is called when exiting the sprintfArgs production.
	ExitSprintfArgs(c *SprintfArgsContext)

	// ExitCompareOperator is called when exiting the compareOperator production.
	ExitCompareOperator(c *CompareOperatorContext)
}
