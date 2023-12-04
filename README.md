# Go Math Expressions

## Overview

Basic math expressions parser and evaluator, written completely in go.

Operator supports unary operators (**-** minus operator only) as well as precedence parsing and brackets.

## Supported functions

### Infix operators

- \+ plus operator
- \- minus operator
- \* multiply operator
- / divide operator

### Prefix operatorators

- min() accepts list of any number of arguments, returns one that resulted in minimum value

- max() accepts list of any number of arguments, returns one that resulted in maximum value

- abs() accepts one argument and returns its absolute value

## Example usage

`go run . "5 + 4 * 3 + abs(min(3, -10, 56))"`

