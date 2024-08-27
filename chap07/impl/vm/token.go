package vm

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	CMD_PUSH TokenType = "push"
	CMD_POP  TokenType = "pop"

	CMD_ADD TokenType = "add"
	CMD_SUB TokenType = "sub"
	CMD_NEG TokenType = "neg"
	CMD_EQ  TokenType = "eq"
	CMD_GT  TokenType = "gt"
	CMD_LT  TokenType = "lt"
	CMD_AND TokenType = "and"
	CMD_OR  TokenType = "or"
	CMD_NOT TokenType = "not"

	IDENT_CONSTANT TokenType = "constant"
	IDENT          TokenType = "IDENT"
	INT            TokenType = "INT"
	LINE_COMMENT   TokenType = "LINE COMMENT"

	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)

var TOKEN_TYPE_MAP = map[string]TokenType{
	"push":     CMD_PUSH,
	"pop":      CMD_POP,
	"add":      CMD_ADD,
	"sub":      CMD_SUB,
	"neg":      CMD_NEG,
	"eq":       CMD_EQ,
	"gt":       CMD_GT,
	"lt":       CMD_LT,
	"and":      CMD_AND,
	"or":       CMD_OR,
	"not":      CMD_NOT,
	"constant": IDENT_CONSTANT,
}

func TokenFromIdentifier(ident string) Token {
	tokenType, ok := TOKEN_TYPE_MAP[ident]
	if !ok {
		return Token{Type: IDENT, Literal: ident}
	}
	return Token{Type: tokenType, Literal: ident}
}
