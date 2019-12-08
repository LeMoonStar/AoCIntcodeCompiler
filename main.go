package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	F_DBG                  bool
	Commands               instructionset
	Labels                 map[string]int
	BuildinInstructionSets map[string]instructionset = map[string]instructionset{
		"standard": instructionset{
			"add": command{
				Instruction: 1,
				ArgCount:    3,
			},
			"multi": command{
				Instruction: 2,
				ArgCount:    3,
			},
			"input": command{
				Instruction: 3,
				ArgCount:    1,
			},
			"output": command{
				Instruction: 4,
				ArgCount:    1,
			},
			"jmpit": command{
				Instruction: 5,
				ArgCount:    2,
			},
			"jmpif": command{
				Instruction: 6,
				ArgCount:    2,
			},
			"less": command{
				Instruction: 7,
				ArgCount:    3,
			},
			"equal": command{
				Instruction: 8,
				ArgCount:    3,
			},
		},
	}
)

type command struct {
	Instruction int
	ArgCount    int
}

type instructionset map[string]command

func main() {
	InputFileName := ""
	OutputFileName := "out.int"
	_ = OutputFileName
	Args := os.Args
	for i := 1; i < len(Args); i++ {
		v := Args[i]
		if v[0] == '-' {
			flag := v[1:]
			switch flag {
			case "dbg":
				F_DBG = true
			case "o":
				i++
				if len(Args) > i {
					OutputFileName = Args[i]
				} else {
					fmt.Printf("-o needs a filename.\n")
					os.Exit(1)
				}

			default:
				CloseMessage(1, "Flag: \"%s\" is not reconized.", flag)
			}
		} else {
			InputFileName = v
		}
	}
	if InputFileName == "" {
		CloseMessage(1, "Please enter a Input file.")
	}

	f, err := os.Open(InputFileName)
	if err != nil {
		CloseMessage(1, "Could not open input file \"%s\"", InputFileName)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	lines := make([][]string, 0, 0)
	WordPattern := regexp.MustCompile("\\S+")
	for scanner.Scan() {
		tmp := make([]string, 0)
		for _, v := range WordPattern.FindAllString(scanner.Text(), -1) {
			tmp = append(tmp, v)
		}
		lines = append(lines, tmp)
	}
	f.Close()

	loadInstructionset("standard")
	Output := make([]*int, 0)
	for LineID, _ := range lines {
		Output = append(Output, CompileCommand(lines[LineID], Output, LineID)...)
	}

}

func CompileCommand(args []string, output []*int, line int) []*int {
	out := make([]*int, 0)
	for i := 0; i < len(args); i++ {
		Type := GetWordType(args[i])
		fmt.Print(args[i], ":", Type, ", ")
		switch Type {
		case T_TEXT:
			if _, ok := Commands[args[i]]; ok {
				//cmdID := i
				cmdArgs := make([]*int, 0)
				ArgMode := make([]bool, Commands[args[i]].ArgCount) //true = value, false = pointer/reference
				if len(args[i+1:]) < Commands[args[i]].ArgCount {
					CloseMessage(1, "Line %d, Word %d: ERROR: command %s needs %d arguments.", line, i, args[i], Commands[args[i]].ArgCount)
				}
				for i2 := 0; i2 < Commands[args[i]].ArgCount; i2++ {
					fmt.Println(Commands[args[i]].ArgCount)
					i++
					Type := GetWordType(args[i])
					fmt.Print(args[i], ";:", Type, ", ")
					switch Type {
					case T_POINTER:
						DebugPrint("Line %d, Word %d: DEBUG: found Pointer", line, i)
						ArgMode[i2] = false
						addr, err := strconv.Atoi(args[i][1:])
						if err != nil {
							CloseMessage(1, "Line %d, Word %d: ERROR: Value of pointer needs to be a number.", line, i)
						}
						cmdArgs = append(cmdArgs, GetSharedIntPointer(addr))
					case T_LITTERAL:
						DebugPrint("Line %d, Word %d: DEBUG: found Litteral", line, i)
						ArgMode[i2] = true
						val, err := strconv.Atoi(args[i][1:])
						if err != nil {
							CloseMessage(1, "Line %d, Word %d: ERROR: Value of litteral needs to be a number.", line, i)
						}
						cmdArgs = append(cmdArgs, GetSharedIntPointer(val))

					}
				}
				out = append(out, GetSharedIntPointer(int(Commands[args[i]].Instruction)))
			} else {
				//sCloseMessage(1, "Command \"%s\" not found.", args[i])
			}
		}

	}
	fmt.Println()
	return out
}

func loadInstructionset(name string) {
	Commands = BuildinInstructionSets[name]
}
