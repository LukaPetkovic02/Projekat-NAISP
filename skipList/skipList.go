package skipList

import (
	"fmt"
	"math/rand"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type SkipList struct {
	maxHeight int
	height    int
	size      int
	//max_capacity int
	head *SkipListNode
}

type SkipListNode struct {
	podatak types.Record
	next    []*SkipListNode
}

func (s *SkipListNode) InitSP(podatak types.Record, level int) SkipListNode {
	//s.key = key
	//s.value = value
	s.next = make([]*SkipListNode, level+1)
	//s.tombstone = tomb
	//s.timestamp = vreme.Unix()
	s.podatak = podatak
	return *s
}

func (s *SkipList) InitSP(maxHeight int, height int) SkipList {
	s.maxHeight = maxHeight
	s.height = height
	s.size = 0
	s.head = &SkipListNode{next: make([]*SkipListNode, height+1)}
	//s.max_capacity = capacity
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
		if x.podatak.Key == searchKey && x.podatak.Tombstone == false {
			return x
		}
		for x.next[i] != nil && x.next[i].podatak.Key <= searchKey {
			x = x.next[i]
		}
	}
	return nil
}

func (s *SkipList) Get(searchKey string) *types.Record { //vraca vrednost koja odgovara kljucu(ako postoji), ako ne postoji vraca nil
	x := s.head
	var i int
	for i = s.height; i >= 0; i-- {
		if x.podatak.Key == searchKey {
			return &x.podatak
		}
		for x.next[i] != nil && x.next[i].podatak.Key <= searchKey {
			x = x.next[i]
		}
	}
	return nil
}

func (s *SkipList) GetSortedRecordsList() []types.Record {
	//ovde treba sortirati sve cvorove po kljucu i vratiti sortiranu listu cvorova
	//cvorovi u skip listi su po difoltu sortirani tako da samo prolazim kroz najnizi nivo
	sortNodeovi := []types.Record{}
	x := s.head
	for x != nil {
		sortNodeovi = append(sortNodeovi, x.podatak)
		x = x.next[0]
	}
	//treba isprazniti listu kada se ona popuni
	s.InitSP(s.maxHeight, s.height)
	return sortNodeovi[1:]
}

func (s *SkipList) Add(podatak types.Record) bool {
	if s.search(podatak.Key) != nil { //ako podatak s tim kljucem vec postoji samo ga izmeni
		x := s.head
		var i int
		for i = s.height; i >= 0; i-- {
			if x.podatak.Key == podatak.Key {
				x.podatak.Value = podatak.Value
				x.podatak.Timestamp = podatak.Timestamp
				x.podatak.Tombstone = podatak.Tombstone
				x.podatak.CRC = podatak.CRC
				x.podatak.KeySize = podatak.KeySize
				x.podatak.ValueSize = podatak.ValueSize
				return true
			}
			for x.next[i] != nil && x.next[i].podatak.Key <= podatak.Key {
				x = x.next[i]
			}
		}
	}

	//if s.size >= s.max_capacity {
	//	return false
	//}
	//inace dodaj
	var pod types.Record
	pod = podatak
	var noviNode SkipListNode
	fmt.Println(s.height)
	//visinapre:=s.height
	noviNode.InitSP(pod, s.roll())
	//visinaposle:=s.height

	x := s.head
	var i int
	fmt.Println(s.height)
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

	return true
}

func (s *SkipList) Delete(key string) bool {
	a := s.search(key)
	if a == nil {
		return false
	} else {
		x := s.head
		var i int
		for i = s.height; i >= 0; i-- {
			if x.podatak.Key == key && x.podatak.Tombstone == false {
				x.podatak.Tombstone = true
				return true
			}
			for x.next[i] != nil && x.next[i].podatak.Key <= key {
				x = x.next[i]
			}
		}
	}
	return true
}

func (s *SkipList) Clear() {
	s.InitSP(s.maxHeight, s.height)
}

func (s *SkipList) GetSize() int {
	return s.size
}
