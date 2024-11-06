grammar VALIDATOR;

// compares like len($)>1 && regexp('^\\w*$')
// msg like msg:sprintf('invalid d: %v', $)
// for preview
// validator: expr (SEMICOLON msg)? (('\r'? '\n') expr)* EOF;
validator: expr (SEMICOLON msg)? EOF;

expr: andExpr (OR expr)?
    ;

andExpr: conditionalExpr (AND andExpr)? ;

conditionalExpr: comparingExpr (QUESTION_MARK expr COLON conditionalExpr)?;

comparingExpr: addingExpr (compareOperator comparingExpr)?;

addingExpr: multiplyingExpr | addingExpr PLUS multiplyingExpr | addingExpr MINUS multiplyingExpr;

multiplyingExpr: factor | multiplyingExpr ASTERISK factor | multiplyingExpr SLASH factor | multiplyingExpr MODULO factor;

factor: value | func;

value: exprSelector
    | constant
    | NIL
    | LEFT_PARENTHESIS expr RIGHT_PARENTHESIS
    ;

constant: STRING
        | INTEGER_VALUE
        | DOUBLE_VALUE
        ;

func: funcName LEFT_PARENTHESIS funcArgs? RIGHT_PARENTHESIS;

funcArgs: expr (COMMA expr)*;

funcName: VAR;

selectorHead: (THIS ('.' fieldExpr)) | SELF;

selectorBody: (bracketSelector | ('.' fieldExpr));

bracketSelector: LEFT_BRACKET expr RIGHT_BRACKET
               ;

exprSelector: selectorHead selectorBody*;

fieldExpr: VAR;

msg: 'msg:' sprintf;

// like sprintf('invalid d: %v', $)
sprintf: 'sprintf' LEFT_PARENTHESIS STRING sprintfArgs* RIGHT_PARENTHESIS;

sprintfArgs: (',' expr);

compareOperator: GT|EGT|LT|ELT|EQ|NOTEQ;

GT: '>';
EGT: '>=';
LT: '<';
ELT: '<=';
EQ: '==';
NOTEQ: '!=';
PLUS     : '+' ;
MINUS    : '-' ;
ASTERISK : '*' ;
SLASH    : '/' ;
MODULO  : '%' ;
COLON: ':';
SEMICOLON: ';';
COMMA: ',';
SELF: '$';
THIS: 'this';
NIL: 'nil';
VAR: ('_' [a-zA-Z0-9]+) | ([a-zA-Z] [_a-zA-Z0-9]*);
AND: '&&';
OR: '||';
LEFT_PARENTHESIS: '(';
RIGHT_PARENTHESIS: ')';
LEFT_BRACKET: '[';
RIGHT_BRACKET: ']';
QUESTION_MARK: '?';
STRING                        : '\'' ( ~'\'' | '\'\'' )* '\'' ;
INTEGER_VALUE                 : DECIMAL_INTEGER ;
DOUBLE_VALUE                  : DIGIT+ ('.' DIGIT*)? EXPONENT?
                              | '.' DIGIT+ EXPONENT?
                              ;
// fragments for literal primitives
fragment DECIMAL_INTEGER      : DIGIT ('_'? DIGIT)* ;
fragment EXPONENT             : 'E' [+-]? DIGIT+ ;
fragment DIGIT                : [0-9];
WS: [ \t\n\r] -> skip;


