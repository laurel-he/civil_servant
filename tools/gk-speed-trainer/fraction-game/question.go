package main

import (
	"math/rand"
	"time"
)

type QuestionPool struct {
	list []Question
	idx  int
}

func NewPool() *QuestionPool {
	p := &QuestionPool{}
	p.shuffle()
	return p
}

func (p *QuestionPool) shuffle() {
	p.list = make([]Question, len(Questions))
	copy(p.list, Questions)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(p.list), func(i, j int) {
		p.list[i], p.list[j] = p.list[j], p.list[i]
	})

	p.idx = 0
}

func (p *QuestionPool) Next() Question {

	if p.idx >= len(p.list) {
		p.shuffle()
	}

	q := p.list[p.idx]
	p.idx++

	return q
}