package btree

import (
	"fmt"
	importi "projekat/utils"
	"time"
)

//kasnije treba import za podatak

type Stablo struct {
	Max  int
	Head *Node
}

type Node struct {
	Value    []importi.Podatak
	Children []*Node
	Parent   *Node
}

func (s *Node) InitSP(Max int, Parent *Node) Node {
	s.Value = make([]importi.Podatak, Max+1)
	s.Value[0] = importi.Podatak{Key: "", Value: nil, Tombstone: 1, Timestamp: time.Now().Unix()} //inicijalizujemo prazan podatak u pocetku
	s.Children = make([]*Node, Max+2)                                                             //jer ce kod podele cvorova u jednom momentu biti vise dece
	s.Parent = Parent
	s.Parent.Value = make([]importi.Podatak, Max+1)
	s.Parent.Children = make([]*Node, Max+2)
	s.Parent.Children[0] = s
	return *s
}

func (st *Stablo) InitSP(Max int) Stablo {
	nilRoditelj := new(Node)
	st.Head = new(Node)
	st.Head.InitSP(Max, nilRoditelj)
	//st.Head.Parent = nilRoditelj
	//st.Head.Children = make([]*Node, Max+2)
	//st.Head.InitSP(Key, Value, Max, nilRoditelj)
	st.Max = Max
	return *st
}

