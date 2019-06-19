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
			input: `wwwwwwwwwWWWWWWwWwwWwWwWWwWwwWwwWWwwWwwwWWwWWWWwWWw`,
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
				{token.READ, token.READ, 30, 1},
				{token.READ, token.READ, 33, 1},
				{token.WRITE, token.WRITE, 36, 1},
				{token.WRITE, token.WRITE, 39, 1},
				{token.OPEN, token.OPEN, 42, 1},
				{token.OPEN, token.OPEN, 45, 1},
				{token.CLOSE, token.CLOSE, 48, 1},
				{token.CLOSE, token.CLOSE, 51, 1},
			},
		},
		{
			name:  "Valid spaced",
			input: `www www www WWW WWW wWw wWw WwW WwW wwW wwW Www Www wWW wWW WWw WWw`,
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
				{token.READ, token.READ, 39, 1},
				{token.READ, token.READ, 43, 1},
				{token.WRITE, token.WRITE, 47, 1},
				{token.WRITE, token.WRITE, 51, 1},
				{token.OPEN, token.OPEN, 55, 1},
				{token.OPEN, token.OPEN, 59, 1},
				{token.CLOSE, token.CLOSE, 63, 1},
				{token.CLOSE, token.CLOSE, 67, 1},
			},
		},
		{
			name: "Valid spaced multi lines",
			input: `www www www WWW WWW wWw 
wWw WwW WwW wwW wwW Www Www wWW
wWW WWw WWw`,
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
				{token.READ, token.READ, 15, 2},
				{token.READ, token.READ, 19, 2},
				{token.WRITE, token.WRITE, 23, 2},
				{token.WRITE, token.WRITE, 27, 2},
				{token.OPEN, token.OPEN, 31, 2},
				{token.OPEN, token.OPEN, 3, 3},
				{token.CLOSE, token.CLOSE, 7, 3},
				{token.CLOSE, token.CLOSE, 11, 3},
			},
		},
		{
			name:  "Illegal",
			input: `wwwwwwwwwWWWWWWwWwwWwWwWWwWwwWwwWWW`,
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
				{token.READ, token.READ, 30, 1},
				{token.READ, token.READ, 33, 1},
				{token.ILLEGAL, token.ILLEGAL, 36, 1},
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
