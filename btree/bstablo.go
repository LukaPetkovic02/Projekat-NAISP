package main

import (
	"fmt"
	"time"
)

type Stablo struct {
	max  int
	head *Node
}

type Podatak struct {
	Key       string
	Value     []byte
	Tombstone byte
	Timestamp int64
}

func (data *Podatak) PrintData() {
	fmt.Print("Key : " + data.Key + " ; Value : " + string(data.Value) + " | Timestamp : ")
	fmt.Print(data.Timestamp)
	fmt.Print(" | Tombstone : ")
	fmt.Println(data.Tombstone)
}

type Node struct {
	podaci   []Podatak
	children []*Node
	parent   *Node
}

func (p *Podatak) InitPrazanPodatak() Podatak {
	p.Key = ""
	p.Value = []byte("")
	p.Tombstone = 0
	p.Timestamp = time.Now().Unix()
	return *p
}

func (s *Node) InitSP(max int, parent *Node) Node {
	//s.key = make([]string, max+1)
	//s.value = make([][]byte, max+1)

	s.podaci = make([]Podatak, max+1)
	s.podaci[0] = Podatak{Key: "", Value: nil, Tombstone: 1, Timestamp: time.Now().Unix()} //inicijalizujemo prazan podatak u pocetku
	s.children = make([]*Node, max+2)                                                      //jer ce kod podele cvorova u jednom momentu biti vise dece
	s.parent = parent
	//s.key[0] = key
	//s.value[0] = value

	s.parent.podaci = make([]Podatak, max+1)
	//s.parent.value = make([][]byte, max+1)
	s.parent.children = make([]*Node, max+2)
	s.parent.children[0] = s
	return *s
}

func (st *Stablo) InitSP(max int) Stablo {
	nilRoditelj := new(Node)
	st.head = new(Node)
	st.head.InitSP(max, nilRoditelj)
	//st.head.parent = nilRoditelj
	//st.head.children = make([]*Node, max+2)
	//st.head.InitSP(key, value, max, nilRoditelj)
	st.max = max
	return *st
}

func (s *Stablo) search(searchKey string) (*Node, int) {
	x := s.head
	t := true
	var i int
	for x != nil {
		t = true
		for i = 0; i < len(x.podaci)-1 && x.podaci[i].Key != ""; i++ {
			if searchKey == x.podaci[i].Key && x.podaci[i].Tombstone == 0 {
				return x, i
			} else if x.podaci[i].Key > searchKey {
				x = x.children[i]
				t = false
				break
			}
		}
		if t {
			x = x.children[i]
		}
	}
	return x, 0
}

func brEl(podaci []Podatak) int {
	var i int

	for i = 0; i < len(podaci) && podaci[i].Key != ""; i++ {
	}
	return i
}
func podeliCvor(x *Node, max int) (*Node, *Node, *Node) {
	var sredina int = max / 2
	x1 := new(Node) //pre sredine
	x1.podaci = make([]Podatak, max+1)
	for i, p := range x1.podaci {
		p.InitPrazanPodatak()
		x1.podaci[i] = p
	}
	//x1.podaci.Key = make([]string, max+1)
	//x1.value = make([][]byte, max+1)
	x1.children = make([]*Node, max+2)
	//deca od x1 su prvih len/2 dece x
	for f := 0; f < len(x.children)/2; f++ {
		x1.children[f] = x.children[f]
	}
	//x1.parent = x.parent
	x2 := new(Node) //sredina
	x2.podaci = make([]Podatak, max+1)
	for i, p := range x2.podaci {
		p.InitPrazanPodatak()
		x2.podaci[i] = p
	}
	x2.children = make([]*Node, max+2)
	//x2 ni ne treba da ima decu jer on svakako ide gore
	//x2.parent = x.parent
	x3 := new(Node) //posle sredine
	x3.podaci = make([]Podatak, max+1)
	for i, p := range x3.podaci {
		p.InitPrazanPodatak()
		x3.podaci[i] = p
	}
	x3.children = make([]*Node, max+2)
	//zadnjih pola dece xa
	for f := len(x.children) / 2; f < len(x.children); f++ {
		x3.children[f-len(x.children)/2] = x.children[f]
	}
	//x3.parent = x.parent
	for i := 0; i < sredina; i++ {
		x1.podaci[i] = x.podaci[i]
		//x1.key[i] = x.key[i]
		//x1.value[i] = x.value[i]
	}
	//x2.key[0] = x.key[sredina]
	//x2.value[0] = x.value[sredina]
	x2.podaci[0] = x.podaci[sredina]
	for i := sredina + 1; i < brEl(x.podaci); i++ {
		//x3.key[i-sredina-1] = x.key[i]
		//x3.value[i-sredina-1] = x.value[i]
		x3.podaci[i-sredina-1] = x.podaci[i]
	}
	return x1, x2, x3
}

