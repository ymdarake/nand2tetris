package vm

import (
	"reflect"
	"testing"
)

func TestNewLexer(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Lexer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLexer(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLexer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexerImpl_NextToken(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Token
	}{
		{
			name: "1",
			fields: fields{
				input: `push constant 7
		push constant 8
		add
		`,
			},
			want: []Token{
				{Type: CMD_PUSH, Literal: "push"},
				{Type: IDENT_CONSTANT, Literal: "constant"},
				{Type: INT, Literal: "7"},
				{Type: CMD_PUSH, Literal: "push"},
				{Type: IDENT_CONSTANT, Literal: "constant"},
				{Type: INT, Literal: "8"},
				{Type: CMD_ADD, Literal: "add"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.fields.input)
			got := make([]Token, 0)

			for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
				got = append(got, tok)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexerImpl.NextToken() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}

func Test_lexerImpl_Peek(t *testing.T) {
	type fields struct {
		input           string
		currentPosition int
		ch              byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			if got := l.Peek(); got != tt.want {
				t.Errorf("lexerImpl.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexerImpl_Advance(t *testing.T) {
	type fields struct {
		input           string
		currentPosition int
		ch              byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			l.Advance()
		})
	}
}

func Test_lexerImpl_skipWhitespace(t *testing.T) {
	type fields struct {
		input           string
		currentPosition int
		ch              byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			l.skipWhitespace()
		})
	}
}

func Test_lexerImpl_readIdentifier(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   struct {
			token      Token
			readCursor int
		}
	}{
		{
			name:   "1",
			fields: fields{input: `push constant 7`},
			want: struct {
				token      Token
				readCursor int
			}{token: Token{Type: CMD_PUSH, Literal: "push"}, readCursor: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			l.Advance()
			if got := l.readIdentifier(); !reflect.DeepEqual(got, tt.want.token) || (l.readCursor != tt.want.readCursor) {
				t.Errorf("lexerImpl.readIdentifier() = %v, readCursor %d, want %v", got, l.readCursor, tt.want)
			}
		})
	}
}

func Test_lexerImpl_readNumber(t *testing.T) {
	type fields struct {
		input           string
		currentPosition int
		ch              byte
	}
	tests := []struct {
		name   string
		fields fields
		want   Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			if got := l.readNumber(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexerImpl.readNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexerImpl_isLetter(t *testing.T) {
	type fields struct {
		input           string
		currentPosition int
		ch              byte
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			if got := l.isLetter(tt.args.ch); got != tt.want {
				t.Errorf("lexerImpl.isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexerImpl_isDigit(t *testing.T) {
	type fields struct {
		input           string
		currentPosition int
		ch              byte
	}
	type args struct {
		ch byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lexerImpl{input: tt.fields.input}
			if got := l.isDigit(tt.args.ch); got != tt.want {
				t.Errorf("lexerImpl.isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
