package main

import (
	"fmt"
	Bloom "projekat/bloomFilter"
	Btree "projekat/btree"
	Cms "projekat/countMinSketch"
	Hajp "projekat/hajperll"
	Util "projekat/utils"
)

func main() {
	//Bloom filter testovi

	fmt.Println("joe")
	bloom := Bloom.NewBloomFilter(5, 0.01)
	// fmt.Println(bloom.M, bloom.K)
	// fmt.Println(bloom.Fns)
	bloom.Add([]byte("wasd"))
	// fmt.Println(bloom.Podaci)
	// fmt.Println(bloom.Search([]byte("wasd")))
	// fmt.Println(bloom.Search([]byte("nesto drugo")))

	//B stablo testovi

	var s Btree.Stablo
	x := Util.NewPodatak("c", []byte("1estodrugo"), 1)
	s.InitSP(x, 3)
	// x = Util.NewPodatak("d", []byte("2estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("e", []byte("3estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("f", []byte("4estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("g", []byte("5estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("h", []byte("6estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("i", []byte("7estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("j", []byte("1estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("k", []byte("2estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("l", []byte("3estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("m", []byte("2estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("n", []byte("3estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("o", []byte("4estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("p", []byte("5estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("q", []byte("6estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("r", []byte("7estodrugo"), 1)
	// s.Add(x)
	// x = Util.NewPodatak("s", []byte("8estodrugo"), 1)
	// s.Add(x)

	// fmt.Println("***********************************************PRE***********************************************")
	// Btree.Ispis(s.Head, 0)
	// fmt.Println("********************************************PRETRAGE********************************************")
	// cvor, i := s.Search("4")
	// fmt.Print("Neuspesno: ")
	// fmt.Println(cvor, i)

	// cvor, i = s.Search("c")
	// fmt.Print("Uspesno: ")
	// fmt.Println(cvor, i)

	// cvor, i = s.Search("p")
	// fmt.Print("Uspesno: ")
	// fmt.Println(cvor, i)

	//CMS testovi

	y := Cms.NewCMS(0.01)
	y.Add([]byte("asdasd"))

	// fmt.Println(y.M, y.K)
	// y.Add([]byte("asdasd"))
	// y.Add([]byte("assd"))
	// //fmt.Println(matrica)
	// fmt.Println(y.Ucestalost([]byte("asdasd")))
	// fmt.Println(y.Ucestalost([]byte("assd")))
	// fmt.Println(y.Ucestalost([]byte("nesto drugo")))

	//HLL testovi
	hajper := Hajp.MakeHLL(4)
	hajper.Add([]byte("dasdas"))
	// fmt.Println(hajper.Estimate())
	// hajper.Add([]byte("ddas"))
	// hajper.Add([]byte("dasd"))
	// hajper.Add([]byte("dasd3131as"))
	// hajper.Add([]byte("dasdas"))
	// hajper.Add([]byte("dasdas"))
	// hajper.Add([]byte("dasdas"))
	// hajper.Add([]byte("dasdas"))
	// hajper.Add([]byte("dasdas"))
	// fmt.Println(hajper.Estimate())
	// fmt.Println(hajper.Reg)
}
