package main

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/App"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/engine"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/lru"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/memtable"
	"github.com/LukaPetkovicSV16/Projekat-NAISP/skipList"
)

func main() {
	engine.CreateDataFolderStructure()
	var sl = &skipList.SkipList{}
	var LRU = lru.NewLRU(int(config.Values.Cache.Size))
	sl.InitSP(10, 10, 100)
	var memtable = memtable.Init(100, sl)
	App.TUI(memtable, LRU)

	//fmt.Println(cms.Deserialize(bajtovi).Encoder)
	//fmt.Println(cms.Deserialize(bajtovi).M)
	//fmt.Println(cms.Deserialize(bajtovi).Data)
	//fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto")))
	//fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto123")))
	//fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto55")))
}
