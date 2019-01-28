%{
package parser

import "fmt"

func setResult(l yyLexer, v *Node) {
  l.(*lexer).result = v
}

%}

%union {
    n       *Node
    b       bool
    i       int
    f       float64
    s       string
    sa      []string
    va      []NVar
    nb      *NBlock
}

// Identifiers + literals
%token<s> IDENT  // foobar
%token<i> INT    // 314
%token<b> TRUE   // true
%token<b> FALSE  // false

// Delimiters
%token NEWLINE // \n
%token COMMA   // ,
%token COLON   // :
%token LPAREN  // (
%token RPAREN  // )
%token LBRACE  // {
%token RBRACE  // }
%token LBRAKET // [
%token RBRAKET // ]

// Operators
%token ADD // +
%token SUB // -
%token MUL // *
%token DIV // /
%token MOD // %

// Keywords
%token DATA       // data
%token CONTRACT   // contract

// Types
%token T_BOOL   // bool
%token T_INT    // int

%type <i> type
%type <sa> ident_list
%type <va> var_declaration
%type <va> var_declarations
%type <va> contract_data
%type <i> expr
%type <n> contract_body
%type <n> contract_declaration

%left ADD SUB
%left MUL DIV MOD
%right UNARYMINUS 

%start contract_declaration

%%

type
    : T_BOOL {$$ = 1}
    | T_INT {$$ = 2}
    ;

statements
    : /*empty*/
    | statements NEWLINE
    | statements statement NEWLINE 
    ;

statement 
    : expr
    ;

expr
    : LPAREN expr RPAREN { $$ = $2; fmt.Println(`PAR`,$2) }
    | INT { $$ = $1; fmt.Println(`INT`,$1)}
    | expr MUL expr { $$ = $1*$3; fmt.Println(`MUL`,$1, $3) }
    | expr DIV expr { /*$$ = $1/$3; */ }
    | expr ADD expr { $$ = $1+$3; fmt.Println(`ADD`,$1, $3) }
    | expr SUB expr { /*$$ = $1-$3; */}
    | expr MOD expr { /*$$ = $1-$3; */}
    | SUB expr %prec UNARYMINUS { $$ = -$2; fmt.Println(`NEG`,$2)}
    ;

ident_list
    : IDENT { $$ = []string{$1} }
    | ident_list IDENT { $$ = append($1, $2) }
    ;

var_declaration
    : type ident_list { $$ = newVars($1, $2) }
    ;

var_declarations
    : var_declaration { $$=$1 }
    | var_declarations NEWLINE var_declaration { $$ = append($1, $3...) }
    ;

contract_data 
    : /*empty*/ { $$ = nil }
    | DATA LBRACE var_declarations NEWLINE RBRACE NEWLINE { $$ = $3 }
    ;

contract_body 
    : contract_data statements {
        $$ = newBlock($1, yylex)
        fmt.Println("BODY", $$)
    }
    ;

contract_declaration 
    : CONTRACT IDENT LBRACE contract_body RBRACE { 
        $$ = newContract($2, $4, yylex)
        setResult(yylex, $$)
        }
    ;  
