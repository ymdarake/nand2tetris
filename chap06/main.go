package main

import (
	"fmt"
	"hack-assembler/assembler"
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	File string
}

func main() {
	arg.MustParse(&args)
	if args.File == "" {
		fmt.Println("Must specify a file to open\nUsage:\n\tgo run main.go --file=<path without ext>")
		os.Exit(1)
	}

	file, err := os.Open(fmt.Sprintf("%s.asm", args.File))
	if err != nil {
		panic(err)
	}

	// output := assembler.AssembleWithoutLabel(file)
	output := assembler.Assemble(file)
	os.WriteFile(fmt.Sprintf("%s.hack", args.File), []byte(output), os.ModePerm)
}
