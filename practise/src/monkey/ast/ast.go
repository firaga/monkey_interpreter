package ast

import "monkey/token"

// Node 节点
type Node interface {
	TokenLiteral() string
}

// Statement 语句 语句不会产生值
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式 表达式会产生值
type Expression interface {
	Node
	expressionNode()
}

// ast/ast.go

// Program monkey程序根节点
type Program struct {
	Statements []Statement //语句切片
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ast/ast.go

// [...]

type LetStatement struct {
	Token token.Token // token.LET词法单元
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier 标识符在let中不会产生值,如let x=5,但标识符也会用在表达式的位置,如 let x= y,y会产生值,所以算表达式
type Identifier struct {
	Token token.Token // token.IDENT词法单元
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ast/ast.go

type ReturnStatement struct {
	Token       token.Token // 'return'词法单元
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
