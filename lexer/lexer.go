package lexer

import (
	"strings"

	"github.com/Code-Hex/shibafu/token"
)

// Lexer is used for tokenizing programs
type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
	col          int
	line         int
}

// New initializes a new lexer with input string
func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
		line:  1,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	l.skipComments()
	return l.readIdents()
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ascii code's null (EOF)
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.col++
}

func (l *Lexer) skipComments() {
	for l.ch != 'w' && l.ch != 'W' {
		if l.ch == '\n' {
			l.col = 0
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) readIdents() *token.Token {
	position := l.position
	defer l.readChar()

	for l.ch == 'w' || l.ch == 'W' {
		input := string(l.input[position:l.readPosition])

		switch input {
		case token.INCR:
			return l.newToken(token.INCR)
		case token.DECR:
			return l.newToken(token.DECR)
		case token.NEXT:
			return l.newToken(token.NEXT)
		case token.PREV:
			return l.newToken(token.PREV)
		case token.READ:
			return l.newToken(token.READ)
		case token.WRITE:
			return l.newToken(token.WRITE)
		case token.OPEN:
			return l.newToken(token.OPEN)
		case token.CLOSE:
			return l.newToken(token.CLOSE)
		}

		switch {
		case strings.HasPrefix(token.INCR, input),
			strings.HasPrefix(token.DECR, input),
			strings.HasPrefix(token.NEXT, input),
			strings.HasPrefix(token.PREV, input),
			strings.HasPrefix(token.READ, input),
			strings.HasPrefix(token.WRITE, input),
			strings.HasPrefix(token.OPEN, input),
			strings.HasPrefix(token.CLOSE, input):
			l.readChar()
			continue
		}

		break
	}
	return l.newToken(token.ILLEGAL)
}

func (l *Lexer) newToken(typ token.Type) *token.Token {
	return &token.Token{
		Type:    typ,
		Literal: string(typ),
		Col:     l.col,
		Line:    l.line,
	}
}
