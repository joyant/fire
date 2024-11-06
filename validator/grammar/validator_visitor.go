// Code generated from ./validator/grammar/VALIDATOR.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // VALIDATOR
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VALIDATORParser.
type VALIDATORVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VALIDATORParser#validator.
	VisitValidator(ctx *ValidatorContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#conditionalExpr.
	VisitConditionalExpr(ctx *ConditionalExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#comparingExpr.
	VisitComparingExpr(ctx *ComparingExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#addingExpr.
	VisitAddingExpr(ctx *AddingExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#multiplyingExpr.
	VisitMultiplyingExpr(ctx *MultiplyingExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#funcArgs.
	VisitFuncArgs(ctx *FuncArgsContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#funcName.
	VisitFuncName(ctx *FuncNameContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#selectorHead.
	VisitSelectorHead(ctx *SelectorHeadContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#selectorBody.
	VisitSelectorBody(ctx *SelectorBodyContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#bracketSelector.
	VisitBracketSelector(ctx *BracketSelectorContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#exprSelector.
	VisitExprSelector(ctx *ExprSelectorContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#fieldExpr.
	VisitFieldExpr(ctx *FieldExprContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#msg.
	VisitMsg(ctx *MsgContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#sprintf.
	VisitSprintf(ctx *SprintfContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#sprintfArgs.
	VisitSprintfArgs(ctx *SprintfArgsContext) interface{}

	// Visit a parse tree produced by VALIDATORParser#compareOperator.
	VisitCompareOperator(ctx *CompareOperatorContext) interface{}
}
