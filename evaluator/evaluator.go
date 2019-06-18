package evaluator

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/Code-Hex/shibafu/lexer"
	"github.com/Code-Hex/shibafu/token"
)

const (
	incr = iota
	decr
	incrval
	decrval
	write
	read
	fwd
	bck

	stackSize = 65535
)

var stdout io.Writer = os.Stdout

type Evaluator struct {
	lexer   lexer.Lexer
	reader  *bufio.Reader
	program []*instruction
}

type instruction struct {
	operator int
	pc       int
}

func New(input string) *Evaluator {
	return &Evaluator{
		lexer:  lexer.New(input),
		reader: bufio.NewReader(os.Stdin),
	}
}

func (e *Evaluator) Evaluate() error {
	if err := e.compile(); err != nil {
		return err
	}
	return e.execute()
}

func (e *Evaluator) compile() error {
	var jmpStack []int
LOOP:
	for pc := 0; ; pc++ {
		t := e.lexer.NextToken()
		switch t.Type {
		case token.EOF:
			break LOOP
		case token.INCR:
			e.program = append(e.program, &instruction{
				operator: incr,
				pc:       pc,
			})
		case token.DECR:
			e.program = append(e.program, &instruction{
				operator: decr,
				pc:       pc,
			})
		case token.NEXT:
			e.program = append(e.program, &instruction{
				operator: incrval,
				pc:       pc,
			})
		case token.PREV:
			e.program = append(e.program, &instruction{
				operator: decrval,
				pc:       pc,
			})
		case token.READ:
			e.program = append(e.program, &instruction{
				operator: read,
				pc:       pc,
			})
		case token.WRITE:
			e.program = append(e.program, &instruction{
				operator: write,
				pc:       pc,
			})
		case token.OPEN:
			e.program = append(e.program, &instruction{
				operator: fwd,
				pc:       pc,
			})
			jmpStack = append(jmpStack, pc)
		case token.CLOSE:
			if len(jmpStack) == 0 {
				return errors.New("compile error")
			}
			jmpLabel := jmpStack[len(jmpStack)-1]
			jmpStack = jmpStack[:len(jmpStack)-1]
			e.program = append(e.program, &instruction{
				operator: bck,
				pc:       jmpLabel,
			})
			e.program[jmpLabel].pc = pc
		case token.ILLEGAL:
			return errors.New("compile error")
		}
	}
	return nil
}

func (e *Evaluator) execute() error {
	var ptr int
	data := make([]byte, stackSize)
	for pc := 0; pc < len(e.program); pc++ {
		inst := e.program[pc]
		switch inst.operator {
		case incr:
			ptr++
		case decr:
			ptr--
		case incrval:
			data[ptr]++
		case decrval:
			data[ptr]--
		case write:
			stdout.Write([]byte{data[ptr]})
		case read:
			rv, _ := e.reader.ReadByte()
			data[ptr] = rv
		case fwd:
			if data[ptr] == 0 {
				pc = e.program[pc].pc
			}
		case bck:
			if data[ptr] != 0 {
				pc = e.program[pc].pc
			}
		default:
			return errors.New("runtime error")
		}
	}
	return nil
}
