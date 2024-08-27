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

	output := ""
	for _, v := range program.Instructions {
		output += fmt.Sprintf("%#v\n", v)
	}
	os.WriteFile(fmt.Sprintf("%s.intermediate", args.File), []byte(output), os.ModePerm)
}
