package token

import "Moonlight/position"

type Type string

type Token struct {
	TokenType Type
	Value     string

	Position position.Position
}

const (
	INVALID    = "Invalid"
	COMMENT    = "Comment"
	WHITESPACE = "Whitespace"

	IDENTIFIER = "Identifier" // add, x ,y, ...
	INT        = "Int"        // 123456
	FLOAT      = "Float"
	STRING     = "String" // "x", "y"

	ADD      = "+"
	SUB      = "-"
	MUL      = "*"
	DIV      = "/"
	MOD      = "%"
	ASSIGN   = "="
	NOT      = "!"
	EQUAL    = "=="
	NOTEQUAL = "!="

	LESS          = "<"
	GREATER       = ">"
	LESSEQUALS    = "<="
	GREATEREQUALS = ">="

	DOT   = "."
	COMMA = ","
	COLON = ":"

	OPENINGPAREN   = "("
	CLOSINGPAREN   = ")"
	OPENINGBRACE   = "{"
	CLOSINGBRACE   = "}"
	OPENINGBRACKET = "["
	CLOSINGBRACKET = "]"

	FUNCTION = "Function"
	TRUE     = "True"
	FALSE    = "False"
	IF       = "If"
	ELSE     = "Else"
	RETURN   = "Return"
)
