package main

import (
	"fmt"
	"hack-vm/vm"
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	File string
}

func main() {
	arg.MustParse(&args)
	if args.File == "" {
		fmt.Println("Must specify a file to open\nUsage:\n\tgo run main.go --file=<hoge.vm>")
		os.Exit(1)
	}

	content, err := os.ReadFile(args.File)
	if err != nil {
		panic(err)
	}
	lexer := vm.NewLexer(string(content))
	parser := vm.NewParser(lexer)
	program := parser.Parse()
	writer := vm.NewCodeWriter()

	// TODO: ディレクトリ対応
	writer.SetFileName(args.File)
	for _, ins := range program.Instructions {
		switch ins.(type) {
		case vm.StackInstruction:
			ins := ins.(vm.StackInstruction)
			writer.WritePushPop(ins)
		case vm.ArithmeticInstruction:
			ins := ins.(vm.ArithmeticInstruction)
			writer.WriteArithmetic(ins)
		}
	}
	writer.Close()
	// os.WriteFile(fmt.Sprintf("%s.intermediate", args.File), []byte(output), os.ModePerm)
}
