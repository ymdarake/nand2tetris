package main

import (
	"fmt"
	"hack-assembler/assembler"
	"io"
	"os"
	"strconv"

	"github.com/alexflint/go-arg"
)

var args struct {
	File string
}

func main() {
	arg.MustParse(&args)
	fmt.Println(args.File)
	if args.File == "" {
		fmt.Println("Must specify a file to open\nUsage:\n\tgo run main.go --file=<path without ext>")
		os.Exit(1)
	}

	file, err := os.Open(fmt.Sprintf("%s.asm", args.File))
	if err != nil {
		panic(err)
	}

	output := RunSimpleAssembler(file)
	os.WriteFile(fmt.Sprintf("%s.hack", args.File), []byte(output), os.ModePerm)
}

func RunSimpleAssembler(input io.Reader) string {
	output := ""
	parser := assembler.NewParser(input)
	for parser.HasMoreCommands() {
		if err := parser.Advance(); err != nil {
			panic(err)
		}
		fmt.Printf("parsing line %d\n", parser.GetLineNumber())
		switch parser.CurrentCommandType() {
		case assembler.A_COMMAND:
			symbol := parser.Symbol()
			val, err := strconv.Atoi(symbol)
			if err != nil {
				panic(err)
			}
			output += fmt.Sprintf("%016s\n", strconv.FormatInt(int64(val), 2))
		case assembler.C_COMMAND:
			// 111a c1c2c3c4 c5c6d1d2 d3j1j2j3
			// dest=comp;jump
			output += fmt.Sprintf("111%s%s%s\n", assembler.GenComp(parser.Comp()), assembler.GenDest(parser.Dest()), assembler.GenJump(parser.Jump()))
		case assembler.COMMENT:
			fmt.Println("skipping comment line")
		}
		// L_COMMANDはスキップ
	}

	return output
}
