package lexer

import (
	"bufio"
	"bytes"
	"io"
)

var eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func IsSeparator(ident Token) bool {
	separators := []Token{EQ, GE, LE, GT, LT, NE}
	for _, j := range separators {
		if j == ident {
			return true
		}
	}
	return false
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// peek returns the next character. Only reads 1 byte ahead so may run into issues with unicode
func (s *Scanner) peek() rune {
	ch, err := s.r.Peek(1)
	if err != nil {
		return eof
	}

	return rune(ch[0])
}

func (s *Scanner) Scan() (tok Token, li string) {
	ch := s.peek()

	if isWhitespace(ch) {
		return s.scanWhitespace()
	} else if isLetter(ch) || isDigit(ch) {
		return s.scanIdent()
	}
	return s.scanSeparator()
}

func (s *Scanner) scanSeparator() (tok Token, lit string) {

	ch := s.read()
	switch ch {
	case eof:
		return EOF, ""
	case '=':
		return EQ, string(ch)
	case '>':
		if s.read() == '=' {
			return GE, ">="
		} else {
			s.unread()
		}
		return GT, string(ch)
	case '<':
		if s.read() == '=' {
			return LE, "<="
		} else {
			s.unread()
		}
		return LT, string(ch)
	case '!':
		if s.read() == '=' {
			return NE, "!="
		} else {
			s.unread()
		}
	}
	return ILLEGAL, string(ch)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return IDENT, buf.String()
}
