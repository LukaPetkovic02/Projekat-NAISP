package bTree

import (
	"fmt"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
	//"time"
)

type Stablo struct {
	Max  int
	Head *Node
	//Max_capacity int
	Cur_capacity int
}

type Node struct {
	Value    []types.Record
	Children []*Node
	Parent   *Node
}

func (s *Node) InitSP(Max int, Parent *Node) Node {
	s.Value = make([]types.Record, Max+1)
	//s.Value[0] = types.Record{Key: "", Value: nil, Tombstone: 1, Timestamp: time.Now().Unix()} //inicijalizujemo prazan podatak u pocetku
	s.Value[0] = types.CreateRecord("", nil, false)
	s.Children = make([]*Node, Max+2) //jer ce kod podele cvorova u jednom momentu biti vise dece
	s.Parent = Parent
	s.Parent.Value = make([]types.Record, Max+1)
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
	//st.Max_capacity = Max_capacity
	st.Cur_capacity = 0
	return *st
}

func (s *Stablo) search(SearchKey string) (*Node, int) { //ako nema vraca nil, ako ima vraca node i njegovu poziciju u nizu podataka u tom cvoru
	x := s.Head
	t := true
	var i int
	for x != nil {
		t = true
		for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
			//byt:=[]byte(SearchKey)
			if SearchKey == x.Value[i].Key && x.Value[i].Tombstone == false {
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

func (s *Stablo) Get(SearchKey string) *types.Record { //ako nema vraca nil, ako ima vraca podatak
	x := s.Head
	t := true
	var i int
	for x != nil {
		t = true
		for i = 0; i < len(x.Value)-1 && string(x.Value[i].Key) != ""; i++ {
			if SearchKey == x.Value[i].Key {
				return &x.Value[i]
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
	return nil
}

func brEl(Value []types.Record) int {
	var i int

	for i = 0; i < len(Value) && Value[i].Key != ""; i++ {
	}
	return i
}

func podeliCvor(x *Node, Max int) (*Node, *Node, *Node) {
	var sredina int = Max / 2
	x1 := new(Node) //pre sredine
	x1.Value = make([]types.Record, Max+1)
	for i, p := range x1.Value {
		p = types.CreateRecord("", nil, false)
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
	x2.Value = make([]types.Record, Max+1)
	for i, p := range x2.Value {
		p = types.CreateRecord("", nil, false)
		x2.Value[i] = p
	}
	x2.Children = make([]*Node, Max+2)
	//x2 ni ne treba da ima decu jer on svakako ide gore
	//x2.Parent = x.Parent
	x3 := new(Node) //posle sredine
	x3.Value = make([]types.Record, Max+1)
	for i, p := range x3.Value {
		p = types.CreateRecord("", nil, false)
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
	for i := sredina + 1; i < brEl(x.Value); i++ {
		//x3.Key[i-sredina-1] = x.Key[i]
		//x3.Value[i-sredina-1] = x.Value[i]
		x3.Value[i-sredina-1] = x.Value[i]
	}
	return x1, x2, x3
}

func (s *Stablo) Delete(key string) bool {
	a, _ := s.search(key)
	if a != nil {
		x := s.Head
		t := true
		var i int
		for x != nil {
			t = true
			for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
				if key == x.Value[i].Key && x.Value[i].Tombstone == false {
					//x.Value[i].Value = podatak.Value
					//x.Value[i].Timestamp = podatak.Timestamp
					x.Value[i].Tombstone = true
					//x.Value[i].Timestamp = podatak.Timestamp
					return true
				} else if x.Value[i].Key > key {
					x = x.Children[i]
					t = false
					break
				}
			}
			if t {
				x = x.Children[i]
			}
		}
		//return true
	}
	fmt.Println("Record with key ", key, " doesn't exist or his tombstone is already true!")
	return false
}

func (s *Stablo) Add(podatak types.Record) bool {
	if s.Head.Value[0].Key == "" { //ako je prazno stablo
		//s.Head.Value[0].Key = addKey
		//s.Head.Value[0].Value = addValue
		//s.Head.Value[0].Tombstone = 0
		//s.Head.Value[0].Timestamp = vreme.Unix()
		s.Head.Value[0] = podatak
		s.Cur_capacity += 1
		return true
	}

	a, _ := s.search(podatak.Key)

	if a != nil { //ako element vec postoji izmeni ga
		x := s.Head
		t := true
		var i int
		for x != nil {
			t = true
			for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
				if podatak.Key == x.Value[i].Key && x.Value[i].Tombstone == false {
					x.Value[i].Value = podatak.Value
					x.Value[i].Timestamp = podatak.Timestamp
					x.Value[i].Tombstone = podatak.Tombstone
					x.Value[i].Timestamp = podatak.Timestamp
					x.Value[i].CRC = podatak.CRC
					x.Value[i].KeySize = podatak.KeySize
					x.Value[i].ValueSize = podatak.ValueSize
					return true
				} else if x.Value[i].Key > podatak.Key {
					x = x.Children[i]
					t = false
					break
				}
			}
			if t {
				x = x.Children[i]
			}
		}
		//fmt.Println("That record already exists!")
		return true
	}

	//if s.Cur_capacity >= s.Max_capacity {
	//	return false
	//}

	//ako ne postoji onda je dodavanje
	x := s.Head
	t := true
	var i int
	for x.Children[0] != nil { //ne moze x.Children!=nil;;ako mu je prvo dete nil znaci da nema dece
		for i = 0; i < len(x.Value)-1 && x.Value[i].Key != ""; i++ {
			if x.Value[i].Key > podatak.Key {
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
	temps := podatak.Key
	tempb := podatak.Value
	var tempt bool
	tempt = podatak.Tombstone
	tempv := podatak.Timestamp
	tempcrc := podatak.CRC
	tempks := podatak.KeySize
	tempvs := podatak.ValueSize
	// kad nadjemo
	for i = 0; i < brEl(x.Value); i++ {
		if temps < x.Value[i].Key {
			x.Value[i].Key, temps = temps, x.Value[i].Key
			x.Value[i].Value, tempb = tempb, x.Value[i].Value
			x.Value[i].Tombstone, tempt = tempt, x.Value[i].Tombstone
			x.Value[i].Timestamp, tempv = tempv, x.Value[i].Timestamp
			x.Value[i].CRC, tempcrc = tempcrc, x.Value[i].CRC
			x.Value[i].KeySize, tempks = tempks, x.Value[i].KeySize
			x.Value[i].ValueSize, tempvs = tempvs, x.Value[i].ValueSize
		}
	}
	x.Value[i].Key = temps
	x.Value[i].Value = tempb
	x.Value[i].Tombstone = tempt
	x.Value[i].Timestamp = tempv
	x.Value[i].CRC = tempcrc
	x.Value[i].KeySize = tempks
	x.Value[i].ValueSize = tempvs

	//ako je doslo do overflowa radimo dodatno
	if brEl(x.Value) == len(x.Value) {
		//if t { //ako ne postoji sibling koji nije popunjen
		//uradi podelu cvorova
		for brEl(x.Value) == s.Max+1 { //dok je trenutni nivo popunjen
			x1, x2, x3 := podeliCvor(x, s.Max)
			x.Value = make([]types.Record, 0)
			for a := 0; a < len(x1.Value) && x1.Value[a].Key != ""; a++ {
				x.Value = append(x.Value, x1.Value[a])
			}
			for a := 0; a < len(x3.Value) && x3.Value[a].Key != ""; a++ {
				x.Value = append(x.Value, x3.Value[a])
			}
			var p types.Record
			p = types.CreateRecord("", nil, false)
			x.Value = append(x.Value, p)

			tempk := x2.Value[0].Key //srednji ima samo jedan Key
			tempv := x2.Value[0].Value
			tempt := x2.Value[0].Tombstone
			tempvr := x2.Value[0].Timestamp
			tempcrc := x2.Value[0].CRC
			tempks := x2.Value[0].KeySize
			tempvs := x2.Value[0].ValueSize
			j := 0
			for j = 0; j < brEl(x.Parent.Value); j++ { //srednji kljuc dajemo roditelju na odgovarajuce mesto
				if tempk < x.Parent.Value[j].Key {
					x.Parent.Value[j].Key, tempk = tempk, x.Parent.Value[j].Key
					x.Parent.Value[j].Value, tempv = tempv, x.Parent.Value[j].Value
					x.Parent.Value[j].Tombstone, tempt = tempt, x.Parent.Value[j].Tombstone
					x.Parent.Value[j].Timestamp, tempvr = tempvr, x.Parent.Value[j].Timestamp
					x.Parent.Value[j].CRC, tempcrc = tempcrc, x.Parent.Value[j].CRC
					x.Parent.Value[j].KeySize, tempks = tempks, x.Parent.Value[j].KeySize
					x.Parent.Value[j].ValueSize, tempvs = tempvs, x.Parent.Value[j].ValueSize
				}

			}
			x.Parent.Value[j].Key = tempk
			x.Parent.Value[j].Value = tempv
			x.Parent.Value[j].Tombstone = tempt
			x.Parent.Value[j].Timestamp = tempvr
			x.Parent.Value[j].CRC = tempcrc
			x.Parent.Value[j].KeySize = tempks
			x.Parent.Value[j].ValueSize = tempvs
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
		if brEl(x.Value) != brEl(s.Head.Parent.Value) {
			isti = false
		} else {
			for o := 0; o < brEl(x.Value); o++ {
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
			var p types.Record
			p = types.CreateRecord("", nil, false)
			s.Head.Parent.Value = make([]types.Record, s.Max+1)
			s.Head.Parent.Value[0] = p
			//s.Head.Parent.Value = append(s.Head.Parent.Value, )
			//s.Head.Parent.Key = make([]string, s.Max+1)
			//s.Head.Parent.Value = make([][]byte, s.Max+1)
			s.Head.Parent.Children = make([]*Node, s.Max+2)
			s.Head.Parent.Children[0] = x
		}

	}
	srediRoditelje(s.Head)
	//proveri dal je popunjen kapacitet
	s.Cur_capacity += 1
	return true
	/*if s.Cur_capacity == s.Max_capacity {
		A1 := s.GetSortedRecordsList()
		s.InitSP(s.Max, s.Max_capacity)
		return A1
	} else {
		return nil
	}*/
}

func srediRoditelje(x *Node) {
	if x != nil {
		for i := 0; i < len(x.Children); i++ {
			if x.Children[i] != nil {
				x.Children[i].Parent = x
				srediRoditelje(x.Children[i])
			}

		}

	}
}

/*func Ispis(x *Node, nivo int) {
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
}*/

func (s *Stablo) GetSortedRecordsList() []types.Record { //vraca listu sortiranih podataka
	Value := make([]types.Record, 0)
	s.GetSortedRecordsListFunc(s.Head, &Value)
	s.InitSP(s.Max) //isprazni stablo
	return Value
}

func (s *Stablo) GetSortedRecordsListFunc(n *Node, Value *[]types.Record) {
	if n == nil {
		return
	}
	for i := 0; i < brEl(n.Value); i++ {
		if n.Children[0] != nil {
			s.GetSortedRecordsListFunc(n.Children[i], Value)
		}
		*Value = append(*Value, n.Value[i])
	}
	if n.Children[0] != nil {
		s.GetSortedRecordsListFunc(n.Children[brEl(n.Value)], Value)
	}
}

func (s *Stablo) Clear() {
	s.InitSP(s.Max)
}
func (s *Stablo) GetSize() int {
	return s.Cur_capacity
}
