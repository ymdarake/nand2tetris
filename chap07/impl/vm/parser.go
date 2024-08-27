package vm

type Parser interface {
	Parse() *Program
	// HasMoreCommand() bool
	// Advance() // Next Command
	// CommandType() *CommandType
	// Arg1() string
	// Arg2() int
}

type CommandType int

const (
	C_ARITHMETIC = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

var COMMAND_TYPE_MAP = map[TokenType]CommandType{
	CMD_PUSH: C_PUSH,
	CMD_POP:  C_POP,
	CMD_ADD:  C_ARITHMETIC,
	CMD_SUB:  C_ARITHMETIC,
	CMD_NEG:  C_ARITHMETIC,
	CMD_EQ:   C_ARITHMETIC,
	CMD_GT:   C_ARITHMETIC,
	CMD_LT:   C_ARITHMETIC,
	CMD_AND:  C_ARITHMETIC,
	CMD_OR:   C_ARITHMETIC,
	CMD_NOT:  C_ARITHMETIC,
}

func LookupCommandType(tokenType TokenType) *CommandType {
	ct, ok := COMMAND_TYPE_MAP[tokenType]
	if !ok {
		return nil
	}
	return &ct
}

type parserImpl struct {
	lexer        Lexer
	currentToken Token
}

// Parse implements Parser.
func (p *parserImpl) Parse() *Program {
	program := &Program{}

	for p.currentToken.Type != EOF {
		instruction := p.parseInstruction()
		if instruction != nil {
			program.Instructions = append(program.Instructions, instruction)
		}
		p.nextToken()
	}

	return program
}

func (p *parserImpl) nextToken() {
	p.currentToken = p.lexer.NextToken()
}

func (p *parserImpl) parseInstruction() Instruction {
	switch p.currentToken.Type {
	case CMD_PUSH, CMD_POP:
		return p.parseStackInstruction()
	case CMD_ADD, CMD_SUB, CMD_EQ, CMD_LT, CMD_GT, CMD_AND, CMD_OR:
		return p.parseBinaryArithmeticInstruction()
	case CMD_NEG, CMD_NOT:
		return p.parseUnaryArithmeticInstruction()
	case LINE_COMMENT:
		// fmt.Println("skip parsing comment")
	}

	return nil
}

func (p *parserImpl) parseStackInstruction() StackInstruction {
	instruction := StackInstruction{}
	instruction.Command = p.currentToken
	p.nextToken()
	instruction.Arg1 = p.currentToken
	p.nextToken()
	instruction.Arg2 = p.currentToken
	return instruction
}

func (p *parserImpl) parseBinaryArithmeticInstruction() BinaryArithmeticInstruction {
	return BinaryArithmeticInstruction{Command: p.currentToken}
}
func (p *parserImpl) parseUnaryArithmeticInstruction() UnaryArithmeticInstruction {
	return UnaryArithmeticInstruction{Command: p.currentToken}
}

func NewParser(lexer Lexer) Parser {
	return &parserImpl{lexer: lexer}
}
