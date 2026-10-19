package lexer

import "github.com/paveltessman/monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {

	case '=':

		tok = newToken(token.ASSIGN, lexer.ch)

	case ';':

		tok = newToken(token.SEMICOLON, lexer.ch)

	case ',':

		tok = newToken(token.COMMA, lexer.ch)

	case '+':

		tok = newToken(token.PLUS, lexer.ch)

	case '-':

		tok = newToken(token.MINUS, lexer.ch)

	case '!':

		tok = newToken(token.BANG, lexer.ch)

	case '/':

		tok = newToken(token.SLASH, lexer.ch)

	case '*':

		tok = newToken(token.ASTERISK, lexer.ch)

	case '<':

		tok = newToken(token.LT, lexer.ch)

	case '>':

		tok = newToken(token.GT, lexer.ch)

	case '(':

		tok = newToken(token.LPAREN, lexer.ch)

	case ')':

		tok = newToken(token.RPAREN, lexer.ch)

	case '{':

		tok = newToken(token.LBRACE, lexer.ch)

	case '}':

		tok = newToken(token.RBRACE, lexer.ch)

	case 0:

		tok.Type = token.EOF
		tok.Literal = ""

	default:

		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		}

		tok = newToken(token.ILLEGAL, lexer.ch)
	}

	lexer.readChar()

	return tok
}

func (lexer *Lexer) readChar() {
	lexer.ch = 0

	if lexer.readPosition < len(lexer.input) {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func (lexer *Lexer) readIdentifier() string {
	start := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[start:lexer.position]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readNumber() string {
	start := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[start:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