func (s *Stablo) put(addKey string, addValue []byte, vreme time.Time) {
	if s.head.podaci[0].Key == "" { //ako je prazno stablo
		s.head.podaci[0].Key = addKey
		s.head.podaci[0].Value = addValue
		s.head.podaci[0].Tombstone = 0
		s.head.podaci[0].Timestamp = vreme.Unix()
		return
	}

	a, _ := s.search(addKey)

	if a != nil { //ako element vec postoji, samo cemo mu promeniti value i timestamp!!!!
		x := s.head
		t := true
		var i int
		for x != nil {
			t = true
			for i = 0; i < len(x.podaci)-1 && x.podaci[i].Key != ""; i++ {
				if addKey == x.podaci[i].Key && x.podaci[i].Tombstone == 0 {
					x.podaci[i].Value = addValue
					x.podaci[i].Timestamp = vreme.Unix()
					return
				} else if x.podaci[i].Key > addKey {
					x = x.children[i]
					t = false
					break
				}
			}
			if t {
				x = x.children[i]
			}
		}
	}

	x := s.head
	t := true
	var i int
	for x.children[0] != nil { //ne moze x.children!=nil;;ako mu je prvo dete nil znaci da nema dece
		for i = 0; i < len(x.podaci)-1 && x.podaci[i].Key != ""; i++ {
			if x.podaci[i].Key > addKey {
				x = x.children[i]
				t = false
				break
			}
		}

		if t {
			x = x.children[i]
		}
		t = true
	}
	temps := addKey
	tempb := addValue
	var tempt byte
	tempt = 0
	tempv := vreme.Unix()
	// kad nadjemo
	for i = 0; i < brEl(x.podaci); i++ {
		if temps < x.podaci[i].Key {
			x.podaci[i].Key, temps = temps, x.podaci[i].Key
			x.podaci[i].Value, tempb = tempb, x.podaci[i].Value
			x.podaci[i].Tombstone, tempt = tempt, x.podaci[i].Tombstone
			x.podaci[i].Timestamp, tempv = tempv, x.podaci[i].Timestamp
		}
	}
	x.podaci[i].Key = temps
	x.podaci[i].Value = tempb
	x.podaci[i].Tombstone = tempt
	x.podaci[i].Timestamp = tempv

	//ako je doslo do overflowa radimo dodatno
	if brEl(x.podaci) == len(x.podaci) {
		//if t { //ako ne postoji sibling koji nije popunjen
		//uradi podelu cvorova
		for brEl(x.podaci) == s.max+1 { //dok je trenutni nivo popunjen
			x1, x2, x3 := podeliCvor(x, s.max)
			x.podaci = make([]Podatak, 0)
			for a := 0; a < len(x1.podaci) && x1.podaci[a].Key != ""; a++ {
				x.podaci = append(x.podaci, x1.podaci[a])
			}
			for a := 0; a < len(x3.podaci) && x3.podaci[a].Key != ""; a++ {
				x.podaci = append(x.podaci, x3.podaci[a])
			}
			var p Podatak
			p.InitPrazanPodatak()
			x.podaci = append(x.podaci, p)

			tempk := x2.podaci[0].Key //srednji ima samo jedan key
			tempv := x2.podaci[0].Value
			tempt := x2.podaci[0].Tombstone
			tempvr := x2.podaci[0].Timestamp
			j := 0
			for j = 0; j < brEl(x.parent.podaci); j++ { //srednji kljuc dajemo roditelju na odgovarajuce mesto
				if tempk < x.parent.podaci[j].Key {
					x.parent.podaci[j].Key, tempk = tempk, x.parent.podaci[j].Key
					x.parent.podaci[j].Value, tempv = tempv, x.parent.podaci[j].Value
					x.parent.podaci[j].Tombstone, tempt = tempt, x.parent.podaci[j].Tombstone
					x.parent.podaci[j].Timestamp, tempvr = tempvr, x.parent.podaci[j].Timestamp
				}

			}
			x.parent.podaci[j].Key = tempk
			x.parent.podaci[j].Value = tempv
			x.parent.podaci[j].Tombstone = tempt
			x.parent.podaci[j].Timestamp = tempvr
			//fmt.Println(x1.parent)
			//fmt.Println("roditelj nakon sredjivanja:", x.parent)
			x1.parent = x.parent
			//x2.parent = x.parent
			x3.parent = x.parent //tek nakon sto sredimo parenta
			//fmt.Println(x.parent)
			//treba pomeriti decu od kraja do j za jedno udesno
			k := 0
			for k = len(x.parent.children) - 1; k > j; k-- {
				x.parent.children[k] = x.parent.children[k-1]
			}
			x.parent.children[k+1] = x3
			x.parent.children[k] = x1

			x = x.parent
		}
		//na kraju postavljamo s.head=x ako se koren rasformirao
		isti := true
		if brEl(x.podaci) != brEl(s.head.parent.podaci) {
			isti = false
		} else {
			for o := 0; o < brEl(x.podaci); o++ {
				if x.podaci[o].Key != s.head.parent.podaci[o].Key {
					isti = false
				}
			}
		}
		if isti {
			//fmt.Println("rasformirao se koren kod ubacivanja", addKey)
			s.head = x //samo ako je x jednak roditelju heada
			//ovde treba inicijalizovati prazan parent
			s.head.parent = new(Node)
			var p Podatak
			p.InitPrazanPodatak()
			s.head.parent.podaci = make([]Podatak, s.max+1)
			s.head.parent.podaci[0] = p
			//s.head.parent.podaci = append(s.head.parent.podaci, )
			//s.head.parent.key = make([]string, s.max+1)
			//s.head.parent.value = make([][]byte, s.max+1)
			s.head.parent.children = make([]*Node, s.max+2)
			s.head.parent.children[0] = x
		}

	}
	srediRoditelje(s.head)
}

