package evaluator

import (
	"bufio"
	"errors"
	"os"

	"github.com/Code-Hex/shibafu/lexer"
	"github.com/Code-Hex/shibafu/token"
)

const stackSize = 65535

type Evaluator struct {
	lexer   *lexer.Lexer
	reader  *bufio.Reader
	program []*instruction
}

type instruction struct {
	operator int
	pc       int
}

const (
	INCR = iota
	DECR
	INCRVAL
	DECRVAL
	WRITE
	READ
	JMP
	BCK
)

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
				operator: INCR,
				pc:       pc,
			})
		case token.DECR:
			e.program = append(e.program, &instruction{
				operator: DECR,
				pc:       pc,
			})
		case token.NEXT:
			e.program = append(e.program, &instruction{
				operator: INCRVAL,
				pc:       pc,
			})
		case token.PREV:
			e.program = append(e.program, &instruction{
				operator: DECRVAL,
				pc:       pc,
			})
		case token.READ:
			e.program = append(e.program, &instruction{
				operator: READ,
				pc:       pc,
			})
		case token.WRITE:
			e.program = append(e.program, &instruction{
				operator: WRITE,
				pc:       pc,
			})
		case token.OPEN:
			e.program = append(e.program, &instruction{
				operator: JMP,
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
				operator: BCK,
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
		case INCR:
			ptr++
		case DECR:
			ptr--
		case INCRVAL:
			data[ptr]++
		case DECRVAL:
			data[ptr]--
		case WRITE:
			os.Stdout.Write([]byte{data[ptr]})
		case READ:
			rv, _ := e.reader.ReadByte()
			data[ptr] = rv
		case JMP:
			if data[ptr] == 0 {
				pc = e.program[pc].pc
			}
		case BCK:
			if data[ptr] != 0 {
				pc = e.program[pc].pc
			}
		default:
			return errors.New("runtime error")
		}
	}
	return nil
}
