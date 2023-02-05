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

// func da(memtable *memtable.Memtable, LRU *lru.LRUCache, token *tokenBucket.TokenBucket) {
// 	App.HandleAdd("a", []byte("fff"), memtable, LRU)
// 	App.HandleAdd("b", []byte("fff"), memtable, LRU)
// 	App.HandleAdd("c", []byte("fff"), memtable, LRU)
// 	App.HandleAdd("a", []byte("gggg"), memtable, LRU)
// 	App.HandleDelete("b", memtable, LRU)
// 	App.HandleAdd("j", []byte("fff"), memtable, LRU)
// }

func main() {
	engine.CreateDataFolderStructure()
	var LRU = lru.NewLRU(int(config.Values.Cache.Size))

	if config.Values.Memtable.Use == "skip-list" {
		var sl = &skipList.SkipList{}
		sl.InitSP(int(config.Values.SkipList.MaxLevel), int(config.Values.SkipList.Height))
		var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
		var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
		App.TUI(memtable, LRU, token)
	} else {
		var sl = &bTree.Stablo{}
		sl.InitSP(config.Values.Btree.MaxNode)
		var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
		var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
		App.TUI(memtable, LRU, token)
	}

	// x := sstable.ReadAllRecordsFromTable("2_1675616501630928100.bin")
	// for _, v := range x {
	// 	fmt.Println(v)
	// }
	// var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
	// var token = tokenBucket.Init(uint64(config.Values.TokenBucket.Size), config.Values.TokenBucket.Rate)
	// App.TUI(memtable, LRU, token)

	// fmt.Println(cms.Deserialize(bajtovi).Encoder)
	// fmt.Println(cms.Deserialize(bajtovi).M)
	// fmt.Println(cms.Deserialize(bajtovi).Data)
	// fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto")))
	// fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto123")))
	// fmt.Println(cms.Deserialize(bajtovi).Frequency([]byte("nesto55")))
}
