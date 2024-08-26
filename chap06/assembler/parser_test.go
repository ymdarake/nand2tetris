package assembler

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func Test_parserImpl_HasMoreCommands(t *testing.T) {
	type fields struct {
		input     bufio.Reader
		readCount int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "1",
			fields: fields{
				input:     *bufio.NewReader(strings.NewReader("@i\nD=M\n@100\n")),
				readCount: 1,
			},
			want: true,
		},
		{
			name: "2",
			fields: fields{
				input:     *bufio.NewReader(strings.NewReader("@i\nD=M\n@100\n")),
				readCount: 2,
			},
			want: true,
		},
		{
			name: "3",
			fields: fields{
				input:     *bufio.NewReader(strings.NewReader("@i\nD=M\n@100\n")),
				readCount: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				input: tt.fields.input,
			}
			for range tt.fields.readCount {
				if err := p.Advance(); err != nil {
					t.Errorf("HasMoreCommands test: ERROR: parserImpl.Advance() = %v,", err)
				}
			}
			if got := p.HasMoreCommands(); got != tt.want {
				t.Errorf("parserImpl.HasMoreCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserImpl_Advance(t *testing.T) {
	type fields struct {
		input             bufio.Reader
		currentLine       string
		currentLineNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				input:             tt.fields.input,
				currentLine:       tt.fields.currentLine,
				currentLineNumber: tt.fields.currentLineNumber,
			}
			if err := p.Advance(); (err != nil) != tt.wantErr {
				t.Errorf("parserImpl.Advance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parserImpl_CurrentCommandType(t *testing.T) {
	type fields struct {
		currentLine string
	}
	tests := []struct {
		name   string
		fields fields
		want   commandType
	}{
		{
			name: "1: L",
			fields: fields{
				currentLine: "(LOOP)",
			},
			want: L_COMMAND,
		},
		{
			name: "2: A: @ prefixed number",
			fields: fields{
				currentLine: "@1234",
			},
			want: A_COMMAND,
		},
		{
			name: "3: C: dest=comp;jump",
			fields: fields{
				currentLine: "dest=comp;jump",
			},
			want: C_COMMAND,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				currentLine: tt.fields.currentLine,
			}
			if got := p.CurrentCommandType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parserImpl.CurrentCommandType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserImpl_Symbol(t *testing.T) {
	type fields struct {
		currentLine string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				currentLine: "(LOOP)",
			},
			want: "LOOP",
		},
		{
			name: "2",
			fields: fields{
				currentLine: "(Label)",
			},
			want: "Label",
		},
		{
			name: "3: @ prefixed number",
			fields: fields{
				currentLine: "@1234",
			},
			want: "1234",
		},
		{
			name: "4: @ prefixed id",
			fields: fields{
				currentLine: "@i",
			},
			want: "i",
		},
		{
			name: "5: @ prefixed id",
			fields: fields{
				currentLine: "@sum",
			},
			want: "sum",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				currentLine: tt.fields.currentLine,
			}
			if got := p.Symbol(); got != tt.want {
				t.Errorf("parserImpl.Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserImpl_Dest(t *testing.T) {
	type fields struct {
		currentLine string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				currentLine: "dest=comp;jump",
			},
			want: "dest",
		},
		{
			name: "2: no jump",
			fields: fields{
				currentLine: "M=0",
			},
			want: "M",
		},
		{
			name: "3: no jump, minus",
			fields: fields{
				currentLine: "D=D-A",
			},
			want: "D",
		},
		{
			name: "4: none(just jump)",
			fields: fields{
				currentLine: "0;JMP",
			},
			want: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				currentLine: tt.fields.currentLine,
			}
			if got := p.Dest(); got != tt.want {
				t.Errorf("parserImpl.Dest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserImpl_Comp(t *testing.T) {
	type fields struct {
		currentLine string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				currentLine: "dest=comp;jump",
			},
			want: "comp",
		},
		{
			name: "2: no jump",
			fields: fields{
				currentLine: "M=0",
			},
			want: "0",
		},
		{
			name: "3: no jump, minus",
			fields: fields{
				currentLine: "D=D-A",
			},
			want: "D-A",
		},
		{
			name: "4: none(just jump)",
			fields: fields{
				currentLine: "0;JMP",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				currentLine: tt.fields.currentLine,
			}
			if got := p.Comp(); got != tt.want {
				t.Errorf("parserImpl.Comp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parserImpl_Jump(t *testing.T) {
	type fields struct {
		input             bufio.Reader
		currentLine       string
		currentLineNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1",
			fields: fields{
				currentLine: "dest=comp;jump",
			},
			want: "jump",
		},
		{
			name: "2: no jump",
			fields: fields{
				currentLine: "M=0",
			},
			want: "null",
		},
		{
			name: "3: no jump, minus",
			fields: fields{
				currentLine: "D=D-A",
			},
			want: "null",
		},
		{
			name: "4: none(just jump)",
			fields: fields{
				currentLine: "0;JMP",
			},
			want: "JMP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parserImpl{
				input:             tt.fields.input,
				currentLine:       tt.fields.currentLine,
				currentLineNumber: tt.fields.currentLineNumber,
			}
			if got := p.Jump(); got != tt.want {
				t.Errorf("parserImpl.Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
