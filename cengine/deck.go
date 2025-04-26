package cengine

import "errors"
import "math/rand"

/*
cengine decks are effectively stack data structures with two backing arrays
and *should* support most stack operations
*/

// Here lies the implementation of piles

type Pile struct {
	backing []Card
	front   int
	clength int
	lmax    int
}

func NewPile(lmax int) Pile {
	if lmax < 1 {
		panic("pile length must be positive")
	}
	return Pile{
		backing: make([]Card, lmax),
		front:   -1,
		lmax:    lmax,
		clength: 0,
	}
}

func (p *Pile) push(c Card) error {
	if p.front == p.lmax-1 {
		return errors.New("pile is full")
	}
	p.front++
	p.backing[p.front] = c
	return nil

}

func (p *Pile) pop() (Card, error) {
	if p.front == -1 {
		return Card{Prop: 0}, errors.New("pile is empty")
	}
	p.front--
	return p.backing[p.front], nil
}
func (p *Pile) peek() (Card, error) {
	if p.front == -1 {
		return Card{Prop: 0}, errors.New("pile is empty")
	}
	return p.backing[p.front], nil
}

func (p *Pile) length() int {
	return p.front + 1
}

func (p *Pile) shuffle() {
	if p.front == -1 {
		return
	}
	var shuf []Card
	if p.front == p.lmax-1 {

		shuf = p.backing[:]
	} else {
		shuf = p.backing[:p.front+1]
	}
	rand.Shuffle(len(shuf), func(i, j int) {
		shuf[i], shuf[j] = shuf[j], shuf[i]
	})
	concat := append(make([]Card, p.lmax-len(shuf)), shuf...)
	p.backing = concat
}
