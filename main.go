package main

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/bTree"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/tokenBucket"
)

func main() {
	engine.CreateDataFolderStructure()
	var sl = &skipList.SkipList{}
	var LRU = lru.NewLRU(int(config.Values.Cache.Size))
	if config.Values.Memtable.Use == "skip-list" {
		var sl = &skipList.SkipList{}
		sl.InitSP(int(config.Values.SkipList.MaxLevel), int(config.Values.SkipList.MaxLevel/2))
	} else {
		var sl = &bTree.Stablo{}
		sl.InitSP(config.Values.Btree.MaxNode)
	}

	var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
	var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
	App.TUI(memtable, LRU, token)

	//fmt.Println(cms.Deserialize(bajtovi).Encoder)
	//fmt.Println(cms.Deserialize(bajtovi).M)
	//fmt.Println(cms.Deserialize(bajtovi).Data)
	//fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto")))
	//fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto123")))
	//fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto55")))
}
