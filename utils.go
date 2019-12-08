package main

import (
	"fmt"
	"os"
	"strconv"
)

var FreeInts []int

const (
	T_UNKNOWN  int8 = -1
	T_LITTERAL int8 = 0
	T_POINTER  int8 = 1
	T_LABEL    int8 = 2
	T_VARIABLE int8 = 3
	T_TEXT     int8 = 4
)

func GetSharedIntPointer(a int) *int {
	for k, _ := range FreeInts {
		if FreeInts[k] == a {
			return &FreeInts[k]
		}
	}
	FreeInts = append(FreeInts, a)
	return &FreeInts[len(FreeInts)-1]
}

func DebugPrint(formatter string, args ...interface{}) {
	fmt.Printf(formatter+"\n", args...)
}

func CloseMessage(code int, formatter string, args ...interface{}) {
	fmt.Printf(formatter+"\n", args...)
	os.Exit(code)
}

func GetWordType(Word string) int8 {
	switch Word[0] {
	case ':':
		return T_LABEL
	case '*':
		return T_POINTER
	case '_':
		return T_VARIABLE
	default:
		if _, err := strconv.Atoi(Word); err != nil {
			return T_TEXT
		} else {
			return T_LITTERAL
		}
	}
}
