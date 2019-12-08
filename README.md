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

### Argument types
every argument can be one of the following types:
| Name      | Prefix    | Description                                                                                                                          |
| :---      | :---      | :---                                                                                                                                 |
| Littertal | NO PREFIX | represents the mode introduced on day 5, which is not a reference to another address, but represents a value.                        |
| Pointer   | *         | represents the first mode introduced on day 2, which is a reference to another address.                                              |
| Label     | :         | defines a name wich can be used anywhere in the code wich will be replaced by a Pointer to the value following the label Definition. |
| Variable  | _         | can be used just like a Pointer, but the address is the specified address plus the address of the last END statement.                |
| Reference | NO PREFIX | a reference will be replaced by the Address of the label wich name is used as argument.                                              |
