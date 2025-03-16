package cengine

type Token int

const (
	ILLEGAL Token = iota
	WHITESPACE
	NEWLINE
	EOF

	// syntax for deckconf
	DECK
	CARD
	ENTRY
	IDENT

	eof = rune(0)
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func newScanner(d []rune) Scanner {
	return Scanner{
		data: d,
		loc:  -1,
	}
}

type Scanner struct {
	data []rune
	loc  int
}

func (s *Scanner) Scan() (Token, string) {
	// move head location
	s.loc++
	//ideally this should never be reached, but just in case
	if s.loc >= len(s.data) {
		return EOF, ""
	}
	c := s.data[s.loc]
	if isWhitespace(c) {
		s.loc--
		//consume whitespace
		return s.ScanWhitespace()
	} else if isLetter(c) {
		s.loc--
		//consume ident
		return s.ScanIdent()
	}
	switch c {
	case eof:
		return EOF, ""
	default:
		return ILLEGAL, string(c)
	}
}

func (s *Scanner) ScanWhitespace() (Token, string) {
	buf := ""
	for {
		s.loc++
		//ideally this should never be reached, but just in case
		if s.loc == len(s.data) {
			return EOF, buf
		}
		c := s.data[s.loc]
		if c == eof {
			return EOF, ""
		} else if !isWhitespace(c) {
			s.loc--
			break
		} else {
			buf = buf + string(c)
		}
	}
	return WHITESPACE, buf
}

// TODO: Look here if we want to add more syntax tokens
func (s *Scanner) ScanIdent() (Token, string) {
	buf := ""
	for {
		s.loc++
		//ideally this should never be reached, but just in case
		if s.loc == len(s.data) {
			return EOF, buf
		}
		c := s.data[s.loc]
		if c == eof {
			return EOF, ""
		} else if isWhitespace(c) {
			s.loc--
			break
		} else {
			buf += string(c)
		}
	}
	switch buf {
	case "DECK":
		return DECK, buf
	case "CARD":
		return CARD, buf
	case "ENTRY":
		return ENTRY, buf
	default:
		return IDENT, buf
	}
}
