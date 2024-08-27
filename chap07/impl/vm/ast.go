package vm

import "fmt"

type Program struct {
	Instructions []Instruction
}

// interface 用
type Instruction interface {
	Literal() string
}

type StackInstruction struct {
	Command Token
	Arg1    Token
	Arg2    Token
}

func (s StackInstruction) Literal() string {
	return fmt.Sprintf("%s %s %s", s.Command.Literal, s.Arg1.Literal, s.Arg2.Literal)
}

type UnaryArithmeticInstruction struct {
	Command Token
}

func (ins UnaryArithmeticInstruction) Literal() string {
	return fmt.Sprintf("%s", ins.Command.Literal)
}

type BinaryArithmeticInstruction struct {
	Command Token
}

func (ins BinaryArithmeticInstruction) Literal() string {
	return fmt.Sprintf("%s", ins.Command.Literal)
}
