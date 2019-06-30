package evaluator

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"

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

	stackSize = 250000000
)

type Evaluator struct {
	lexer   lexer.Lexer
	reader  *bufio.Reader
	writer  io.Writer
	program []*instruction
}

type instruction struct {
	operator int
	pc       int
}

func New(input string, r io.Reader, w io.Writer) *Evaluator {
	return &Evaluator{
		lexer:  lexer.New(input),
		reader: bufio.NewReader(r),
		writer: w,
	}
}

func (e *Evaluator) Evaluate(ctx context.Context) error {
	if err := e.compile(); err != nil {
		return err
	}
	return e.execute(ctx)
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
				return compileError(t, "undefined `wWW`")
			}
			jmpLabel := jmpStack[len(jmpStack)-1]
			jmpStack = jmpStack[:len(jmpStack)-1]
			e.program = append(e.program, &instruction{
				operator: bck,
				pc:       jmpLabel,
			})
			e.program[jmpLabel].pc = pc
		case token.ILLEGAL:
			return compileError(t, "compile illegal token")
		}
	}
	return nil
}

func (e *Evaluator) execute(ctx context.Context) error {
	var ptr int
	data := make([]byte, stackSize)
	for pc := 0; pc < len(e.program); pc++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if len(data) <= ptr {
				return fmt.Errorf("stack exceeds %d-byte limit", stackSize)
			}
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
				e.writer.Write([]byte{data[ptr]})
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
	}
	return nil
}

func compileError(t *token.Token, msg string) error {
	return fmt.Errorf("compile:%d:%d: %s", t.Line, t.Col, msg)
}
