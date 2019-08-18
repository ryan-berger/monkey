package parser

import (
	"github.com/ryan-berger/monkey/ast"
	"github.com/ryan-berger/monkey/lexer"
	"github.com/ryan-berger/monkey/token"
)

type Parser struct {
	l         lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(lexer lexer.Lexer) *Parser  {
	p := &Parser{
		l: lexer,
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}