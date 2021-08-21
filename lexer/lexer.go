package lexer

import (
	"Moonlight/position"
	"Moonlight/token"
)

type Lexer struct {
	Source string

	Filename string

	position     position.Position
	currentToken token.Token
}

func GenTokens(lexer *Lexer) []token.Token {
	var tokens []token.Token
	column := 0
	line := 0
	for _, char := range lexer.Source {
		column += 1
		lexer.position = position.Position{
			Filename: lexer.Filename,
			Column:   column,
			Line:     line,
		}
		lexer.currentToken = token.Token{TokenType: token.INVALID, Position: lexer.position}
		switch char {
		// WHITESPACES
		case '\n':
			if lexer.currentToken.TokenType != token.STRING && lexer.currentToken.TokenType != token.COMMENT {
				line += 1
				column = 0
			}
			fallthrough
		case ' ':
			fallthrough
		case '\t':
			fallthrough
		case '\r':
			if lexer.currentToken.TokenType != token.WHITESPACE {
				lexer.currentToken.TokenType = token.WHITESPACE
			}
			fallthrough

		// NUMERIC LITERALS
		case '0':
			fallthrough
		case '1':
			fallthrough
		case '2':
			fallthrough
		case '3':
			fallthrough
		case '4':
			fallthrough
		case '5':
			fallthrough
		case '6':
			fallthrough
		case '7':
			fallthrough
		case '8':
			fallthrough
		case '9':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.INT
			}
			fallthrough

		// OPERATORS + SEPARATORS
		case ':':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.COLON
			}
			fallthrough
		case ',':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.COMMA
			}
			fallthrough
		case '.':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.DOT
			}
			fallthrough

		case '+':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.ADD
			}
			fallthrough
		case '-':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.SUB
			}
			fallthrough
		case '*':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.MUL
			}
			fallthrough
		case '/':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.DIV
			}
			fallthrough
		case '%':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.MOD
			}
			fallthrough
		case '=':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.ASSIGN
			}
			fallthrough
		case '<':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.LESS
			}
			fallthrough
		case '>':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.GREATER
			}
			fallthrough
		case '!':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.NOT
			}
			fallthrough

		// BRACKETS
		case '(':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.OPENINGPAREN
			}
			fallthrough
		case ')':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.CLOSINGPAREN
			}
			fallthrough
		case '[':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.OPENINGBRACKET
			}
			fallthrough
		case ']':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.CLOSINGBRACKET
			}
			fallthrough
		case '{':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.OPENINGBRACE
			}
			fallthrough
		case '}':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.CLOSINGBRACE
			}
			fallthrough

		// STRINGS
		case '"':
			fallthrough
		case '\'':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.STRING
			}
			fallthrough
		case '#':
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.COMMENT
			}
			lexer.currentToken.Value += string(char)
			break
		default:
			if lexer.currentToken.TokenType == token.INVALID {
				lexer.currentToken.TokenType = token.IDENTIFIER
			}
			lexer.currentToken.Value += string(char)
			break
		}
		if len(tokens) > 0 {
			if lexer.currentToken.TokenType != tokens[len(tokens)-1].TokenType {
				tokens = append(tokens, lexer.currentToken)
			} else if lexer.currentToken.TokenType == tokens[len(tokens)-1].TokenType {
				tokens[len(tokens)-1].Value += lexer.currentToken.Value
			}
		} else if len(tokens) == 0 {
			tokens = append(tokens, lexer.currentToken)
		}
	}
	return tokens
}

func ImproveTyping(tokens []token.Token) []token.Token {
	for index, currentToken := range tokens {
		if currentToken.Value == "==" && currentToken.TokenType != token.EQUAL {
			currentToken.TokenType = token.EQUAL
		} else if currentToken.Value == "!=" && currentToken.TokenType != token.NOTEQUAL {
			currentToken.TokenType = token.NOTEQUAL
		} else if currentToken.Value == "<=" && currentToken.TokenType != token.LESSEQUALS {
			currentToken.TokenType = token.LESSEQUALS
		} else if currentToken.Value == ">=" && currentToken.TokenType != token.GREATEREQUALS {
			currentToken.TokenType = token.GREATEREQUALS
		} else if currentToken.Value == "def" && currentToken.TokenType != token.FUNCTION {
			currentToken.TokenType = token.FUNCTION
		} else if currentToken.Value == "true" && currentToken.TokenType != token.TRUE {
			currentToken.TokenType = token.TRUE
		} else if currentToken.Value == "false" && currentToken.TokenType != token.FALSE {
			currentToken.TokenType = token.FALSE
		} else if currentToken.Value == "if" && currentToken.TokenType != token.IF {
			currentToken.TokenType = token.IF
		} else if currentToken.Value == "else" && currentToken.TokenType != token.ELSE {
			currentToken.TokenType = token.ELSE
		} else if currentToken.Value == "return" && currentToken.TokenType != token.RETURN {
			currentToken.TokenType = token.RETURN
		} else if currentToken.TokenType == token.DOT {
			println('1')
			if len(tokens) > index && index > 0 && tokens[index-1].TokenType == token.INT &&
				tokens[index+1].TokenType == token.INT {
				currentToken.TokenType = token.FLOAT
			}
		}
	}
	return tokens
}
