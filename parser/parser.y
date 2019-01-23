%{
package parser

import "fmt"

%}

%union {
    b       bool
    i       int
    f       float64
    s       string
}

// Identifiers + literals
%token<s> IDENT  // foobar
%token<i> INT    // 314
%token<b> TRUE   // true
%token<b> FALSE  // false

// Delimiters
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

%type <s> contract_declaration

%start translation_unit

%%

func_body 
    : /*empty*/
    | INT
    ;

contract_data 
    : /*empty*/
    | DATA LBRACE RBRACE
    ;

contract_body 
    : contract_data func_body
    ;

contract_declaration 
    : CONTRACT IDENT LBRACE contract_body RBRACE { 
        $$ = "OOOOPS"
        fmt.Println("OOOPS", $$, $2)
        }
    ;  

translation_unit
    : contract_declaration
    ;