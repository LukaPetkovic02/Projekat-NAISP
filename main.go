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

func da(memtable *memtable.Memtable, LRU *lru.LRUCache, token *tokenBucket.TokenBucket) {
	App.HandleAdd("ada", []byte("fff"), memtable, LRU)
	App.HandleAdd("abd", []byte("fff"), memtable, LRU)
	App.HandleAdd("adb", []byte("fff"), memtable, LRU)
	App.HandleAdd("adc", []byte("gggg"), memtable, LRU)
	App.HandleDelete("ada", memtable, LRU)
	App.HandleAdd("adj", []byte("gggg"), memtable, LRU)
	App.HandleAdd("ja", []byte("fff"), memtable, LRU)
	App.HandleAdd("adh", []byte("gggg"), memtable, LRU)
}

func main() {
	engine.CreateDataFolderStructure()
	var LRU = lru.NewLRU(int(config.Values.Cache.Size))

	if config.Values.Memtable.Use == "skip-list" {
		var sl = &skipList.SkipList{}
		sl.InitSP(int(config.Values.SkipList.MaxLevel), int(config.Values.SkipList.Height))
		var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
		var token = tokenBucket.Init()
		da(memtable, LRU, token)
		App.TUI(memtable, LRU, token)
	} else {
		var sl = &bTree.Stablo{}
		sl.InitSP(config.Values.Btree.MaxNode)
		var memtable = memtable.Init(int(config.Values.Memtable.Size), sl)
		var token = tokenBucket.Init()
		da(memtable, LRU, token)
		App.TUI(memtable, LRU, token)
	}
}
