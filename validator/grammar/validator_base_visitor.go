// Code generated from ./validator/grammar/VALIDATOR.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // VALIDATOR
import "github.com/antlr4-go/antlr/v4"

type BaseVALIDATORVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVALIDATORVisitor) VisitValidator(ctx *ValidatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitConditionalExpr(ctx *ConditionalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitComparingExpr(ctx *ComparingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitAddingExpr(ctx *AddingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitMultiplyingExpr(ctx *MultiplyingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitFuncArgs(ctx *FuncArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitFuncName(ctx *FuncNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitSelectorHead(ctx *SelectorHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitSelectorBody(ctx *SelectorBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitBracketSelector(ctx *BracketSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitExprSelector(ctx *ExprSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitFieldExpr(ctx *FieldExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitMsg(ctx *MsgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitSprintf(ctx *SprintfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitSprintfArgs(ctx *SprintfArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVALIDATORVisitor) VisitCompareOperator(ctx *CompareOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
