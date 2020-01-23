%{
package main

import (
        "fmt"
        "strconv"
)
%}

%union{
String string
Number float64
}


%token<String> NUMBER IDENTIFIER

%type <Number> expr

%%
start: expr {fmt.Println($1)}
     | assignment;

expr:
      NUMBER { 
        var err error
        $$, err = strconv.ParseFloat($1, 64)
        if err != nil{
                yylex.Error(err.Error())
        }
        }
    | IDENTIFIER {
        var ok bool
        $$, ok = yylex.(*interpreter).vars[$1]
        if !ok {
                yylex.Error(fmt.Sprintf("Variable undefined: %s\n", $1))
        }
        }
    | expr '+' expr { $$ = $1 + $3 }
    | expr '-' expr { $$ = $1 - $3 }
    | expr '*' expr { $$ = $1 * $3 }
    | expr '/' expr { $$ = $1 / $3 }
    | '(' expr ')'  { $$ = $2 }
    | '-' expr { $$ = -$2 }
    ;

assignment:
          IDENTIFIER '=' expr {
                if !yylex.(*interpreter).evaluationFailed {
                        yylex.(*interpreter).vars[$1] = $3 
                }};
%%
