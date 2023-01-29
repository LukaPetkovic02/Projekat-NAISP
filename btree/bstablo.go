package btree

import (
	"fmt"
	importi "projekat/utils"
)

//kasnije treba import za podatak

type Stablo struct {
	Max  int
	Head *Node
}

type Node struct {
	Key      []string
	Value    []importi.Podatak
	Children []*Node
	Parent   *Node
}

// inicijalzuje Node
func (s *Node) InitSP(data importi.Podatak, max int, parent *Node) Node {
	s.Key = make([]string, max+1)            //s key je lista kljuceva cvorova
	s.Value = make([]importi.Podatak, max+1) //s value je lista vrednosti cvorova
	s.Children = make([]*Node, max+2)        //s children je lista dece koja ce biti za jedan veca od liste cvorova
	s.Parent = parent                        //postavljamo mu parenta
	s.Key[0] = data.Key                      // kljuc svake vrednosti Podatak je vrednost njenog kljuca
	s.Value[0] = data

	s.Parent.Key = make([]string, max+1)
	s.Parent.Value = make([]importi.Podatak, max+1)
	s.Parent.Children = make([]*Node, max+2)
	s.Parent.Children[0] = s
	return *s
}

// inicijalizuje stablo
func (st *Stablo) InitSP(data importi.Podatak, max int) Stablo {
	nilRoditelj := new(Node) //stablo u sebi sadrzi pocetni node kao i maximalni br cvorova u nodu
	st.Head = new(Node)      //st.head je parent pocetnog cvora koji je prazan element
	st.Head.InitSP(data, max, nilRoditelj)
	st.Max = max
	return *st
}

func (s *Stablo) Search(searchKey string) (*Node, int) {
	x := s.Head
	t := true
	var i int
	for x != nil {
		t = true
		for i = 0; i < len(x.Key)-1 && x.Key[i] != ""; i++ {
			if searchKey == x.Key[i] {
				return x, i
			} else if x.Key[i] > searchKey {
				x = x.Children[i]
				t = false
				break
			}
		}
		if t {
			x = x.Children[i]
		}
	}
	return x, 0
}

func BrEl(Key []string) int {
	var i int

	for i = 0; i < len(Key) && Key[i] != ""; i++ {
	}
	return i
}

func PodeliCvor(x *Node, max int) (*Node, *Node, *Node) {
	var sredina int = max / 2
	x1 := new(Node) //pre sredine
	x1.Key = make([]string, max+1)
	x1.Value = make([]importi.Podatak, max+1)
	x1.Children = make([]*Node, max+2)
	//deca od x1 su prvih len/2 dece x
	for f := 0; f < len(x.Children)/2; f++ {
		x1.Children[f] = x.Children[f]
	}
	//x1.parent = x.parent
	x2 := new(Node) //sredina
	x2.Key = make([]string, max+1)
	x2.Value = make([]importi.Podatak, max+1)
	x2.Children = make([]*Node, max+2)
	//x2 ni ne treba da ima decu jer on svakako ide gore
	//x2.parent = x.parent
	x3 := new(Node) //posle sredine
	x3.Key = make([]string, max+1)
	x3.Value = make([]importi.Podatak, max+1)
	x3.Children = make([]*Node, max+2)
	//zadnjih pola dece xa
	for f := len(x.Children) / 2; f < len(x.Children); f++ {
		x3.Children[f-len(x.Children)/2] = x.Children[f]
	}
	//x3.parent = x.parent
	for i := 0; i < sredina; i++ {
		x1.Key[i] = x.Key[i]
		x1.Value[i] = x.Value[i]
	}
	x2.Key[0] = x.Key[sredina]
	x2.Value[0] = x.Value[sredina]
	for i := sredina + 1; i < BrEl(x.Key); i++ {
		x3.Key[i-sredina-1] = x.Key[i]
		x3.Value[i-sredina-1] = x.Value[i]
	}
	return x1, x2, x3
}

