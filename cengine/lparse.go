package cengine

import (
	"errors"
	"log"
	"strconv"
)

type Parser struct {
	s   *Scanner
	buf struct {
		lastTok Token  // last read token
		lastLit string // last read literal
		n       int    // buffer size (max=1)
	}
}

func newParser(d []rune) Parser {
	a := newScanner(d)
	return Parser{
		s: &a,
	}
}

type LoaderData struct {
	Decks      Collection
	EntryPoint string
}

func (p *Parser) Parse() (ld LoaderData) {
	//rewrite parsing logic
	ld = LoaderData{
		Decks:      make(Collection),
		EntryPoint: ""}
	for {
		t, s := p.parseSkipWhitespace()
		switch t {
		case EOF:
			return ld
		case DECK:
			t, s = p.parseExpect(IDENT)
			ld.Decks[s] = &Deck{Active: make([]Card, 0), Discard: make([]Card, 0)}
		case CARD:
			_, id := p.parseExpect(IDENT)
			_, s = p.parseExpect(IDENT)
			val, err := strconv.ParseUint(id, 10, 8)
			if err != nil {
				log.Fatal(sendErr())
			}
			ld.Decks[s].Active = append(ld.Decks[s].Active, Card{Prop: uint8(val)})
		case ENTRY:
			t, s = p.parseExpect(IDENT)
			ld.EntryPoint = s
		case ILLEGAL:
			log.Fatal(sendErr())
		}
	}
}

// Move to next token, skipping whitespace
func (p *Parser) parseSkipWhitespace() (t Token, s string) {
	for {
		t, s = p.s.Scan()
		if t != WHITESPACE {
			break
		}
	}
	return
}

// Move to next token, expecting Ident
func (p *Parser) parseExpect(e Token) (t Token, s string) {
	t, s = p.parseSkipWhitespace()
	if t == e {
		return
	}
	log.Fatal("Unexpected token")
	return
}

func sendErr() error {
	return errors.New("Illegal token encountered")
}
