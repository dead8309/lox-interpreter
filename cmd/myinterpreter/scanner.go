package main

type Scanner struct {
	Source  []byte
	Tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
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

func (s *Scanner) AddToken(t TokenType, literal any) {
	text := string(s.Source[s.start:s.current])
	s.Tokens = append(s.Tokens, Token{t, text, nil, s.line})
}

func (s *Scanner) ScanToken() {
	switch s.Advance() {
	case '(':
		s.AddToken(LeftParen, nil)
	case ')':
		s.AddToken(RightParen, nil)
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
