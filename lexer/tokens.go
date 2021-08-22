package lexer

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	IDENT
	GT
	LT
	GE
	LE
	EQ
	NE
	WS
)
