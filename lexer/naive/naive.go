package naive

import (
	"github.com/ryan-berger/monkey/token"
)

type Lexer struct {

}

func (l *Lexer) NextToken() token.Token {
	return token.Token{}
}

func NewLexer(input string) *Lexer  {
	return &Lexer{}
}