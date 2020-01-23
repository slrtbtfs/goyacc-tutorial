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
%type<Number> expr

%left '+' '-'
%left '*' '/'

%%
     
start: expr {
        if !yylex.(*interpreter).evaluationFailed{
                fmt.Println($1)
        }};
        
expr: NUMBER{ 
        var err error
        $$, err = strconv.ParseFloat($1, 64)
        if err != nil{
                yylex.Error(err.Error())
        }
        }
    | expr '+' expr { $$ = $1 + $3 }
    | expr '-' expr { $$ = $1 - $3 }
    | expr '*' expr { $$ = $1 * $3 }
    | expr '/' expr { $$ = $1 / $3 }
    | '(' expr ')'  { $$ = $2 }
    | '-' expr %prec '*' { $$ = -$2 }
    ;
%%
