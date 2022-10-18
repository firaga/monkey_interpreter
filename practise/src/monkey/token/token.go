// token/token.go

package token

//let five = 5;
//let ten = 10;
//
//let add = fn(x, y) {
//  x + y;
//};
//
//let result = add(five, ten);

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// token/token.go

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符+字面量
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// 运算符
	ASSIGN = "="
	PLUS   = "+"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		//返回匹配的keyword
		return tok
	}
	//返回identifier
	return IDENT
}
