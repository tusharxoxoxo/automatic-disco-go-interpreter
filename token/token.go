//token/token.go 

//in our token package we are going to define our token struct and token type

package token

type TokenType string

type Token struct{
    Type TokenType
    Literal string
}

//let five = 5;
//let ten = 10;
//
//let add = fn(x, y) {
//    x + y;
//};  
//
//let result = add(five, ten);


const (
    //Special Types
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    //Identifiers + Literals
    IDENT = "IDENT" //add, foobar, x, y
    INT = "INT" //123456

    //Operators
    ASSIGN = "="
    PLUS = "+"

    //Delimiters    
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    //Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"


)



