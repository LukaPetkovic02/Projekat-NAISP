package main

import (
	"fmt"
	"projekat/bloomFilter"
	Btree "projekat/btree"
	Cms "projekat/countMinSketch"
	Hajp "projekat/hajperll"
	Lru "projekat/lru"
	Util "projekat/utils"
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

	//B stablo testovi

	var s Btree.Stablo
	s.InitSP(3, 20)
	fmt.Println(s)
	s.Put(Util.NewPodatak("a", []byte("nesto"), 1))
	s.Put(Util.NewPodatak("b", []byte("nesto2"), 1))

	sortirani := s.GetAllData()
	for i := 0; i < len(sortirani); i++ {
		sortirani[i].PrintData()
	}

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
