package ast

import (
	"bytes"
	"fmt"
	"github.com/ryan-berger/monkey/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

func (l *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("%s %s = ", l.TokenLiteral(), l.Name.String()))

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
