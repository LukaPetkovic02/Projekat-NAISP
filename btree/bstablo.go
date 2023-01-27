package main

import "fmt"

type Stablo struct {
	max  int
	head *Node
}

type Node struct {
	key      []string
	value    [][]byte
	children []*Node
	parent   *Node
}

func (s *Node) InitSP(key string, value []byte, max int, parent *Node) Node {
	s.key = make([]string, max+1)
	s.value = make([][]byte, max+1)
	s.children = make([]*Node, max+2) //jer ce kod podele cvorova u jednom momentu biti vise dece
	s.parent = parent
	s.key[0] = key
	s.value[0] = value

	s.parent.key = make([]string, max+1)
	s.parent.value = make([][]byte, max+1)
	s.parent.children = make([]*Node, max+2)
	s.parent.children[0] = s
	return *s
}

func (st *Stablo) InitSP(key string, value []byte, max int) Stablo {
	nilRoditelj := new(Node)
	st.head = new(Node)
	st.head.InitSP(key, value, max, nilRoditelj)
	st.max = max
	return *st
}

func (s *Stablo) search(searchKey string) (*Node, int) {
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

func brEl(key []string) int {
	var i int

	for i = 0; i < len(key) && key[i] != ""; i++ {
	}
	return i
}
func podeliCvor(x *Node, max int) (*Node, *Node, *Node) {
	var sredina int = max / 2
	x1 := new(Node) //pre sredine
	x1.key = make([]string, max+1)
	x1.value = make([][]byte, max+1)
	x1.children = make([]*Node, max+2)
	//deca od x1 su prvih len/2 dece x
	for f := 0; f < len(x.children)/2; f++ {
		x1.children[f] = x.children[f]
	}
	//x1.parent = x.parent
	x2 := new(Node) //sredina
	x2.key = make([]string, max+1)
	x2.value = make([][]byte, max+1)
	x2.children = make([]*Node, max+2)
	//x2 ni ne treba da ima decu jer on svakako ide gore
	//x2.parent = x.parent
	x3 := new(Node) //posle sredine
	x3.key = make([]string, max+1)
	x3.value = make([][]byte, max+1)
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
	for i := sredina + 1; i < brEl(x.key); i++ {
		x3.key[i-sredina-1] = x.key[i]
		x3.value[i-sredina-1] = x.value[i]
	}
	return x1, x2, x3
}
func (s *Stablo) add(addKey string, addValue []byte) {
	a, _ := s.search(addKey)
	if a != nil {
		fmt.Println("Element vec postoji sa tim kljucem.")
		return
	}

	x := s.head
	t := true
	var i int
	for x.children[0] != nil { //ne moze x.children!=nil;;ako mu je prvo dete nil znaci da nema dece
		for i = 0; i < len(x.key)-1 && x.key[i] != ""; i++ {
			if x.key[i] > addKey {
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
	// kad nadjemo
	for i = 0; i < brEl(x.key); i++ {
		if temps < x.key[i] {
			x.key[i], temps = temps, x.key[i]
			x.value[i], tempb = tempb, x.value[i]
		}
	}
	x.key[i] = temps
	x.value[i] = tempb

	//ako je doslo do overflowa radimo dodatno
	if brEl(x.key) == len(x.key) {
		//if t { //ako ne postoji sibling koji nije popunjen
		//uradi podelu cvorova
		for brEl(x.key) == s.max+1 { //dok je trenutni nivo popunjen
			x1, x2, x3 := podeliCvor(x, s.max)
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
			for j = 0; j < brEl(x.parent.key); j++ { //srednji kljuc dajemo roditelju na odgovarajuce mesto
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
		if brEl(x.key) != brEl(s.head.parent.key) {
			isti = false
		} else {
			for o := 0; o < brEl(x.key); o++ {
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
			s.head.parent.value = make([][]byte, s.max+1)
			s.head.parent.children = make([]*Node, s.max+2)
			s.head.parent.children[0] = x
		}

	}
	srediRoditelje(s.head)
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
		for j := 0; j < brEl(x.key); j++ {
			fmt.Print(x.key[j], " ")
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

func main() {
	var s Stablo
	s.InitSP("a", []byte("koren"), 3)
	s.add("b", []byte("nesto"))
	//cvor, i := s.search("4")
	//fmt.Println(cvor, i)
	s.add("c", []byte("1estodrugo"))
	s.add("d", []byte("2estodrugo"))
	s.add("e", []byte("3estodrugo"))
	s.add("f", []byte("4estodrugo"))
	s.add("g", []byte("5estodrugo"))
	s.add("h", []byte("6estodrugo"))
	s.add("i", []byte("7estodrugo"))
	s.add("j", []byte("1estodrugo"))
	s.add("k", []byte("2estodrugo"))
	s.add("l", []byte("3estodrugo"))
	s.add("m", []byte("2estodrugo"))
	s.add("n", []byte("3estodrugo"))
	s.add("o", []byte("4estodrugo"))
	s.add("p", []byte("5estodrugo"))
	s.add("q", []byte("6estodrugo"))
	s.add("r", []byte("7estodrugo"))
	s.add("s", []byte("1estodrugo"))
	s.add("t", []byte("2estodrugo"))
	s.add("u", []byte("3estodrugo"))
	s.add("v", []byte("4estodrugo"))
	s.add("w", []byte("5estodrugo"))
	s.add("x", []byte("6estodrugo"))
	s.add("y", []byte("7estodrugo"))
	s.add("z", []byte("7estodrugo"))
	fmt.Println("PRE")
	ispis(s.head, 0)
}
