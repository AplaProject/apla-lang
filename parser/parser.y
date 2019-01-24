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
%type <nb> contract_body
%type <n> contract_declaration

%start contract_declaration

%%

type
    : T_BOOL {$$ = 1}
    | T_INT {$$ = 2}
    ;

ident_list
    : IDENT { $$ = []string{$1} }
    | ident_list IDENT { $$ = append($1, $2) }
    ;

var_declaration
    : type ident_list {
        va := make([]NVar, len($2))
        for i, name := range $2 {
            va[i] = NVar{
                Type: $1,
                Name: name,
            }
        }
        $$ = va
    }
    ;

var_declarations
    : var_declaration { $$=$1 }
    | var_declarations var_declaration { $$ = append($1, $2...) }
    ;

func_body 
    : /*empty*/
    | INT {
        fmt.Println("FUNC")
    }
    ;

contract_data 
    : /*empty*/ { $$ = nil }
    | DATA LBRACE var_declarations RBRACE { $$ = $3 }
    ;

contract_body 
    : contract_data func_body {
        $$ = newBlock($1)
        fmt.Println("BODY", $1)
    }
    ;

contract_declaration 
    : CONTRACT IDENT LBRACE contract_body RBRACE { 
        $$ = newContract($2, $4)
        setResult(yylex, $$)
        }
    ;  
