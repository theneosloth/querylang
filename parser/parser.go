package parser

import (
	"fmt"
	"io"

	"github.com/theneosloth/querylang/lexer"
)

type Query struct {
	Subject   string
	Operation string
	Query     string
}

type Group struct {
	Queries []Query
}

type Parser struct {
	s   *lexer.Scanner
	buf struct {
		tok lexer.Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: lexer.NewScanner(r)}
}

func (p *Parser) scan() (tok lexer.Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhitespace() (tok lexer.Token, lit string) {
	tok, lit = p.scan()
	if tok == lexer.WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) Parse() (*Group, error) {
	group := &Group{}
	for {
		query := Query{}
		tok, lit := p.scanIgnoreWhitespace()
		if tok == lexer.EOF {
			return group, nil
		}
		if tok != lexer.IDENT {
			return nil, fmt.Errorf("Found %q, expected an identifier", lit)
		}
		query.Subject = lit

		// If the next token is a separator, otherwise we're done
		tok, tokLit := p.scan()
		if !lexer.IsSeparator(tok) {
			group.Queries = append(group.Queries, query)
			p.unscan()
			continue
		}

		q, lit := p.scanIgnoreWhitespace()
		if q != lexer.IDENT {
			return nil, fmt.Errorf("Expected identifier, found %q", lit)
		}
		query.Query = lit
		query.Operation = tokLit
		group.Queries = append(group.Queries, query)

	}
}

// func (q *Query) String() string {
// 	return fmt.Sprintf("[%s %s %s]", q.Operation, q.Subject, q.Query)
// }

// func (g *Group) String() string {
// 	out := "{\n"
// 	for _, q := range g.Queries {
// 		out = fmt.Sprintf("%s\t%s\n", out, q.String())
// 	}
// 	out = fmt.Sprintf("%s\n}\n", out)
// 	return out
// }
