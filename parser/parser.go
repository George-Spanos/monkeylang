package parser

import (
	"monkeylang/ast"
	"monkeylang/lexer"
	"monkeylang/token"
)

type Parser struct {
	l *lexer.Lexer

	curlToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curlToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curlToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
