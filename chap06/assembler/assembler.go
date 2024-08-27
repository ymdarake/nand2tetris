package assembler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

func Assemble(input io.Reader) string {
	output := ""
	symbolTable := NewSymbolTable()
	nextVarAddress := 16

	var teedBuf bytes.Buffer
	tee := io.TeeReader(input, &teedBuf)
	labelParser := NewParser(tee)
	currentAddress := 0
	for labelParser.HasMoreCommands() {
		if err := labelParser.Advance(); err != nil {
			panic(err)
		}
		switch labelParser.CurrentCommandType() {
		case A_COMMAND, C_COMMAND:
			currentAddress++
		case L_COMMAND:
			symbolTable.AddEntry(labelParser.Symbol(), currentAddress)
		}
	}

	parser := NewParser(bufio.NewReader(&teedBuf))
	for parser.HasMoreCommands() {
		if err := parser.Advance(); err != nil {
			panic(err)
		}
		switch parser.CurrentCommandType() {
		case A_COMMAND:
			symbol := parser.Symbol()
			val, err := strconv.Atoi(symbol)
			if err != nil {
				if symbolTable.Contains(symbol) {
					// @LABEL
					output += fmt.Sprintf("%016s\n", strconv.FormatInt(int64(symbolTable.GetAddress(symbol)), 2))
					continue
				}
				// @var
				symbolTable.AddEntry(symbol, nextVarAddress)
				output += fmt.Sprintf("%016s\n", strconv.FormatInt(int64(nextVarAddress), 2))
				nextVarAddress++
				continue
			}
			// @1234
			output += fmt.Sprintf("%016s\n", strconv.FormatInt(int64(val), 2))
		case C_COMMAND:
			// 111a c1c2c3c4 c5c6d1d2 d3j1j2j3
			// dest=comp;jump
			output += fmt.Sprintf("111%s%s%s\n", GenComp(parser.Comp()), GenDest(parser.Dest()), GenJump(parser.Jump()))
		case L_COMMAND, COMMENT:
		}
	}

	return output
}

func AssembleWithoutLabel(input io.Reader) string {
	output := ""
	parser := NewParser(input)
	for parser.HasMoreCommands() {
		if err := parser.Advance(); err != nil {
			panic(err)
		}
		fmt.Printf("parsing line %d\n", parser.GetLineNumber())
		switch parser.CurrentCommandType() {
		case A_COMMAND:
			symbol := parser.Symbol()
			val, err := strconv.Atoi(symbol)
			if err != nil {
				panic(err)
			}
			output += fmt.Sprintf("%016s\n", strconv.FormatInt(int64(val), 2))
		case C_COMMAND:
			// 111a c1c2c3c4 c5c6d1d2 d3j1j2j3
			// dest=comp;jump
			output += fmt.Sprintf("111%s%s%s\n", GenComp(parser.Comp()), GenDest(parser.Dest()), GenJump(parser.Jump()))
		case COMMENT:
			fmt.Println("skipping comment line")
		}
		// L_COMMANDはスキップ
	}

	return output
}
