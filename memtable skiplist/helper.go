package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ZA SAD SAM SAMO DEO SA SKIP LISTOM URADIO, PROVERITE
type SkipList struct {
	maxHeight    int
	height       int
	size         int
	max_capacity int
	head         *SkipListNode
}

type SkipListNode struct {
	key       string
	value     []byte
	next      []*SkipListNode
	tombstone bool
	timestamp time.Time //vreme poslednje promene
}

func (s *SkipListNode) InitSP(key string, value []byte, level int, tomb bool, vreme time.Time) SkipListNode {
	s.key = key
	s.value = value
	s.next = make([]*SkipListNode, level+1)
	s.tombstone = tomb
	s.timestamp = vreme
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
		if x.key == searchKey && !x.tombstone {
			return x
		}
		for x.next[i] != nil && x.next[i].key <= searchKey {
			x = x.next[i]
		}
	}
	return nil
}

func (s *SkipList) insert(novi string, vrednost []byte, vreme time.Time) []SkipListNode { //dodaje novi element i vraca sortiranu listu ako je skip lista puna, inace vraca nil
	if s.search(novi) != nil {
		return nil
	}

	var noviNode SkipListNode
	noviNode.InitSP(novi, vrednost, s.roll(), false, vreme)
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
	s.size += 1
	fmt.Println("Uspesno ubaceno!!!")

	if s.size >= s.max_capacity {
		fmt.Println("Skip lista je popunjena!")
		//ovde treba sortirati sve cvorove po kljucu i vratiti sortiranu listu cvorova
		//cvorovi u skip listi su po difoltu sortirani tako da samo prolazim kroz najnizi nivo
		sortNodeovi := []SkipListNode{}
		x := s.head
		for x != nil {
			sortNodeovi = append(sortNodeovi, *x)
			x = x.next[0]
		}
		//treba isprazniti listu kada se ona popuni
		s.InitSP(10, 5, 4)
		return sortNodeovi
	}

	return nil
}

func (s *SkipList) put(key string, newValue []byte, vreme time.Time) { //menja value cvora sa datim kljucem i updateuje vreme
	if s.search(key) == nil {
		//fmt.Println("Element sa tim kljucem ne postoji!!!")
		return
	}

	x := s.head
	var i int
	for i = s.height; i >= 0; i-- {
		if x.key == key && !x.tombstone {
			x.value = newValue
			x.timestamp = vreme
		}
		for x.next[i] != nil && x.next[i].key <= key {
			x = x.next[i]
		}
	}
}

func (s *SkipList) delete(brisanje string, vreme time.Time) { //logicko brisanje (postavljamo tombstone na true) i postavljamo vreme
	p := s.search(brisanje)
	if p == nil {
		fmt.Println("Nije pronadjen taj string.")
		return
	}

	x := s.head
	var i int
	for i = s.height; i >= 0; i-- {
		if x.key == brisanje && !x.tombstone {
			x.tombstone = true
			x.timestamp = vreme
		}
		for x.next[i] != nil && x.next[i].key <= brisanje {
			x = x.next[i]
		}
	}

	/*x := s.head
	var i int
	for i = len(p.next) - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < brisanje {
			x = x.next[i]
		}

		x.next[i] = x.next[i].next[i]

	}
	fmt.Println("Uspesno obrisano!")*/

}
func main() {
	var b SkipList
	b.InitSP(10, 5, 4) //ovo je samo primer, posle cemo ubaciti neke konstante koje cemo citati iz nekog spoljasnjeg fajla
	//fmt.Println(b.size)
	b.insert("12", []byte("poruka"), time.Now())
	b.insert("14", []byte("poruka"), time.Now())
	b.insert("13", []byte("poruka"), time.Now())
	f1 := b.insert("16", []byte("poruka"), time.Now())
	fmt.Println(f1)
	//nakon ovoga se isprazni skip lista jer se popunila
	f2 := b.insert("15", []byte("poruka"), time.Now())
	fmt.Println(f2)
	//fmt.Println(b.size)
	//b.search("asd")

	//fmt.Println(b.search("12").value)
	//b.insert("asd", []byte("nebitno"), time.Now())
	//b.delete("12", time.Now())
	//b.delete("13", time.Now())
	//fmt.Println(b.search("12"))
}
