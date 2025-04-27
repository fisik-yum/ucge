package cengine

import "errors"
import "math/rand"

/*
Cengine piles are effectively stack data structures with two backing arrays
and *should* support most stack operations

The pile type is backed by a constant length slice that is determined at
initialization. Extreme care is taken to not reallocate memory.
*/ // Here lies the implementation of piles

type Pile struct {
	Backing []Card `json:"Backing"`
	Front   int    `json:"Front"`
	Clength int    `json:"Clength"`
	Lmax    int    `json:"Lmax"`
}

func NewPile(lmax int) Pile {
	if lmax < 1 {
		panic("pile length must be positive")
	}
	return Pile{
		Backing: make([]Card, lmax),
		Front:   -1,
		Lmax:    lmax,
		Clength: 0,
	}
}

func (p *Pile) push(c Card) error {
	if p.Front == p.Lmax-1 {
		return errors.New("pile is full")
	}
	p.Front++
	p.Backing[p.Front] = c
	return nil

}

func (p *Pile) pop() (Card, error) {
	if p.Front == -1 {
		return Card{Prop: 0}, errors.New("pile is empty")
	}
	p.Front--
	return p.Backing[p.Front+1], nil
}
func (p *Pile) peek() (Card, error) {
	if p.Front == -1 {
		return Card{Prop: 0}, errors.New("pile is empty")
	}
	return p.Backing[p.Front], nil
}

func (p *Pile) length() int {
	 return p.Front + 1
}

func (p *Pile) shuffle() {
	if p.Front == -1 {
		return
	} 
	var shuf []Card
	if p.Front == p.Lmax-1 {

		shuf = p.Backing[:]
	} else {
		shuf = p.Backing[:p.Front+1]
	}
	rand.Shuffle(len(shuf), func(i, j int) {
		shuf[i], shuf[j] = shuf[j], shuf[i]
	})
	concat := append(make([]Card, p.Lmax-len(shuf)), shuf...)
	p.Backing = concat
}

func (p*Pile)popN(n int) ([]Card, error){
 if(n>p.length()){
	return []Card{},errors.New("pile too small")
 }
 ret:= p.Backing[p.Front:p.Front+n]
p.Front+=n
return ret,nil
}

func (p*Pile) String() string{
	ret:=""
	for i:=p.Front;i>-1;i--{
		ret=ret+p.Backing[i].String()
	}
	return ret
}