func (s *Stablo) delete(deleteKey string, vreme time.Time) { //logicko brisanje
	a, _ := s.search(deleteKey)

	if a != nil { //ako element postoji, samo cemo mu promeniti timestamp!
		x := s.head
		t := true
		var i int
		for x != nil {
			t = true
			for i = 0; i < len(x.podaci)-1 && x.podaci[i].Key != ""; i++ {
				if deleteKey == x.podaci[i].Key && x.podaci[i].Tombstone == 0 {
					x.podaci[i].Tombstone = 1
					x.podaci[i].Timestamp = vreme.Unix()
					return
				} else if x.podaci[i].Key > deleteKey {
					x = x.children[i]
					t = false
					break
				}
			}
			if t {
				x = x.children[i]
			}
		}
	} else {
		fmt.Println("Element sa tim kljucem ne postoji!")
		return
	}
}
func srediRoditelje(x *Node) {
	if x != nil {
		for i := 0; i < len(x.children); i++ {
			if x.children[i] != nil {
				x.children[i].parent = x
				srediRoditelje(x.children[i])
			}

		}

	}
}
func ispis(x *Node, nivo int) {
	if x != nil {
		fmt.Print("nivo ", nivo, ":")
		for j := 0; j < brEl(x.podaci); j++ {
			x.podaci[j].PrintData()
		}
		fmt.Println()
		nivo++
		if x.children[0] != nil { //ako ima decu
			for i := 0; i < len(x.children) && x.children[i] != nil; i++ {
				ispis(x.children[i], nivo)
			}
		}
	}
}
func (s *Stablo) AllDataSorted() []Podatak { //vraca listu sortiranih podataka
	podaci := make([]Podatak, 0)
	s.allDataSorted(s.head, &podaci)
	return podaci
}
func (s *Stablo) allDataSorted(n *Node, podaci *[]Podatak) {
	if n == nil {
		return
	}
	for i := 0; i < brEl(n.podaci); i++ {
		if n.children[0] != nil {
			s.allDataSorted(n.children[i], podaci)
		}
		*podaci = append(*podaci, n.podaci[i])
	}
	if n.children[0] != nil {
		s.allDataSorted(n.children[brEl(n.podaci)], podaci)
	}
}
func main() {
	var s Stablo
	s.InitSP(3)
	fmt.Println(s)
	s.put("a", []byte("nesto"), time.Now())
	s.put("b", []byte("nesto"), time.Now())
	//cvor, i := s.search("4")
	//fmt.Println(cvor, i)
	s.put("c", []byte("1estodrugo"), time.Now())
	s.put("d", []byte("2estodrugo"), time.Now())
	s.put("e", []byte("3estodrugo"), time.Now())
	s.put("f", []byte("4estodrugo"), time.Now())
	s.put("g", []byte("5estodrugo"), time.Now())
	s.put("h", []byte("6estodrugo"), time.Now())
	s.put("i", []byte("7estodrugo"), time.Now())
	s.put("j", []byte("1estodrugo"), time.Now())
	s.put("k", []byte("2estodrugo"), time.Now())
	s.put("l", []byte("3estodrugo"), time.Now())
	s.put("m", []byte("2estodrugo"), time.Now())
	s.put("n", []byte("3estodrugo"), time.Now())
	s.put("o", []byte("4estodrugo"), time.Now())
	s.put("p", []byte("5estodrugo"), time.Now())
	s.put("q", []byte("6estodrugo"), time.Now())
	s.put("r", []byte("7estodrugo"), time.Now())
	s.put("s", []byte("1estodrugo"), time.Now())
	s.put("t", []byte("2estodrugo"), time.Now())
	s.put("u", []byte("3estodrugo"), time.Now())
	s.put("v", []byte("4estodrugo"), time.Now())
	s.put("w", []byte("5estodrugo"), time.Now())
	s.put("x", []byte("6estodrugo"), time.Now())
	s.put("y", []byte("7estodrugo"), time.Now())
	s.put("z", []byte("7estodrugo"), time.Now())

	fmt.Println("PRE")
	s.put("a", []byte("promenjeno"), time.Now())
	s.delete("a", time.Now())
	sortirani := s.AllDataSorted()
	for i := 0; i < len(sortirani); i++ {
		sortirani[i].PrintData()
	}
	//ispis(s.head, 0)
	//v1, _ := s.search("a")
	//v2, _ := s.search("b")
	//fmt.Println(v1)
	//fmt.Println(v2)
}
