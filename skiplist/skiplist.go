package skipList

import (
	"fmt"
	"math/rand"
	Importi "projekat/utils"
)

type SkipList struct {
	maxHeight    int
	height       int
	size         int
	max_capacity int
	head         *SkipListNode
}

type SkipListNode struct {
	podatak Importi.Podatak
	next    []*SkipListNode
}

func (s *SkipListNode) InitSP(podatak Importi.Podatak, level int) SkipListNode {
	s.next = make([]*SkipListNode, level+1)
	s.podatak = podatak
	return *s
}

func (s *SkipList) InitSP(maxHeight int, height int, capacity int) SkipList {
	s.maxHeight = maxHeight
	s.height = height
	s.size = 0
	s.head = &SkipListNode{next: make([]*SkipListNode, height+1)}
	s.max_capacity = capacity
	return *s
}

func (s *SkipList) Roll() int {
	level := 0
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

func (s *SkipList) Search(searchKey string) *SkipListNode { //vraca vrednost koja odgovara kljucu(ako postoji), ako ne postoji vraca nil
	x := s.head
	var i int
	for i = s.height; i >= 0; i-- {
		if x.podatak.Key == searchKey && x.podatak.Tombstone != 1 {
			return x
		}
		for x.next[i] != nil && x.next[i].podatak.Key <= searchKey {
			x = x.next[i]
		}
	}
	return nil
}

func (s *SkipList) Put(podatak Importi.Podatak) []Importi.Podatak {
	if s.Search(podatak.Key) != nil { //ako podatak s tim kljucem vec postoji samo ga izmeni
		x := s.head
		var i int
		for i = s.height; i >= 0; i-- {
			if x.podatak.Key == podatak.Key {
				x.podatak.Value = podatak.Value
				x.podatak.Timestamp = podatak.Timestamp
				x.podatak.Tombstone = podatak.Tombstone
				return nil
			}
			for x.next[i] != nil && x.next[i].podatak.Key <= podatak.Key {
				x = x.next[i]
			}
		}
	}
	//inace dodaj
	var pod Importi.Podatak
	pod = podatak
	var noviNode SkipListNode
	noviNode.InitSP(pod, s.Roll())
	x := s.head
	var i int

	for i = len(noviNode.next) - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].podatak.Key < noviNode.podatak.Key {
			x = x.next[i]
		}

		if x.next[i] != nil {
			noviNode.next[i] = x.next[i]
		}

		x.next[i] = &noviNode
	}
	s.size += 1
	fmt.Println("Uspesno ubaceno!!!")

	if s.size >= s.max_capacity {
		fmt.Println("Skip lista je popunjena!")
		//ovde treba sortirati sve cvorove po kljucu i vratiti sortiranu listu cvorova
		//cvorovi u skip listi su po difoltu sortirani tako da samo prolazim kroz najnizi nivo
		sortNodeovi := []Importi.Podatak{}
		x := s.head
		for x != nil {
			sortNodeovi = append(sortNodeovi, x.podatak)
			x = x.next[0]
		}
		//treba isprazniti listu kada se ona popuni
		s.InitSP(s.maxHeight, s.height, s.max_capacity)
		return sortNodeovi
	}

	return nil
}

func (s *SkipList) GetAllData(podatak Importi.Podatak) []Importi.Podatak {
	sortNodeovi := []Importi.Podatak{}
	x := s.head
	for x != nil {
		sortNodeovi = append(sortNodeovi, x.podatak)
		x = x.next[0]
	}
	//treba isprazniti listu kada se ona popuni
	s.InitSP(s.maxHeight, s.height, s.max_capacity)
	return sortNodeovi
}
