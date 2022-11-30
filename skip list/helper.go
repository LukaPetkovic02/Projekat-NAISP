package main

import (
	"fmt"
	"math/rand"
)

type SkipList struct {
	maxHeight int
	height    int
	size      int
	head      *SkipListNode
}

type SkipListNode struct {
	key   string
	value []byte
	next  []*SkipListNode
}

func (s *SkipListNode) InitSP(key string, value []byte, level int) SkipListNode {
	s.key = key
	s.value = value
	s.next = make([]*SkipListNode, level+1)
	return *s
}

func (s *SkipList) InitSP(maxHeight int, height int, size int, glava *SkipListNode) SkipList {
	s.maxHeight = maxHeight
	s.height = height
	s.size = size
	s.head = glava
	return *s
}

func (s *SkipList) roll() int {
	level := 0
	// possible ret values from rand are 0 and 1
	// we stop shen we get a 0
	for ; rand.Int31n(2) == 1; level++ {
		if level >= s.maxHeight {
			if level > s.height {
				s.height = level
			}
			return level
		}
	}
	if level > s.height {
		s.height = level
	}
	return level
}

func (s *SkipList) search(searchKey string) *SkipListNode { //vraca vrednost koja odgovara kljucu(ako postoji), ako ne postoji vraca nil
	x := s.head
	var i int
	for i = s.height; i >= 0; i-- {
		if x.key == searchKey {
			return x
		}
		for x.next[i] != nil && x.next[i].key <= searchKey {
			x = x.next[i]
		}
	}
	return nil
}

func (s *SkipList) insert(novi string, vrednost []byte) {
	if s.search(novi) != nil {
		return
	}
	var noviNode SkipListNode
	noviNode.InitSP(novi, vrednost, s.roll())
	x := s.head
	var i int

	for i = len(noviNode.next) - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < noviNode.key {
			x = x.next[i]
		}

		if x.next[i] != nil {
			noviNode.next[i] = x.next[i]
		}

		x.next[i] = &noviNode
	}
	fmt.Println("Uspesno ubaceno!!!")
}

func (s *SkipList) delete(brisanje string) {
	p := s.search(brisanje)
	if p == nil {
		fmt.Println("Nije pronadjen taj string.")
		return
	}

	x := s.head
	var i int
	for i = len(p.next) - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < brisanje {
			x = x.next[i]
		}

		x.next[i] = x.next[i].next[i]

	}
	fmt.Println("Uspesno obrisano!")

}
func main() {
	var a SkipListNode
	a.InitSP("koren", []byte("dsad"), 5)
	var b SkipList
	b.InitSP(10, 5, 10, &a)

	b.insert("12", []byte("poruka"))
	b.search("asd")

	fmt.Println(b.search("12").value)
	b.insert("asd", []byte("nebitno"))
	b.delete("12")
	b.delete("13")
	fmt.Println(b.search("12"))
}
