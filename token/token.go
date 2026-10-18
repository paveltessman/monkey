package token

type Type string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(literal string) Type {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