func (s *Stablo) Search(SearchKey string) (*Node, int) {
	x := s.Head
	t := true
	var i int
	for x != nil {
		t = true
		for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
			if SearchKey == x.Value[i].Key && x.Value[i].Tombstone == 0 {
				return x, i
			} else if x.Value[i].Key > SearchKey {
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

func BrEl(Value []importi.Podatak) int {
	var i int

	for i = 0; i < len(Value) && Value[i].Key != ""; i++ {
	}
	return i
}

func PodeliCvor(x *Node, Max int) (*Node, *Node, *Node) {
	var sredina int = Max / 2
	x1 := new(Node) //pre sredine
	x1.Value = make([]importi.Podatak, Max+1)
	for i, p := range x1.Value {
		p.InitPrazanPodatak()
		x1.Value[i] = p
	}
	//x1.Value.Key = make([]string, Max+1)
	//x1.Value = make([][]byte, Max+1)
	x1.Children = make([]*Node, Max+2)
	//deca od x1 su prvih len/2 dece x
	for f := 0; f < len(x.Children)/2; f++ {
		x1.Children[f] = x.Children[f]
	}
	//x1.Parent = x.Parent
	x2 := new(Node) //sredina
	x2.Value = make([]importi.Podatak, Max+1)
	for i, p := range x2.Value {
		p.InitPrazanPodatak()
		x2.Value[i] = p
	}
	x2.Children = make([]*Node, Max+2)
	//x2 ni ne treba da ima decu jer on svakako ide gore
	//x2.Parent = x.Parent
	x3 := new(Node) //posle sredine
	x3.Value = make([]importi.Podatak, Max+1)
	for i, p := range x3.Value {
		p.InitPrazanPodatak()
		x3.Value[i] = p
	}
	x3.Children = make([]*Node, Max+2)
	//zadnjih pola dece xa
	for f := len(x.Children) / 2; f < len(x.Children); f++ {
		x3.Children[f-len(x.Children)/2] = x.Children[f]
	}
	//x3.Parent = x.Parent
	for i := 0; i < sredina; i++ {
		x1.Value[i] = x.Value[i]
		//x1.Key[i] = x.Key[i]
		//x1.Value[i] = x.Value[i]
	}
	//x2.Key[0] = x.Key[sredina]
	//x2.Value[0] = x.Value[sredina]
	x2.Value[0] = x.Value[sredina]
	for i := sredina + 1; i < BrEl(x.Value); i++ {
		//x3.Key[i-sredina-1] = x.Key[i]
		//x3.Value[i-sredina-1] = x.Value[i]
		x3.Value[i-sredina-1] = x.Value[i]
	}
	return x1, x2, x3
}

func (s *Stablo) Put(addKey string, addValue []byte, vreme time.Time) {
	if s.Head.Value[0].Key == "" { //ako je prazno stablo
		s.Head.Value[0].Key = addKey
		s.Head.Value[0].Value = addValue
		s.Head.Value[0].Tombstone = 0
		s.Head.Value[0].Timestamp = vreme.Unix()
		return
	}

	a, _ := s.Search(addKey)

	if a != nil { //ako element vec postoji, samo cemo mu promeniti Value i timestamp!!!!
		x := s.Head
		t := true
		var i int
		for x != nil {
			t = true
			for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
				if addKey == x.Value[i].Key && x.Value[i].Tombstone == 0 {
					x.Value[i].Value = addValue
					x.Value[i].Timestamp = vreme.Unix()
					return
				} else if x.Value[i].Key > addKey {
					x = x.Children[i]
					t = false
					break
				}
			}
			if t {
				x = x.Children[i]
			}
		}
	}

	x := s.Head
	t := true
	var i int
	for x.Children[0] != nil { //ne moze x.Children!=nil;;ako mu je prvo dete nil znaci da nema dece
		for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
			if x.Value[i].Key > addKey {
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
	temps := addKey
	tempb := addValue
	var tempt byte
	tempt = 0
	tempv := vreme.Unix()
	// kad nadjemo
	for i = 0; i < BrEl(x.Value); i++ {
		if temps < x.Value[i].Key {
			x.Value[i].Key, temps = temps, x.Value[i].Key
			x.Value[i].Value, tempb = tempb, x.Value[i].Value
			x.Value[i].Tombstone, tempt = tempt, x.Value[i].Tombstone
			x.Value[i].Timestamp, tempv = tempv, x.Value[i].Timestamp
		}
	}
	x.Value[i].Key = temps
	x.Value[i].Value = tempb
	x.Value[i].Tombstone = tempt
	x.Value[i].Timestamp = tempv

	//ako je doslo do overflowa radimo dodatno
	if BrEl(x.Value) == len(x.Value) {
		//if t { //ako ne postoji sibling koji nije popunjen
		//uradi podelu cvorova
		for BrEl(x.Value) == s.Max+1 { //dok je trenutni nivo popunjen
			x1, x2, x3 := PodeliCvor(x, s.Max)
			x.Value = make([]importi.Podatak, 0)
			for a := 0; a < len(x1.Value) && x1.Value[a].Key != ""; a++ {
				x.Value = append(x.Value, x1.Value[a])
			}
			for a := 0; a < len(x3.Value) && x3.Value[a].Key != ""; a++ {
				x.Value = append(x.Value, x3.Value[a])
			}
			var p importi.Podatak
			p.InitPrazanPodatak()
			x.Value = append(x.Value, p)

			tempk := x2.Value[0].Key //srednji ima samo jedan Key
			tempv := x2.Value[0].Value
			tempt := x2.Value[0].Tombstone
			tempvr := x2.Value[0].Timestamp
			j := 0
			for j = 0; j < BrEl(x.Parent.Value); j++ { //srednji kljuc dajemo roditelju na odgovarajuce mesto
				if tempk < x.Parent.Value[j].Key {
					x.Parent.Value[j].Key, tempk = tempk, x.Parent.Value[j].Key
					x.Parent.Value[j].Value, tempv = tempv, x.Parent.Value[j].Value
					x.Parent.Value[j].Tombstone, tempt = tempt, x.Parent.Value[j].Tombstone
					x.Parent.Value[j].Timestamp, tempvr = tempvr, x.Parent.Value[j].Timestamp
				}

			}
			x.Parent.Value[j].Key = tempk
			x.Parent.Value[j].Value = tempv
			x.Parent.Value[j].Tombstone = tempt
			x.Parent.Value[j].Timestamp = tempvr
			//fmt.Println(x1.Parent)
			//fmt.Println("roditelj nakon sredjivanja:", x.Parent)
			x1.Parent = x.Parent
			//x2.Parent = x.Parent
			x3.Parent = x.Parent //tek nakon sto sredimo Parenta
			//fmt.Println(x.Parent)
			//treba pomeriti decu od kraja do j za jedno udesno
			k := 0
			for k = len(x.Parent.Children) - 1; k > j; k-- {
				x.Parent.Children[k] = x.Parent.Children[k-1]
			}
			x.Parent.Children[k+1] = x3
			x.Parent.Children[k] = x1

			x = x.Parent
		}
		//na kraju postavljamo s.Head=x ako se koren rasformirao
		isti := true
		if BrEl(x.Value) != BrEl(s.Head.Parent.Value) {
			isti = false
		} else {
			for o := 0; o < BrEl(x.Value); o++ {
				if x.Value[o].Key != s.Head.Parent.Value[o].Key {
					isti = false
				}
			}
		}
		if isti {
			//fmt.Println("rasformirao se koren kod ubacivanja", addKey)
			s.Head = x //samo ako je x jednak roditelju Heada
			//ovde treba inicijalizovati prazan Parent
			s.Head.Parent = new(Node)
			var p importi.Podatak
			p.InitPrazanPodatak()
			s.Head.Parent.Value = make([]importi.Podatak, s.Max+1)
			s.Head.Parent.Value[0] = p
			//s.Head.Parent.Value = append(s.Head.Parent.Value, )
			//s.Head.Parent.Key = make([]string, s.Max+1)
			//s.Head.Parent.Value = make([][]byte, s.Max+1)
			s.Head.Parent.Children = make([]*Node, s.Max+2)
			s.Head.Parent.Children[0] = x
		}

	}
	SrediRoditelje(s.Head)
}

func (s *Stablo) Delete(DeleteKey string, vreme time.Time) { //logicko brisanje
	a, _ := s.Search(DeleteKey)

	if a != nil { //ako element postoji, samo cemo mu promeniti timestamp!
		x := s.Head
		t := true
		var i int
		for x != nil {
			t = true
			for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
				if DeleteKey == x.Value[i].Key && x.Value[i].Tombstone == 0 {
					x.Value[i].Tombstone = 1
					x.Value[i].Timestamp = vreme.Unix()
					return
				} else if x.Value[i].Key > DeleteKey {
					x = x.Children[i]
					t = false
					break
				}
			}
			if t {
				x = x.Children[i]
			}
		}
	} else {
		fmt.Println("Element sa tim kljucem ne postoji!")
		return
	}
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
		for j := 0; j < BrEl(x.Value); j++ {
			x.Value[j].PrintData()
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

func (s *Stablo) AllDataSortedBegin() []importi.Podatak { //vraca listu sortiranih podataka
	Value := make([]importi.Podatak, 0)
	s.AllDataSorted(s.Head, &Value)
	return Value
}

func (s *Stablo) AllDataSorted(n *Node, Value *[]importi.Podatak) {
	if n == nil {
		return
	}
	for i := 0; i < BrEl(n.Value); i++ {
		if n.Children[0] != nil {
			s.AllDataSorted(n.Children[i], Value)
		}
		*Value = append(*Value, n.Value[i])
	}
	if n.Children[0] != nil {
		s.AllDataSorted(n.Children[BrEl(n.Value)], Value)
	}
}

// func main() {
// 	var s Stablo
// 	s.InitSP(3)
// 	fmt.Println(s)
// 	s.Put("a", []byte("nesto"), time.Now())
// 	s.Put("b", []byte("nesto"), time.Now())
// 	//cvor, i := s.Search("4")
// 	//fmt.Println(cvor, i)
// 	s.Put("c", []byte("1estodrugo"), time.Now())
// 	s.Put("d", []byte("2estodrugo"), time.Now())
// 	s.Put("e", []byte("3estodrugo"), time.Now())
// 	s.Put("f", []byte("4estodrugo"), time.Now())
// 	s.Put("g", []byte("5estodrugo"), time.Now())
// 	s.Put("h", []byte("6estodrugo"), time.Now())
// 	s.Put("i", []byte("7estodrugo"), time.Now())
// 	s.Put("j", []byte("1estodrugo"), time.Now())
// 	s.Put("k", []byte("2estodrugo"), time.Now())
// 	s.Put("l", []byte("3estodrugo"), time.Now())
// 	s.Put("m", []byte("2estodrugo"), time.Now())
// 	s.Put("n", []byte("3estodrugo"), time.Now())
// 	s.Put("o", []byte("4estodrugo"), time.Now())
// 	s.Put("p", []byte("5estodrugo"), time.Now())
// 	s.Put("q", []byte("6estodrugo"), time.Now())
// 	s.Put("r", []byte("7estodrugo"), time.Now())
// 	s.Put("s", []byte("1estodrugo"), time.Now())
// 	s.Put("t", []byte("2estodrugo"), time.Now())
// 	s.Put("u", []byte("3estodrugo"), time.Now())
// 	s.Put("v", []byte("4estodrugo"), time.Now())
// 	s.Put("w", []byte("5estodrugo"), time.Now())
// 	s.Put("x", []byte("6estodrugo"), time.Now())
// 	s.Put("y", []byte("7estodrugo"), time.Now())
// 	s.Put("z", []byte("7estodrugo"), time.Now())

// 	fmt.Println("PRE")
// 	s.Put("a", []byte("promenjeno"), time.Now())
// 	s.Delete("a", time.Now())
// 	sortirani := s.AllDataSorted()
// 	for i := 0; i < len(sortirani); i++ {
// 		sortirani[i].PrintData()
// 	}
// 	//Ispis(s.Head, 0)
// 	//v1, _ := s.Search("a")
// 	//v2, _ := s.Search("b")
// 	//fmt.Println(v1)
// 	//fmt.Println(v2)
// }
