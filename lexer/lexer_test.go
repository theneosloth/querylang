package lexer

import (
	"strings"
	"testing"
)

type result struct {
	token   Token
	literal string
}

func TestScan(t *testing.T) {

	tests := map[string]result{
		"query": {
			IDENT,
			"query",
		},
		">=": {
			GE,
			">=",
		},
		"\t": {
			WS,
			"\t",
		},
	}

	for test, result := range tests {
		s := NewScanner(strings.NewReader(test))
		tok, lit := s.Scan()
		if tok != result.token {
			t.Errorf("Result does not match expected. Input: %v, Expected: %v", tok, result.token)
		}

		if lit != result.literal {
			t.Errorf("Literal not scanned. Input: %s, Expected: %s", lit, result.literal)
		}
	}
}
