package lexer 

import (
    "testing"
    "GoblinShark/token"
)


func TestNextToken(t *testing.T) {
    input := `=+(){},;`

    tests = []struct {
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)
    m := 
