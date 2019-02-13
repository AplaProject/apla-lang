%{
package parser

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
}

// Identifiers + literals
%token<s> IDENT  // foobar
%token<s> CALL  // foobar(
%token<i> INT    // 314
%token<s> STRING  // "string"
%token<s> QSTRING  // `string`
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
%token QUESTION // ?

// Operators
%token ADD // +
%token SUB // -
%token MUL // *
%token DIV // /
%token MOD // %

%token ADD_ASSIGN // +=
%token SUB_ASSIGN // -=
%token MUL_ASSIGN // *=
%token DIV_ASSIGN // /=
%token MOD_ASSIGN // %=
%token ASSIGN // =

%token AND // &&
%token OR  // ||

%token EQ     // ==
%token NOT_EQ // !=
%token NOT    // !

%token LT     // <
%token GT     // >
%token LTE    // <=
%token GTE    // >=

// Keywords
%token DATA       // data
%token CONTRACT   // contract
%token IF       // if
%token ELIF     // elif
%token ELSE     // else
%token RETURN   // return
%token WHILE   // while
%token FUNC    // func

// Types
%token T_INT    // int
%token T_BOOL   // bool
%token T_STR  // str

%type <i> type
%type <i> rettype
%type <sa> ident_list
%type <va> var_declaration
%type <va> var_declarations
%type <va> par_declaration
%type <va> par_declarations
%type <va> contract_data
%type <n> params
%type <n> var
%type <n> expr
%type <n> elif
%type <n> else
%type <n> statement
%type <n> statements
%type <n> contract_body
%type <n> contract_declaration

%left AND 
%left OR
%left LTE GTE LT GT EQ NOT_EQ
%left ADD SUB
%left MUL DIV MOD
%right UNARYMINUS UNARYNOT

%start contract_declaration

%%

type
    : T_BOOL {$$ = VBool}
    | T_INT {$$ = VInt}
    | T_STR {$$ = VStr}
    ;

rettype
    : /*empty*/ { $$ = VVoid }
    | type { $$ = $1 }
    ;

statements
    : /*empty*/ { $$ = nil }
    | statements NEWLINE { $$ = $1 }
    | statements statement NEWLINE { $$ = addStatement($1, $2, yylex)}
    ;

params
    : /*empty*/ { $$ = nil }
    | expr { $$ = newParam( $1, yylex ) }
    | params COMMA expr { $$ = addParam($1, $3)}
    ;

var 
    : IDENT { $$ = newVarValue($1, yylex); }

else 
   : /*empty*/ {$$ = nil}
   | ELSE LBRACE statements RBRACE { $$ = $3 }
   ;

elif
   : /*empty*/ {$$ = nil}
   | elif ELIF expr LBRACE statements RBRACE { $$ = newElif($1, $3, $5, yylex) }
   ;

statement 
    : var ASSIGN expr { $$ = newBinary($1, $3, ASSIGN, yylex) }
    | var ADD_ASSIGN expr { $$ = newBinary($1, $3, ADD_ASSIGN, yylex); }
    | var SUB_ASSIGN expr { $$ = newBinary($1, $3, SUB_ASSIGN, yylex); }
    | var MUL_ASSIGN expr { $$ = newBinary($1, $3, MUL_ASSIGN, yylex); }
    | var DIV_ASSIGN expr { $$ = newBinary($1, $3, DIV_ASSIGN, yylex); }
    | var MOD_ASSIGN expr { $$ = newBinary($1, $3, MOD_ASSIGN, yylex); } 
    | type IDENT ASSIGN expr { $$ = newBinary( newVarDecl( $1, []string{$2}, yylex ), $4, ASSIGN, yylex) }
    | type ident_list { $$ = newVarDecl( $1, $2, yylex )}
    | IF expr LBRACE statements RBRACE elif else { $$ = newIf( $2, $4, $6, $7, yylex )}
    | RETURN { $$ = newReturn(nil, yylex); }
    | RETURN expr { $$ = newReturn($2, yylex); }
    | WHILE expr LBRACE statements RBRACE { $$ = newWhile( $2, $4, yylex )}
    | FUNC CALL par_declarations RPAREN rettype LBRACE statements RBRACE { 
           $$ = newFunc($2, $3, $5, $7, yylex)
           }
    | CALL params RPAREN { $$ = newCallFunc($1, $2, yylex);}
    ;

expr
    : LPAREN expr RPAREN { $$ = $2; }
    | INT { $$ = newValue($1, yylex);}
    | STRING { $$ = newValue($1, yylex);}
    | QSTRING { $$ = newValue($1, yylex);}
    | TRUE { $$ = newValue(true, yylex);}
    | FALSE { $$ = newValue(false, yylex);}
    | CALL params RPAREN { $$ = newCallFunc($1, $2, yylex);}
    | IDENT { $$ = newGetVar($1, yylex);}
    | QUESTION LPAREN expr COMMA expr COMMA expr RPAREN { $$ = newQuestion($3, $5, $7, yylex);}
    | expr MUL expr { $$ = newBinary($1, $3, MUL, yylex); }
    | expr DIV expr { $$ = newBinary($1, $3, DIV, yylex);  }
    | expr ADD expr { $$ = newBinary($1, $3, ADD, yylex); }
    | expr SUB expr { $$ = newBinary($1, $3, SUB, yylex);}
    | expr MOD expr { $$ = newBinary($1, $3, MOD, yylex); } 
    | expr AND expr { $$ = newBinary($1, $3, AND, yylex); }
    | expr OR expr { $$ = newBinary($1, $3, OR, yylex);  }
    | expr EQ expr { $$ = newBinary($1, $3, EQ, yylex); }
    | expr NOT_EQ expr { $$ = newBinary($1, $3, NOT_EQ, yylex);}
    | expr LTE expr { $$ = newBinary($1, $3, LTE, yylex); }
    | expr GTE expr { $$ = newBinary($1, $3, GTE, yylex);  }
    | expr LT expr { $$ = newBinary($1, $3, LT, yylex); }
    | expr GT expr { $$ = newBinary($1, $3, GT, yylex);}

    | SUB expr %prec UNARYMINUS { $$ = newUnary($2, SUB, yylex)}
    | NOT expr %prec UNARYNOT { $$ = newUnary($2, NOT, yylex)}
    ;

ident_list
    : IDENT { $$ = []string{$1} }
    | ident_list IDENT { $$ = append($1, $2) }
    ;

par_declaration
    : type ident_list { $$ = newVars($1, $2)}
    ;

par_declarations
    : /*empty*/ {$$=nil}
    | par_declaration { $$ = $1}
    | par_declarations COMMA par_declaration { $$ = append($1, $3...) }
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
        $$ = newBlock($1, $2, yylex)
    }
    ;

contract_declaration 
    : CONTRACT IDENT LBRACE contract_body RBRACE { 
        $$ = newContract($2, $4, yylex)
        setResult(yylex, $$)
        }
    | contract_declaration NEWLINE
    ;  
