package assembler

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type commandType int

const (
	A_COMMAND commandType = iota
	C_COMMAND
	L_COMMAND
	COMMENT
)

func (ct commandType) String() string {
	switch ct {
	case A_COMMAND:
		return "A_COMMAND"
	case C_COMMAND:
		return "C_COMMAND"
	case L_COMMAND:
		return "L_COMMAND"
	}

	return "unknown"
}

type Parser interface {
	HasMoreCommands() bool
	Advance() error
	CurrentCommandType() commandType
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
	GetLineNumber() int
}

type parserImpl struct {
	input             bufio.Reader
	currentLine       string
	currentLineNumber int // 0始まり
}

func NewParser(in io.Reader) Parser {
	return &parserImpl{
		input: *bufio.NewReader(in),
	}
}

func (p *parserImpl) GetLineNumber() int {
	return p.currentLineNumber
}

func (p *parserImpl) skipWhitespace() {
	ch, err := p.input.Peek(1)
	if err != nil {
		return
	}
	r, _ := utf8.DecodeRune(ch)
	for unicode.IsSpace(r) {
		if r == '\n' {
			p.currentLineNumber++
		}
		p.input.Discard(1)
		ch, err := p.input.Peek(1)
		if err != nil {
			return
		}
		r, _ = utf8.DecodeRune(ch)
	}
}

func (p *parserImpl) HasMoreCommands() bool {
	_, err := p.input.Peek(1)
	return err != io.EOF
}
func (p *parserImpl) Advance() error {
	p.skipWhitespace()
	line, err := p.input.ReadString('\n')
	p.currentLineNumber++
	if err != nil {
		return fmt.Errorf("%w: line %d", err, p.currentLineNumber)
	}
	p.currentLine = strings.TrimSpace(line)
	return nil
}

// NOTE: エラー検知はしていない
func (p *parserImpl) CurrentCommandType() commandType {
	firstChar := p.currentLine[0]
	switch firstChar {
	case '(':
		return L_COMMAND
	case '@':
		return A_COMMAND
	case '/':
		return COMMENT
	default:
		return C_COMMAND
	}
}

// NOTE: ASCII only (1 byte chars)
func (p *parserImpl) Symbol() string {
	switch p.CurrentCommandType() {
	case L_COMMAND:
		return p.currentLine[1 : len(p.currentLine)-1]
	case A_COMMAND:
		return p.currentLine[1:]
	default:
		fmt.Println("ERROR: L/Aコマンド以外でSymbol()が呼び出されました.")
		return ""
	}
}

// dest=comp;jump
func (p *parserImpl) Dest() string {
	eqIndex := strings.Index(p.currentLine, "=")
	if eqIndex == -1 {
		return "null"
	}
	return p.currentLine[:eqIndex]
}

func (p *parserImpl) Comp() string {
	dest := p.Dest()
	startIndex := 0
	endIndex := len(p.currentLine)
	if dest != "null" {
		startIndex = len(dest) + len("=") // =も含む.
	}
	jump := p.Jump()
	if jump != "null" {
		endIndex -= (len(jump) + len(";"))
	}
	return p.currentLine[startIndex:endIndex]
}

func (p *parserImpl) Jump() string {
	semicolonIndex := strings.Index(p.currentLine, ";")
	if semicolonIndex == -1 {
		return "null"
	}
	return p.currentLine[semicolonIndex+1:]
}
