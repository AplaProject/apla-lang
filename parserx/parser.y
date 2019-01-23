%{
package parser

// import "fmt"

%}

%union {
    b       bool
    i       int
    f       float64
    s       string
    a       []string
}

// Identifiers + literals
%token<s> IDENT  // foobar
%token<i> INT    // 314
%token<f> FLOAT  // 3.14
%token<s> STRING // "string"
%token<b> TRUE   // true
%token<b> FALSE  // false

// Operators
%token ADD // +
%token SUB // -
%token MUL // *
%token DIV // /
%token MOD // %

%token ASSIGN     // =
%token ADD_ASSIGN // +=
%token SUB_ASSIGN // -=
%token MUL_ASSIGN // *=
%token DIV_ASSIGN // /=
%token MOD_ASSIGN // %=

%token AND // &&
%token OR  // ||

%token INC // ++
%token DEC // --

%token EQ     // ==
%token NOT_EQ // !=
%token NOT    // !

%token LT     // <
%token GT     // >
%token LTE    // <=
%token GTE    // >=

%token ELLIPSIS // ...

// Delimiters
%token DOT     // .
%token COMMA   // ,
%token COLON   // :
%token LPAREN  // (
%token RPAREN  // )
%token LBRACE  // {
%token RBRACE  // }
%token LBRAKET // [
%token RBRAKET // ]

// Keywords
%token CONTRACT   // contract
%token DATA       // data
%token CONDITION  // condition
%token ACTION     // action
%token FUNC       // func
%token VAR        // var
%token EXTEND_VAR // $foo
%token IF         // if
%token ELSE       // else
%token WHILE      // while
%token BREAK      // break
%token CONTINUE   // continue
%token INFO       // info
%token WARNING    // warning
%token ERROR      // error
%token NIL        // nil
%token RETURN     // return

// Types
%token T_BOOL   // bool
%token T_INT    // int
%token T_FLOAT  // float
%token T_MONEY  // money
%token T_STRING // string
%token T_BYTES  // bytes
%token T_ARRAY  // array
%token T_MAP    // map
%token T_FILE   // file

%start translation_unit

%nonassoc LOWER_THAN_LPAREN
%nonassoc LESS_THAN_ELSE
%nonassoc LPAREN
%nonassoc ELSE

%%

primary_expression
    : IDENT LPAREN argument_expression_list RPAREN
    | IDENT LPAREN RPAREN
    | IDENT %prec LOWER_THAN_LPAREN ;
    | INT
    | STRING
    | LPAREN expression RPAREN
    ;

postfix_expression
    : primary_expression
    | postfix_expression LBRAKET expression RBRAKET
    ;

argument_expression_list
    : assignment_expression
    | argument_expression_list COMMA assignment_expression
    ;

unary_expression
    : postfix_expression
    | INC unary_expression
    | DEC unary_expression
    | unary_operator unary_expression
    ;

unary_operator
    : ADD
    | SUB
    | NOT
    ;
/*
multiplicative_expression
    : unary_expression
    | multiplicative_expression MUL unary_expression
    | multiplicative_expression DIV unary_expression
    | multiplicative_expression MOD unary_expression
    ;

additive_expression
    : multiplicative_expression
    | additive_expression ADD multiplicative_expression
    | additive_expression SUB multiplicative_expression
    ;

relational_expression
    : additive_expression
    | relational_expression LT additive_expression
    | relational_expression GT additive_expression
    | relational_expression LTE additive_expression
    | relational_expression GTE additive_expression
    ;

equality_expression
    : relational_expression
    | equality_expression EQ relational_expression
    | equality_expression NOT_EQ relational_expression
    ;

logical_and_expression
    : equality_expression
    | logical_and_expression AND equality_expression
    ;

logical_or_expression
    : logical_and_expression
    | logical_or_expression OR logical_and_expression
    ;
*/
assignment_expression
//    : logical_or_expression
    : unary_expression
    | IDENT assignment_operator assignment_expression
    ;

assignment_operator:
    ASSIGN
    ;

expression
    : assignment_expression
    | expression COMMA assignment_expression
    ;

block_item_list
    : block_item
    | block_item_list block_item
    ;

compound_statement
    : LBRACE RBRACE
    | LBRACE block_item_list RBRACE
    ;

block_item
    : declaration
    | statement
    ;

type
    : T_BOOL
    | T_INT
    | T_FLOAT
    | T_MONEY
    | T_STRING
    | T_BYTES
    | T_ARRAY
    | T_MAP
    | T_FILE
    ;

ident_list
    : IDENT
    | ident_list COMMA IDENT
    ;

variable_declaration:
    VAR ident_list type

variable_specifier
    : ident_list type
    | variable_specifier COMMA ident_list type
    | variable_specifier COMMA IDENT ELLIPSIS
    ;

function_argument_list
    : LPAREN RPAREN
    | LPAREN variable_specifier RPAREN
    ;

function_signature
    : function_argument_list type
    | function_argument_list
    ;

function_declaration:
    FUNC IDENT function_signature compound_statement

declaration
    : function_declaration
    | variable_declaration
    ;

statement
    : compound_statement
    | expression
    | selection_statement
    | iteration_statement
    | jump_statement
    ;

selection_statement
    : IF expression statement %prec LESS_THAN_ELSE
    | IF expression statement ELSE statement
    ;

iteration_statement
    : WHILE expression statement
    ;

jump_statement
    : CONTINUE
    | BREAK
    | INFO expression
    | WARNING expression
    | ERROR expression
    | RETURN /* empty */ 
//    | RETURN expression
    ;

contract_argument
    : IDENT type STRING
    | IDENT type
    ;

contract_argument_list
    : contract_argument
    | contract_argument_list contract_argument
    ;

contract_argument_declaration
    : DATA LBRACE contract_argument_list RBRACE
    | DATA LBRACE RBRACE
    ;

contract_condition_declaration:
    CONDITION compound_statement

contract_action_declaration:
    ACTION compound_statement

contract_specifier
    : contract_argument_declaration
    | contract_condition_declaration
    | contract_action_declaration
    | function_declaration
    ;

contract_specifiers
    : contract_specifier
    | contract_specifiers contract_specifier
    ;

contract_declaration
    : CONTRACT IDENT LBRACE contract_specifiers RBRACE
    | CONTRACT IDENT LBRACE RBRACE
    ;

top_declaration
    : function_declaration
    | contract_declaration
    ;

translation_unit
    : top_declaration
    | translation_unit top_declaration
    ;