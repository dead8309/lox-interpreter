package main

import "fmt"

type TokenType int

const (
	// Single character tokens
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	// Literals
	IDENTIFIER
	STRING
	NUMBER
	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FOR
	FUN
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	ERROR
	// EOF
	EOF
)

var tokeTypeName = map[TokenType]string{
	LeftParen:  "LEFT_PAREN",
	RightParen: "RIGHT_PAREN",
	LeftBrace:  "LEFT_BRACE",
	RightBrace: "RIGHT_BRACE",
	COMMA:      "COMMA",
	DOT:        "DOT",
	MINUS:      "MINUS",
	PLUS:       "PLUS",
	SEMICOLON:  "SEMICOLON",
	SLASH:      "SLASH",
	STAR:       "STAR",
	// One or two character tokens
	BANG:          "BANG",
	BANG_EQUAL:    "BANG_EQUAL",
	EQUAL:         "EQUAL",
	EQUAL_EQUAL:   "EQUAL_EQUAL",
	GREATER:       "GREATER",
	GREATER_EQUAL: "GREATER_EQUAL",
	LESS:          "LESS",
	LESS_EQUAL:    "LESS_EQUAL",
	// Literals
	IDENTIFIER: "IDENTIFIER",
	STRING:     "STRING",
	NUMBER:     "NUMBER",
	// Keywords
	AND:    "AND",
	CLASS:  "CLASS",
	ELSE:   "ELSE",
	FALSE:  "FALSE",
	FOR:    "FOR",
	FUN:    "FUN",
	IF:     "IF",
	NIL:    "NIL",
	OR:     "OR",
	PRINT:  "PRINT",
	RETURN: "RETURN",
	SUPER:  "SUPER",
	THIS:   "THIS",
	TRUE:   "TRUE",
	VAR:    "VAR",
	WHILE:  "WHILE",

	ERROR: "ERROR",
	// EOF
	EOF: "EOF",
}

func (t *TokenType) String() string {
	return tokeTypeName[*t]
}

type Token struct {
	Type    TokenType
	Laxeme  string
	Literal any
	Line    int
}

func (t *Token) String() string {
	l := "null"
	return fmt.Sprintf("%s %s %v", t.Type.String(), t.Laxeme, l)
}
