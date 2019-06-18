package evaluator

import (
	"bufio"
	"os"

	"github.com/Code-Hex/shibafu/lexer"
	"github.com/Code-Hex/shibafu/token"
)

const stackSize = 65535

type Evaluator struct {
	lexer   *lexer.Lexer
	reader  *bufio.Reader
	program []token.Type
}

const (
	INC = iota
	DEC
	INCVAL
	DECVAL
	OUT
	IN
	JMP
	BCK
)

func New(input string) *Evaluator {
	return &Evaluator{
		lexer:  lexer.New(input),
		reader: bufio.NewReader(os.Stdin),
	}
}

func (e *Evaluator) Evaluate() {
	var ptr int
	data := make([]byte, stackSize)
	for pc := 0; pc < len(e.program); pc++ {
		switch t.Type {
		case token.EOF:
			break
		case token.INCR:
			ptr++
		case token.DECR:
			ptr--
		case token.NEXT:
			data[ptr]++
		case token.PREV:
			data[ptr]--
		case token.READ:
			rv, _ := e.reader.ReadByte()
			data[ptr] = rv
		case token.WRITE:
			os.Stdout.Write([]byte{data[ptr]})
		case token.OPEN:

		}
	}

}

func (e *Evaluator) compile() {
	for {
		t := e.lexer.NextToken()
		e.program = append(e.program, t.Type)
	}
}

func (e *Evaluator) execute() {
	
}