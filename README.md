# goyacc-tutorial

This is a simple toy calculator written in golang with yacc.

It has been written for a talk at GoDays Berlin 2020, a recording of which can be found [here](https://youtu.be/N1kOV4biSRw).

The directory `calculator` contains a version of the parser which uses direct evaluation. The multiple iteration of the grammar are in files named `grammar_<iteration>_<description>.y`.


The directory `ast_calculator` contains a version that uses an abstract syntax tree.

## Building and running

* Install goyacc `go get golang.org/x/tools/cmd/goyacc`
* Go to the directory with the desired parser version.
* Run `goyacc <desired grammar file>.y`.
* Run `go build`.
* Run the created calculator binary.
