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


%%
     
start: expr {
        if !yylex.(*interpreter).evaluationFailed{
                fmt.Println($1)
        }};
        
expr:  NUMBER{ 
        var err error
        $$, err = strconv.ParseFloat($1, 64)
        if err != nil{
                yylex.Error(err.Error())
        }
        }
    ;
%%