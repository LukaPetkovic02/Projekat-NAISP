package btree

import "fmt"

//kasnije treba import za podatak
type Podatak struct {
	key       string
	value     []byte
	tombstone byte
	timestamp int64
}

type Stablo struct {
	max  int
	head *Node
}

type Node struct {
	key      []string
	value    []Podatak
	children []*Node
	parent   *Node
}

//inicijalzuje Node
func (s *Node) InitSP(data Podatak, max int, parent *Node) Node {
	s.key = make([]string, max+1)     //s key je lista kljuceva cvorova
	s.value = make([]Podatak, max+1)  //s value je lista vrednosti cvorova
	s.children = make([]*Node, max+2) //s children je lista dece koja ce biti za jedan veca od liste cvorova
	s.parent = parent                 //postavljamo mu parenta
	s.key[0] = data.key               // kljuc svake vrednosti Podatak je vrednost njenog kljuca
	s.value[0] = data

	s.parent.key = make([]string, max+1)
	s.parent.value = make([]Podatak, max+1)
	s.parent.children = make([]*Node, max+2)
	s.parent.children[0] = s
	return *s
}

//inicijalizuje stablo
func (st *Stablo) InitSP(data Podatak, max int) Stablo {
	nilRoditelj := new(Node) //stablo u sebi sadrzi pocetni node kao i maximalni br cvorova u nodu
	st.head = new(Node)      //st.head je parent pocetnog cvora koji je prazan element
	st.head.InitSP(data, max, nilRoditelj)
	st.max = max
	return *st
}

func (s *Stablo) Search(searchKey string) (*Node, int) {
	x := s.head
	t := true
	var i int
	for x != nil {
		t = true
		for i = 0; i < len(x.key)-1 && x.key[i] != ""; i++ {
			if searchKey == x.key[i] {
				return x, i
			} else if x.key[i] > searchKey {
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

func BrEl(key []string) int {
	var i int

	for i = 0; i < len(key) && key[i] != ""; i++ {
	}
	return i
}

func PodeliCvor(x *Node, max int) (*Node, *Node, *Node) {
	var sredina int = max / 2
	x1 := new(Node) //pre sredine
	x1.key = make([]string, max+1)
	x1.value = make([]Podatak, max+1)
	x1.children = make([]*Node, max+2)
	//deca od x1 su prvih len/2 dece x
	for f := 0; f < len(x.children)/2; f++ {
		x1.children[f] = x.children[f]
	}
	//x1.parent = x.parent
	x2 := new(Node) //sredina
	x2.key = make([]string, max+1)
	x2.value = make([]Podatak, max+1)
	x2.children = make([]*Node, max+2)
	//x2 ni ne treba da ima decu jer on svakako ide gore
	//x2.parent = x.parent
	x3 := new(Node) //posle sredine
	x3.key = make([]string, max+1)
	x3.value = make([]Podatak, max+1)
	x3.children = make([]*Node, max+2)
	//zadnjih pola dece xa
	for f := len(x.children) / 2; f < len(x.children); f++ {
		x3.children[f-len(x.children)/2] = x.children[f]
	}
	//x3.parent = x.parent
	for i := 0; i < sredina; i++ {
		x1.key[i] = x.key[i]
		x1.value[i] = x.value[i]
	}
	x2.key[0] = x.key[sredina]
	x2.value[0] = x.value[sredina]
	for i := sredina + 1; i < BrEl(x.key); i++ {
		x3.key[i-sredina-1] = x.key[i]
		x3.value[i-sredina-1] = x.value[i]
	}
	return x1, x2, x3
}

func (s *Stablo) Add(data Podatak) {
	a, _ := s.Search(data.key)
	if a != nil {
		fmt.Println("Element vec postoji sa tim kljucem.")
		return
	}

	x := s.head
	t := true
	var i int
	for x.children[0] != nil { //ne moze x.children!=nil;;ako mu je prvo dete nil znaci da nema dece
		for i = 0; i < len(x.key)-1 && x.key[i] != ""; i++ {
			if x.key[i] > data.key {
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
	temps := data.key
	tempb := data
	// kad nadjemo
	for i = 0; i < BrEl(x.key); i++ {
		if temps < x.key[i] {
			x.key[i], temps = temps, x.key[i]
			x.value[i], tempb = tempb, x.value[i]
		}
	}
	x.key[i] = temps
	x.value[i] = tempb

	//ako je doslo do overflowa radimo dodatno
	if BrEl(x.key) == len(x.key) {
		//if t { //ako ne postoji sibling koji nije popunjen
		//uradi podelu cvorova
		for BrEl(x.key) == s.max+1 { //dok je trenutni nivo popunjen
			x1, x2, x3 := PodeliCvor(x, s.max)
			x.key = make([]string, 0)
			for a := 0; a < len(x1.key) && x1.key[a] != ""; a++ {
				x.key = append(x.key, x1.key[a])
			}
			for a := 0; a < len(x3.key) && x3.key[a] != ""; a++ {
				x.key = append(x.key, x3.key[a])
			}
			x.key = append(x.key, "")

			tempk := x2.key[0] //srednji ima samo jedan key
			tempv := x2.value[0]
			j := 0
			for j = 0; j < BrEl(x.parent.key); j++ { //srednji kljuc dajemo roditelju na odgovarajuce mesto
				if tempk < x.parent.key[j] {
					x.parent.key[j], tempk = tempk, x.parent.key[j]
					x.parent.value[j], tempv = tempv, x.parent.value[j]
				}

			}
			x.parent.key[j] = tempk
			x.parent.value[j] = tempv
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
		if BrEl(x.key) != BrEl(s.head.parent.key) {
			isti = false
		} else {
			for o := 0; o < BrEl(x.key); o++ {
				if x.key[o] != s.head.parent.key[o] {
					isti = false
				}
			}
		}
		if isti {
			//fmt.Println("rasformirao se koren kod ubacivanja", addKey)
			s.head = x //samo ako je x jednak roditelju heada
			//ovde treba inicijalizovati prazan parent
			s.head.parent = new(Node)
			s.head.parent.key = make([]string, s.max+1)
			s.head.parent.value = make([]Podatak, s.max+1)
			s.head.parent.children = make([]*Node, s.max+2)
			s.head.parent.children[0] = x
		}

	}
	SrediRoditelje(s.head)
}

func SrediRoditelje(x *Node) {
	if x != nil {
		for i := 0; i < len(x.children); i++ {
			if x.children[i] != nil {
				x.children[i].parent = x
				SrediRoditelje(x.children[i])
			}

		}

	}
}

func Ispis(x *Node, nivo int) {
	if x != nil {
		fmt.Print("nivo ", nivo, ":")
		for j := 0; j < BrEl(x.key); j++ {
			fmt.Print(x.key[j], " ")
		}
		fmt.Println()
		nivo++
		if x.children[0] != nil { //ako ima decu
			for i := 0; i < len(x.children) && x.children[i] != nil; i++ {
				Ispis(x.children[i], nivo)
			}
		}
	}
}
