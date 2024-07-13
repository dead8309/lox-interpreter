package main

import (
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	Source  []byte
	Tokens  []Token
	start   int
	current int
	line    int
	errors  int
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
		errors:  0,
	}
}

func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) Advance() byte {
	v := s.Source[s.current]
	s.current += 1
	return v
}

func (s *Scanner) Peek() byte {
	if s.IsAtEnd() {
		return '\x00'
	}
	return s.Source[s.current]
}

func (s *Scanner) PeekNext() byte {
	if s.current+1 >= len(s.Source) {
		return '\x00'
	}
	return s.Source[s.current+1]
}

func (s *Scanner) AddToken(t TokenType, literal any) {
	text := string(s.Source[s.start:s.current])
	s.Tokens = append(s.Tokens, Token{t, text, literal, s.line})
}

func (s *Scanner) ScanToken() {
	switch c := s.Advance(); c {
	case '(':
		s.AddToken(LeftParen, nil)
	case ')':
		s.AddToken(RightParen, nil)
	case '{':
		s.AddToken(LeftBrace, nil)
	case '}':
		s.AddToken(RightBrace, nil)
	case ',':
		s.AddToken(COMMA, nil)
	case '.':
		s.AddToken(DOT, nil)
	case '-':
		s.AddToken(MINUS, nil)
	case '+':
		s.AddToken(PLUS, nil)
	case ';':
		s.AddToken(SEMICOLON, nil)
	case '*':
		s.AddToken(STAR, nil)
	case '/':
		if !s.IsAtEnd() && s.Source[s.current] == '/' {
			for s.Peek() != '\n' && !s.IsAtEnd() {
				s.Advance()
			}
		} else {
			s.AddToken(SLASH, nil)
		}
	case '\n':
		s.line++
	case ' ', '\r', '\t':
		break
	case '=':
		if !s.IsAtEnd() && s.Source[s.current] == '=' {
			s.Advance()
			s.AddToken(EQUAL_EQUAL, nil)
		} else {
			s.AddToken(EQUAL, nil)
		}
	case '!':
		if !s.IsAtEnd() && s.Source[s.current] == '=' {
			s.Advance()
			s.AddToken(BANG_EQUAL, nil)
		} else {
			s.AddToken(BANG, nil)
		}
	case '<':
		if !s.IsAtEnd() && s.Source[s.current] == '=' {
			s.Advance()
			s.AddToken(LESS_EQUAL, nil)
		} else {
			s.AddToken(LESS, nil)
		}
	case '>':
		if !s.IsAtEnd() && s.Source[s.current] == '=' {
			s.Advance()
			s.AddToken(GREATER_EQUAL, nil)
		} else {
			s.AddToken(GREATER, nil)
		}
	case '"':
		s.ParseStrings()
	default:
		if IsDigit(c) {
			s.ParseDigits()
		} else if IsAlpha(c) {
			s.ParseIdentifier()
		} else {
			char := s.Source[s.start:s.current]
			s.error(fmt.Sprintf("Unexpected character: %s", char))
		}
	}
}

func (s *Scanner) ScanContent() []Token {
	for i := 0; !s.IsAtEnd(); i++ {
		s.start = s.current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, Token{EOF, "", nil, s.line})
	return s.Tokens
}

func (s *Scanner) error(msg string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error: %s\n", s.line, msg)
	s.errors++
}

func (s *Scanner) ParseStrings() {
	for s.Peek() != '"' && !s.IsAtEnd() {
		if s.Peek() == '\n' {
			s.line++
		}
		s.Advance()
	}
	if s.IsAtEnd() {
		s.error("Unterminated string.")
		return
	}
	// For closing "
	s.Advance()
	value := s.Source[s.start+1 : s.current-1]
	s.AddToken(STRING, value)
}

func (s *Scanner) ParseDigits() {
	for IsDigit(s.Peek()) {
		s.Advance()
	}

	if s.Peek() == '.' && IsDigit(s.PeekNext()) {
		// comsume '.'
		s.Advance()

		for IsDigit(s.Peek()) {
			s.Advance()
		}
	}

	value := string(s.Source[s.start:s.current])
	num, _ := strconv.ParseFloat(value, 64)
	s.AddToken(NUMBER, num)
}

func (s *Scanner) ParseIdentifier() {
	for IsAlphaNumeric(s.Peek()) {
		s.Advance()
	}
	s.AddToken(IDENTIFIER, nil)
}

func IsDigit(s byte) bool {
	return s >= '0' && s <= '9'
}

func IsAlpha(s byte) bool {
	return (s >= 'a' && s <= 'z') ||
		(s >= 'A' && s <= 'Z') ||
		s == '_'
}

func IsAlphaNumeric(s byte) bool {
	return IsAlpha(s) || IsDigit(s)
}
