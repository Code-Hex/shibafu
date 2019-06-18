package lexer

import (
	"testing"

	"github.com/Code-Hex/shibafu/token"
)

func TestNextToken(t *testing.T) {
	type want struct {
		Type    token.Type
		Literal string
		Col     int
		Line    int
	}

	tests := []struct {
		name  string
		input string
		wants []want
	}{
		{
			name:  "Valid joined",
			input: `wwwwwwwwwWWWWWWwWwwWwWwWWwWwwWWwwWWWWwwWWwwwwWwwwwWwwWWwWWWWwWW`,
			wants: []want{
				{token.INCR, token.INCR, 3, 1},
				{token.INCR, token.INCR, 6, 1},
				{token.INCR, token.INCR, 9, 1},
				{token.DECR, token.DECR, 12, 1},
				{token.DECR, token.DECR, 15, 1},
				{token.NEXT, token.NEXT, 18, 1},
				{token.NEXT, token.NEXT, 21, 1},
				{token.PREV, token.PREV, 24, 1},
				{token.PREV, token.PREV, 27, 1},
				{token.READ, token.READ, 31, 1},
				{token.READ, token.READ, 35, 1},
				{token.WRITE, token.WRITE, 39, 1},
				{token.WRITE, token.WRITE, 43, 1},
				{token.OPEN, token.OPEN, 48, 1},
				{token.OPEN, token.OPEN, 53, 1},
				{token.CLOSE, token.CLOSE, 58, 1},
				{token.CLOSE, token.CLOSE, 63, 1},
			},
		},
		{
			name:  "Valid spaced",
			input: `www www www WWW WWW wWw wWw WwW WwW wwWW wwWW WWww WWww wwWww wwWww WWwWW WWwWW`,
			wants: []want{
				{token.INCR, token.INCR, 3, 1},
				{token.INCR, token.INCR, 7, 1},
				{token.INCR, token.INCR, 11, 1},
				{token.DECR, token.DECR, 15, 1},
				{token.DECR, token.DECR, 19, 1},
				{token.NEXT, token.NEXT, 23, 1},
				{token.NEXT, token.NEXT, 27, 1},
				{token.PREV, token.PREV, 31, 1},
				{token.PREV, token.PREV, 35, 1},
				{token.READ, token.READ, 40, 1},
				{token.READ, token.READ, 45, 1},
				{token.WRITE, token.WRITE, 50, 1},
				{token.WRITE, token.WRITE, 55, 1},
				{token.OPEN, token.OPEN, 61, 1},
				{token.OPEN, token.OPEN, 67, 1},
				{token.CLOSE, token.CLOSE, 73, 1},
				{token.CLOSE, token.CLOSE, 79, 1},
			},
		},
		{
			name: "Valid spaced multi lines",
			input: `www www www WWW WWW wWw 
wWw WwW WwW wwWW wwWW WWww WWww wwWww
wwWww WWwWW WWwWW`,
			wants: []want{
				{token.INCR, token.INCR, 3, 1},
				{token.INCR, token.INCR, 7, 1},
				{token.INCR, token.INCR, 11, 1},
				{token.DECR, token.DECR, 15, 1},
				{token.DECR, token.DECR, 19, 1},
				{token.NEXT, token.NEXT, 23, 1},
				{token.NEXT, token.NEXT, 3, 2},
				{token.PREV, token.PREV, 7, 2},
				{token.PREV, token.PREV, 11, 2},
				{token.READ, token.READ, 16, 2},
				{token.READ, token.READ, 21, 2},
				{token.WRITE, token.WRITE, 26, 2},
				{token.WRITE, token.WRITE, 31, 2},
				{token.OPEN, token.OPEN, 37, 2},
				{token.OPEN, token.OPEN, 5, 3},
				{token.CLOSE, token.CLOSE, 11, 3},
				{token.CLOSE, token.CLOSE, 17, 3},
			},
		},
		{
			name:  "Illegal",
			input: `wwwwwwwwwWWWWWWwWwwWwWwWWwWwwWWwwWWWWw`,
			wants: []want{
				{token.INCR, token.INCR, 3, 1},
				{token.INCR, token.INCR, 6, 1},
				{token.INCR, token.INCR, 9, 1},
				{token.DECR, token.DECR, 12, 1},
				{token.DECR, token.DECR, 15, 1},
				{token.NEXT, token.NEXT, 18, 1},
				{token.NEXT, token.NEXT, 21, 1},
				{token.PREV, token.PREV, 24, 1},
				{token.PREV, token.PREV, 27, 1},
				{token.READ, token.READ, 31, 1},
				{token.READ, token.READ, 35, 1},
				{token.ILLEGAL, token.ILLEGAL, 39, 1},
			},
		},
		{
			name:  "EOF",
			input: `wwwwwwwww`,
			wants: []want{
				{token.INCR, token.INCR, 3, 1},
				{token.INCR, token.INCR, 6, 1},
				{token.INCR, token.INCR, 9, 1},
				{token.EOF, token.EOF, 10, 1},
				{token.EOF, token.EOF, 10, 1},
				{token.EOF, token.EOF, 10, 1},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.input)
			for i, want := range tc.wants {
				tok := l.NextToken()
				if tok.Type != want.Type {
					t.Errorf("tests[%d] - tokentype wrong. want=%q, got=%q", i, want.Type, tok.Type)
				}
				if tok.Literal != want.Literal {
					t.Errorf("tests[%d] - literal wrong. want=%q, got=%q", i, want.Literal, tok.Literal)
				}
				if tok.Col != want.Col {
					t.Errorf("tests[%d] - column wrong. want=%d, got=%d", i, want.Col, tok.Col)
				}
				if tok.Line != want.Line {
					t.Errorf("tests[%d] - line number wrong. want=%d, got=%d", i, want.Line, tok.Line)
				}
			}
		})
	}
}
