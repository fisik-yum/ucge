package cengine

import (
	"fmt"
	"math/rand"
	"slices"
)

/*
Cengine handles most operations in terms of Cards and Decks. A Card is a single
unit with specific properties and type.
*/

const (
	ToneSerious uint8 = iota
	ToneNeutral
	ToneFriendly
)

const (
	ResourceCoins uint8 = iota + 16
	ResourceWater
	ResourceFood
)

type Card struct {
	Prop uint8 `json:"Prop"`
}

func (b Card) Property() uint8 {
	return b.Prop
}

func (c *Card) Equals(o Card) bool {
	return c.Prop == o.Prop
}

func (b *Card) String() string {
	return fmt.Sprintf("%d", b.Prop)
}

/*
Cards are grouped into Decks, which maintain an active and discard pile. Decks
can generate new Hands, which are the suggested method to expose game state to
a player. Optionally Decks can be part of a Collection, which organizes each
Deck by an identifier, creating a quasi-inventory to manage different Decks.
*/
type Collection map[string]*Deck

func NewInventory() Collection {
	return make(Collection)
}

type Deck struct {
	Active  []Pile `json:"Active"`
	Discard []Pile `json:"Discard"`
}

func NewDeck(length int) Deck {
	return Deck{
		Active:  NewPile(length),
		Discard: NewPile(length),
	}
}

func (d *Deck) String() string {
	return fmt.Sprintf("Active: %v\nDiscard: %v", d.Active, d.Discard)
}

func (d *Deck) SetAt(i int, n *Card) {
	*&d.Active[i] = *n
}

func (d *Deck) GetAt(i int) Card {
	return *&d.Active[i]
}

func (d *Deck) Shuffle() {
	dr := *d
	rand.Shuffle(len(dr.Active), func(i, j int) {
		dr.Active[i], dr.Active[j] = dr.Active[j], dr.Active[i]
	})
}

func (d *Deck) DiscardAt(i int) {
	d.Discard = append(d.Discard, d.Active[i])
	d.Active = slices.Delete(d.Active, i, i+1)
}

func (d *Deck) NewHand(n int) *Hand {
	h := Hand{
		Cards: d.Active[0:n],
	}
	d.Active = d.Active[n:]
	return &h
}

type Hand struct {
	Cards []Card `json:"Cards"`
}
