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
	line         int
}

// New initializes a new lexer with input string
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipComments()
	l.readChar()
	for {
		switch {
		case strings.ContainsRune(token.INCR, l.ch),
			strings.ContainsRune(token.DECR, l.ch),
			strings.ContainsRune(token.NEXT, l.ch),
			strings.ContainsRune(token.PREV, l.ch),
			strings.ContainsRune(token.READ, l.ch),
			strings.ContainsRune(token.WRITE, l.ch),
			strings.ContainsRune(token.OPEN, l.ch),
			strings.ContainsRune(token.CLOSE, l.ch):
			continue
		}
	}
	return token.Token{}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ascii code's null
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipComments() {
	for l.ch != 'w' || l.ch != 'W' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}
