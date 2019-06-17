package token

type Type string

// Token is structure for identifying input stream of characters
type Token struct {
	Type    Type
	Literal string
	Line    int
}

// Literals
const (
	EOF = "EOF"

	INCR  = "wWw"
	DECR  = "WWw"
	NEXT  = "www"
	PREV  = "WwWw"
	READ  = "wWWW"
	WRITE = "wWwWw"
	OPEN  = "wwWww"
	CLOSE = "WWwWWW"
)
