package lexer

import "github.com/ryan-berger/monkey/token"

type Lexer interface {
	NextToken() token.Token
}

type Builder func(input string) Lexer