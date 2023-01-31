package main

import (
	"fmt"
	"projekat/bloomFilter"
	Btree "projekat/btree"
	Cms "projekat/countMinSketch"
	Hajp "projekat/hajperll"
	Lru "projekat/lru"
	Memtable "projekat/meemtable"
	SkipList "projekat/skipList"
	Util "projekat/utils"
	//"time"
)

func main() {
	//Bloom filter testovi

	fmt.Println("joe")
	bloom := bloomFilter.NewBloomFilter(5, 0.01)
	// fmt.Println(bloom.M, bloom.K)
	// fmt.Println(bloom.Fns)
	bloom.Add([]byte("wasd"))
	// fmt.Println(bloom.Podaci)
	// fmt.Println(bloom.Search([]byte("wasd")))
	// fmt.Println(bloom.Search([]byte("nesto drugo")))
	var p Util.Podatak
	p = Util.NewPodatak("1", []byte("bajtovi"), 0)
	//skip lista testovi
	var sk SkipList.SkipList
	sk.InitSP(10, 5, 5)
	sk.Put(p)
	//B stablo testovi

	var s Btree.Stablo
	s.InitSP(3, 10)

	//testiranje memtablea
	var mem Memtable.MemTable
	//mem = &s
	mem.Put(p)
	fmt.Println("memtable kao stablo") //samo gledam dal radi
	fmt.Println(mem)

	mem = &sk
	mem.Put(p)
	fmt.Println("memtable kao skip lista")
	fmt.Println(mem)
	//fmt.Println(s)
	//s.Put("a", []byte("nesto"), time.Now())
	//s.Put("b", []byte("nesto"), time.Now())
	//cvor, i := s.Search("4")
	//fmt.Println(cvor, i)
	/*s.Put("c", []byte("1estodrugo"), time.Now())
	s.Put("d", []byte("2estodrugo"), time.Now())
	s.Put("e", []byte("3estodrugo"), time.Now())
	s.Put("f", []byte("4estodrugo"), time.Now())
	s.Put("g", []byte("5estodrugo"), time.Now())
	s.Put("h", []byte("6estodrugo"), time.Now())
	s.Put("i", []byte("7estodrugo"), time.Now())
	s.Put("j", []byte("1estodrugo"), time.Now())
	s.Put("k", []byte("2estodrugo"), time.Now())
	s.Put("l", []byte("3estodrugo"), time.Now())
	s.Put("m", []byte("2estodrugo"), time.Now())
	s.Put("n", []byte("3estodrugo"), time.Now())
	s.Put("o", []byte("4estodrugo"), time.Now())
	s.Put("p", []byte("5estodrugo"), time.Now())
	s.Put("q", []byte("6estodrugo"), time.Now())
	s.Put("r", []byte("7estodrugo"), time.Now())
	s.Put("s", []byte("1estodrugo"), time.Now())
	s.Put("t", []byte("2estodrugo"), time.Now())
	s.Put("u", []byte("3estodrugo"), time.Now())
	s.Put("v", []byte("4estodrugo"), time.Now())
	s.Put("w", []byte("5estodrugo"), time.Now())
	s.Put("x", []byte("6estodrugo"), time.Now())
	s.Put("y", []byte("7estodrugo"), time.Now())
	s.Put("z", []byte("7estodrugo"), time.Now())*/

	fmt.Println("PRE")
	//s.Put("a", []byte("promenjeno"), time.Now())
	//s.Delete("a", time.Now())

	//Btree.Ispis(s.Head, 0)
	v1 := s.Get("a")
	v2 := s.Get("b")
	fmt.Println(v1)
	fmt.Println(v2)

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

	//lru test

	lru_cache := Lru.NoviLRU(5)

	lru_cache.Dodaj(Util.NewPodatak("a", []byte("1estodrugo"), 1))
	lru_cache.Dodaj(Util.NewPodatak("b", []byte("2estodrugo"), 1))
	lru_cache.Dodaj(Util.NewPodatak("c", []byte("3estodrugo"), 1))
	lru_cache.Dodaj(Util.NewPodatak("d", []byte("4estodrugo"), 1))
	lru_cache.Dodaj(Util.NewPodatak("e", []byte("4estodrugo"), 1)) //zakomentarisi par elemenata da b ostane u lru-u da bi radilo
	lru_cache.Dodaj(Util.NewPodatak("f", []byte("4estodrugo"), 1))
	lru_cache.Dodaj(Util.NewPodatak("a", []byte("7estodrugo"), 1))

	if lru_cache.Citaj("b") == nil {
		fmt.Println("Ne postoji")
	} else {
		fmt.Println(string(lru_cache.Citaj("b")))
	}
}
