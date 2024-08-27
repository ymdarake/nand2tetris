package vm

import (
	"fmt"
	"os"
	"strconv"
)

type CodeWriter interface {
	SetFileName(filename string)
	// NOTE: 7章ではひとまず算術命令、スタック命令のみを実装する
	WriteArithmetic(ins ArithmeticInstruction)
	WritePushPop(ins StackInstruction)
	Close() error // flush的な.
}

type codeWriterImpl struct {
	currentFileName string
	code            string
}

func NewCodeWriter() CodeWriter {
	return &codeWriterImpl{}
}

// SetFileName implements CodeWriter.
func (c *codeWriterImpl) SetFileName(filename string) {
	c.currentFileName = filename
}

// WriteArithmetic implements CodeWriter.
func (c *codeWriterImpl) WriteArithmetic(ins ArithmeticInstruction) {
	switch ins.GetCommand().Type {
	case CMD_ADD, CMD_SUB, CMD_EQ, CMD_LT, CMD_GT, CMD_AND, CMD_OR:
		c.writeBinaryArithmeticInstruction(ins)
		return
	case CMD_NEG, CMD_NOT:
		// c.writeUnaryArithmeticInstruction(ins)
		return
	}
}

// WritePushPop implements CodeWriter.
func (c *codeWriterImpl) WritePushPop(ins StackInstruction) {
	switch ins.Command.Type {
	case CMD_PUSH:
		c.writePush(ins)
	case CMD_POP:
		c.writePop(ins)
	default:
		fmt.Println("default")
	}
}

// Close implements CodeWriter.
func (c *codeWriterImpl) Close() error {
	if err := os.WriteFile(fmt.Sprintf("%s.gen.asm", c.currentFileName), []byte(c.code), os.ModePerm); err != nil {
		return err
	}
	c.currentFileName = ""
	c.code = ""
	return nil
}

func (c *codeWriterImpl) writePush(ins StackInstruction) error {
	switch ins.Segment.Literal {
	case "constant":
		c.code += fmt.Sprintf("@%s\n", ins.Index.Literal) // Aレジスタに値を入れる
		c.code += "D=A\n"                                 // DレジスタにAレジスタの値を入れる(定数値)
		c.pushFromDRegister()
		return nil
	case "local", "argument", "this", "that":
		// セグメントのindexにある値をDレジスタに入れる
		c.code += fmt.Sprintf("@%s\n", SegmentLiteralToRegister(ins.Segment.Literal))
		index, err := strconv.Atoi(ins.Index.Literal)
		if err != nil {
			return err
		}
		for i := 0; i < index; i++ {
			c.code += fmt.Sprintf("A=A+1\n")
		}
		c.code += "D=M\n"
		c.pushFromDRegister()
		return nil
	case "pointer":
	case "temp":
	case "static":
	}

	return nil
}

func (c *codeWriterImpl) pushFromDRegister() {
	// Dレジスタにある値をスタックに積む
	c.code += "@SP\n" // MレジスタにSPのアドレスを入れる
	c.code += "A=M\n" // AレジスタにMレジスタの値(SPのアドレス)を入れる
	c.code += "M=D\n" // MレジスタにDレジスタの値を入れる(定数値)
	// スタックポインタをひとつ進める
	c.code += "@SP\n"   // MレジスタにSPのアドレスを入れる
	c.code += "M=M+1\n" // SPのアドレスをひとつ進める
}
func (c *codeWriterImpl) writePop(ins StackInstruction) error {
	switch ins.Segment.Literal {
	case "local", "argument", "this", "that":
		// pushと逆順
		// スタックポインタをひとつ戻す
		c.code += "@SP\n"
		c.code += "M=M-1\n"
		c.code += "A=M\n"

		// ひとつ戻ったスタックポインタの指す値をDレジスタに入れる.
		c.code += "D=M\n"

		// Dレジスタの値をセグメントのindexに入れる
		c.code += fmt.Sprintf("@%s\n", SegmentLiteralToRegister(ins.Segment.Literal))
		c.code += "A=M\n"
		index, err := strconv.Atoi(ins.Index.Literal)
		if err != nil {
			return err
		}
		for i := 0; i < index; i++ {
			c.code += fmt.Sprintf("A=A+1\n")
		}
		c.code += "M=D\n"
		return nil
	case "pointer":
	case "temp":
	case "static":
	}

	return nil
}

func (c *codeWriterImpl) writeBinaryArithmeticInstruction(ins ArithmeticInstruction) error {
	// ひとつ目の引数をDレジスタに持ってくる
	c.code += "@SP\n"
	c.code += "M=M-1\n"
	c.code += "D=M\n"

	// ふたつ目の引数をAレジスタに持ってくる
	c.code += "@SP\n"
	c.code += "M=M-1\n"
	c.code += "A=M\n"

	switch ins.GetCommand().Type {
	case CMD_ADD:
		c.code += "D=D+M\n"
	case CMD_SUB:
		c.code += "D=M-D\n" // popしてるし、逆順なのはある種直感的.
	case CMD_AND:
		c.code += "D=D&M\n"
	case CMD_OR:
		c.code += "D=D|M\n"
	case CMD_EQ, CMD_LT, CMD_GT:
	}
	c.pushFromDRegister()
	return nil
}

func SegmentLiteralToRegister(segment string) string {
	mp := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
	}
	return mp[segment]
}
