package token

type Type string

// Token is structure for identifying input stream of characters
type Token struct {
	Type    Type
	Literal string
	Col     int
	Line    int
}

// Literals
const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	INCR  = "www"
	DECR  = "WWW"
	NEXT  = "wWw"
	PREV  = "WwW"
	READ  = "wwWW"
	WRITE = "WWww"
	OPEN  = "wwWww"
	CLOSE = "WWwWW"
)