func (s *Stablo) Add(data importi.Podatak) {
	a, _ := s.Search(data.Key)
	if a != nil {
		fmt.Println("Element vec postoji sa tim kljucem.")
		return
	}

	x := s.Head
	t := true
	var i int
	for x.Children[0] != nil { //ne moze x.children!=nil;;ako mu je prvo dete nil znaci da nema dece
		for i = 0; i < len(x.Key)-1 && x.Key[i] != ""; i++ {
			if x.Key[i] > data.Key {
				x = x.Children[i]
				t = false
				break
			}
		}

		if t {
			x = x.Children[i]
		}
		t = true
	}
	temps := data.Key
	tempb := data
	// kad nadjemo
	for i = 0; i < BrEl(x.Key); i++ {
		if temps < x.Key[i] {
			x.Key[i], temps = temps, x.Key[i]
			x.Value[i], tempb = tempb, x.Value[i]
		}
	}
	x.Key[i] = temps
	x.Value[i] = tempb

	//ako je doslo do overflowa radimo dodatno
	if BrEl(x.Key) == len(x.Key) {
		//if t { //ako ne postoji sibling koji nije popunjen
		//uradi podelu cvorova
		for BrEl(x.Key) == s.Max+1 { //dok je trenutni nivo popunjen
			x1, x2, x3 := PodeliCvor(x, s.Max)
			x.Key = make([]string, 0)
			for a := 0; a < len(x1.Key) && x1.Key[a] != ""; a++ {
				x.Key = append(x.Key, x1.Key[a])
			}
			for a := 0; a < len(x3.Key) && x3.Key[a] != ""; a++ {
				x.Key = append(x.Key, x3.Key[a])
			}
			x.Key = append(x.Key, "")

			tempk := x2.Key[0] //srednji ima samo jedan key
			tempv := x2.Value[0]
			j := 0
			for j = 0; j < BrEl(x.Parent.Key); j++ { //srednji kljuc dajemo roditelju na odgovarajuce mesto
				if tempk < x.Parent.Key[j] {
					x.Parent.Key[j], tempk = tempk, x.Parent.Key[j]
					x.Parent.Value[j], tempv = tempv, x.Parent.Value[j]
				}

			}
			x.Parent.Key[j] = tempk
			x.Parent.Value[j] = tempv
			//fmt.Println(x1.parent)
			//fmt.Println("roditelj nakon sredjivanja:", x.parent)
			x1.Parent = x.Parent
			//x2.parent = x.parent
			x3.Parent = x.Parent //tek nakon sto sredimo parenta
			//fmt.Println(x.parent)
			//treba pomeriti decu od kraja do j za jedno udesno
			k := 0
			for k = len(x.Parent.Children) - 1; k > j; k-- {
				x.Parent.Children[k] = x.Parent.Children[k-1]
			}
			x.Parent.Children[k+1] = x3
			x.Parent.Children[k] = x1

			x = x.Parent
		}
		//na kraju postavljamo s.head=x ako se koren rasformirao
		isti := true
		if BrEl(x.Key) != BrEl(s.Head.Parent.Key) {
			isti = false
		} else {
			for o := 0; o < BrEl(x.Key); o++ {
				if x.Key[o] != s.Head.Parent.Key[o] {
					isti = false
				}
			}
		}
		if isti {
			//fmt.Println("rasformirao se koren kod ubacivanja", addKey)
			s.Head = x //samo ako je x jednak roditelju heada
			//ovde treba inicijalizovati prazan parent
			s.Head.Parent = new(Node)
			s.Head.Parent.Key = make([]string, s.Max+1)
			s.Head.Parent.Value = make([]importi.Podatak, s.Max+1)
			s.Head.Parent.Children = make([]*Node, s.Max+2)
			s.Head.Parent.Children[0] = x
		}

	}
	SrediRoditelje(s.Head)
}

func SrediRoditelje(x *Node) {
	if x != nil {
		for i := 0; i < len(x.Children); i++ {
			if x.Children[i] != nil {
				x.Children[i].Parent = x
				SrediRoditelje(x.Children[i])
			}

		}

	}
}

func Ispis(x *Node, nivo int) {
	if x != nil {
		fmt.Print("nivo ", nivo, ":")
		for j := 0; j < BrEl(x.Key); j++ {
			fmt.Print(x.Key[j], " ")
		}
		fmt.Println()
		nivo++
		if x.Children[0] != nil { //ako ima decu
			for i := 0; i < len(x.Children) && x.Children[i] != nil; i++ {
				Ispis(x.Children[i], nivo)
			}
		}
	}
}
