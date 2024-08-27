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
	Segment Token
	Index   Token
}

func (s StackInstruction) Literal() string {
	return fmt.Sprintf("%s %s %s", s.Command.Literal, s.Segment.Literal, s.Index.Literal)
}

type ArithmeticInstruction interface {
	// ただのinterfaceつくる用
	GetCommand() Token
}

type UnaryArithmeticInstruction struct {
	Command Token
}

func (ins UnaryArithmeticInstruction) Literal() string {
	return fmt.Sprintf("%s", ins.Command.Literal)
}
func (ins UnaryArithmeticInstruction) GetCommand() Token {
	return ins.Command
}

type BinaryArithmeticInstruction struct {
	Command Token
}

func (ins BinaryArithmeticInstruction) Literal() string {
	return fmt.Sprintf("%s", ins.Command.Literal)
}

func (ins BinaryArithmeticInstruction) GetCommand() Token {
	return ins.Command
}
