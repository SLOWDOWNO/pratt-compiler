package parser

import (
	"ppc/lexer"
	"ppc/token"
	"strconv"
)

type Parser struct {
	lexer   *lexer.Lexer
	current token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer, current: lexer.NextToken()}
}

func (p *Parser) nextToken() {
	p.current = p.lexer.NextToken()
}

func (p *Parser) ParseExpression(precedence int) int {
	left := p.parsePrefix()

	for precedence < p.getPrecedence(p.current.Type) {
		left = p.parseInfix(left)
	}

	return left
}

func (p *Parser) getPrecedence(tokenType token.TokenType) int {
	switch tokenType {
	case token.PLUS, token.MINUS:
		return 1
	case token.ASTERISK, token.SLASH:
		return 2
	}
	return 0
}

func (p *Parser) parsePrefix() int {
	switch p.current.Type {
	case token.NUMBER:
		// 处理数字
		value, _ := strconv.Atoi(p.current.Value)
		p.nextToken()
		return value
	case token.MINUS:
		// 处理负号
		p.nextToken()
		value := p.ParseExpression(p.getPrecedence(token.MINUS))
		return -value
		// 处理左括号
	case token.LPAREN:
		p.nextToken()
		// 递归解析括号内的表达式
		value := p.ParseExpression(0)
		if p.current.Type != token.RPAREN {
			panic("missing closing parenthesis")
		}
		p.nextToken()
		return value
	default:
		panic("unexpected tokken: " + p.current.Value)
	}
}

func (p *Parser) parseInfix(left int) int {
	op := p.current
	precedence := p.getPrecedence(op.Type)
	p.nextToken()
	right := p.ParseExpression(precedence)
	return p.applyOperator(left, right, op.Type)
}

func (p *Parser) applyOperator(left, right int, op token.TokenType) int {
	switch op {
	case token.PLUS:
		return left + right
	case token.MINUS:
		return left - right
	case token.ASTERISK:
		return left * right
	case token.SLASH:
		return left / right
	}
	panic("unexpected operator: " + strconv.Itoa(int(op)))
}
