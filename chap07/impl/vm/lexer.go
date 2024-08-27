package vm

import (
	"unicode"
)

type Lexer interface {
	Advance()
	Peek() byte
	NextToken() Token
}

type lexerImpl struct {
	input            string
	lastCharPosition int
	readCursor       int
	ch               byte
}

func NewLexer(input string) Lexer {
	l := &lexerImpl{input: input}
	l.Advance()
	return l
}

// NextToken implements Lexer.
func (l *lexerImpl) NextToken() Token {
	l.skipWhitespace()

	if l.ch == '/' && l.Peek() == '/' {
		return l.readLineComment()
	} else if l.isLetter(l.ch) {
		return l.readIdentifier()
	} else if l.isDigit(l.ch) {
		return l.readNumber()
	} else if l.ch == 0 {
		return Token{Type: EOF, Literal: ""}
	} else {
		return Token{Type: ILLEGAL, Literal: string(l.ch)}
	}
}

// Peek implements Lexer.
func (l *lexerImpl) Peek() byte {
	if l.readCursor >= len(l.input) {
		return 0
	}
	return l.input[l.readCursor]
}

// Advance implements Lexer.
func (l *lexerImpl) Advance() {
	if l.readCursor >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readCursor]
	}
	l.lastCharPosition = l.readCursor
	l.readCursor++
}

func (l *lexerImpl) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.Advance()
	}
}

func (l *lexerImpl) readLineComment() Token {
	startPos := l.lastCharPosition
	for l.ch != '\n' {
		l.Advance()
	}
	return Token{Type: LINE_COMMENT, Literal: l.input[startPos:l.lastCharPosition]}
}
func (l *lexerImpl) readIdentifier() Token {
	startPos := l.lastCharPosition
	for l.isLetter(l.ch) {
		l.Advance()
	}
	ident := l.input[startPos:l.lastCharPosition]
	return TokenFromIdentifier(ident)
}

func (l *lexerImpl) readNumber() Token {
	startPos := l.lastCharPosition
	for l.isDigit(l.ch) {
		l.Advance()
	}
	str := l.input[startPos:l.lastCharPosition]
	return Token{Type: INT, Literal: str}
}

func (l *lexerImpl) isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func (l *lexerImpl) isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}
