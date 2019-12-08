# AoC intcode Compiler
a assembly like language for the [AdventOfCode2019](https://adventofcode.com/2019/about) Intcode (Day [2](https://adventofcode.com/2019/day/2) and [5](https://adventofcode.com/2019/day/5))

## Language Difinition

### Examples
the following example is a program that takes 2 inputs, multiplies them and then outputs the result:
```
input mul1
input mul2
multi :mul1 0 :mul2 0 out
output :out 0
END
```
this program uses labels to get the addresses for the values of the multiplication and the output command, which are used as references in other commands.

the next example shows how to use labels in combination with jump commands:
```
input in
equal :in 0 5 flag
jmpit :flag jmpto
output 999
END
:jmpto output 0
END
```
this program takes an input and checks whether the input is equal to 5, if yes, it outputs 0, if not, it outputs 999.