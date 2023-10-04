//token/token.go 

//in our token package we are going to define our token struct and token type

package token

type TokenType string

type Token struct{
    Type TokenType
    Literal string
}


