package main

import (
	"fmt"
	Bloom "projekat/bloomFilter"
	Btree "projekat/btree"
	Util "projekat/utils"
)

func main() {
	fmt.Println("joe")
	bloom := Bloom.NewBloomFilter(5, 0.01)
	fmt.Println(bloom.M, bloom.K)
	fmt.Println(bloom.Fns)
	bloom.Add([]byte("wasd"))
	fmt.Println(bloom.Podaci)
	fmt.Println(bloom.Search([]byte("wasd")))
	fmt.Println(bloom.Search([]byte("nesto drugo")))

	var s Btree.Stablo
	x := Util.NewPodatak("c", []byte("1estodrugo"), 1)
	s.InitSP(x, 3)
	//cvor, i := s.search("4")
	//fmt.Println(cvor, i)
	x = Util.NewPodatak("d", []byte("2estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("e", []byte("3estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("f", []byte("4estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("g", []byte("5estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("h", []byte("6estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("i", []byte("7estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("j", []byte("1estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("k", []byte("2estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("l", []byte("3estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("m", []byte("2estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("n", []byte("3estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("o", []byte("4estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("p", []byte("5estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("q", []byte("6estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("r", []byte("7estodrugo"), 1)
	s.Add(x)
	x = Util.NewPodatak("s", []byte("8estodrugo"), 1)
	s.Add(x)

	fmt.Println("PRE***********************************************")
	Btree.Ispis(s.Head, 0)
}
