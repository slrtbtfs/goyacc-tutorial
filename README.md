# goyacc-tutorial

This is a simple toy calculator written in golang with yacc.

It has been written for a talk ad GoDays.

The directory `calculator` contains a version of the parser which uses direct evaluation. The multiple iteration of the grammar are in files named `grammar_<iteration>_<description>.y`.


The directory `ast_calculator` contains a version that uses an abstract syntax tree.

## Building and running

* Install goyacc `go get golang.org/x/tools/cmd/goyacc`
* Go to the directory with the desired parser version.
* Run `goyacc <desired grammar file>.y`.
* Run `go build`.
* Run the created calculator binary.
