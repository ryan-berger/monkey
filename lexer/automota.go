package lexer

import (
	"strings"
	"unicode"

	"github.com/ryan-berger/monkey/token"
)

func lexText(lexer *lexer) stateFn {
	for {
		if lexer.input[lexer.start:lexer.pos] == "let" {
			lexer.emit(token.LET)
			return lexIdentifier
		}

		next := lexer.next()
		if next == eof {
			lexer.emit(token.EOF)
			return lexer.errorf("reached eof")
		}
	}
}

func lexIdentifier(lexer *lexer) stateFn {
	for {
		if string(lexer.peek()) == "=" {
			lexer.emit(token.IDENT)
			return lexAssign
		}

		next := lexer.next()

		if next == eof {
			lexer.emit(token.EOF)
			return lexer.errorf("could not find equal sign")
		}
	}
}

func lexAssign(lexer *lexer) stateFn {
	lexer.pos += len(token.ASSIGN)
	lexer.emit(token.ASSIGN)
	return lexExpression
}

func lexExpression(lexer *lexer) stateFn {
	for {
		if strings.HasPrefix(lexer.input[lexer.pos:], "fn") {
			return lexFunction
		}

		if next := lexer.next(); unicode.IsSpace(next) {
			lexer.ignore()
		}
	}
}

func lexFunction(lexer *lexer) stateFn {
	lexer.pos += len("fn")
	lexer.emit(token.FUNCTION)
	return lexBeginArgList
}

func lexBeginArgList(lexer *lexer) stateFn {
	for {
		switch next := lexer.next(); {
		case strings.HasPrefix(lexer.input[lexer.start:], "("):
			lexer.emit(token.LPAREN)
			return lexArgList
		case unicode.IsSpace(next):
			lexer.ignore()
		}
	}

	return nil
}

func lexArgList(lexer *lexer) stateFn {
	for {
		if strings.HasPrefix(lexer.input[lexer.pos:], ")") {
			return lexEndArgList
		}
		lexer.next()
		return nil
	}
}

func lexEndArgList(lexer *lexer) stateFn {
	lexer.pos += len(")")
	lexer.emit(token.RPAREN)
	return lexBeginFunc
}

func lexBeginFunc(lexer *lexer) stateFn {
	for {
		switch next := lexer.next(); {
		case strings.HasPrefix(lexer.input[lexer.start:], "{"):
			lexer.emit(token.LBRACE)
			return lexText
		case unicode.IsSpace(next):
			lexer.ignore()
		case next == eof:
			return lexer.errorf("expected { found eof")
		}
	}
}

func lexNumericalExpression(lexer *lexer) stateFn {
	return nil
}
