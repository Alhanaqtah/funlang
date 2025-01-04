package lexer

import (
	"funlang/internal/token"
	"unicode"
)

type Lexer struct {
	input   []rune
	pos     int  // pointer to current charecter in input
	readPos int  // pointer to charecter readind in input
	ch      rune // current charecter
}

func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}

	l.readChar()

	return l
}

func (l *Lexer) Next() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = token.New(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = token.New(token.PLUS, string(l.ch))
	case '-':
		tok = token.New(token.MINUS, string(l.ch))
	case '/':
		tok = token.New(token.SLASH, string(l.ch))
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = token.New(token.BANG, string(l.ch))
		}
	case '<':
		tok = token.New(token.LT, string(l.ch))
	case '>':
		tok = token.New(token.GT, string(l.ch))
	case ',':
		tok = token.New(token.COMMA, string(l.ch))
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch))
	case '(':
		tok = token.New(token.LPAREN, string(l.ch))
	case ')':
		tok = token.New(token.RPAREN, string(l.ch))
	case '{':
		tok = token.New(token.LBRACE, string(l.ch))
	case '}':
		tok = token.New(token.RBRACE, string(l.ch))
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if unicode.IsLetter(l.ch) {
			tok.Literal = string(l.readIdentifier())
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = string(l.readNumber())
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) readIdentifier() []rune {
	pos := l.pos

	for unicode.IsLetter(l.ch) || l.ch == '_' {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() []rune {
	pos := l.pos

	for unicode.IsDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}
