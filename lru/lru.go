package lru

import (
	"container/list"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

type LRUCache struct {
	LRU_Size int
	Recent   *list.List
	Hash_Map map[string]*list.Element
}

func (lru *LRUCache) Add(rec types.Record) {

	el, in := lru.Hash_Map[rec.Key]

	if in {

		lru.Recent.MoveToFront(el)
		el.Value = &rec
	} else {

		lru.Hash_Map[rec.Key] = lru.Recent.PushFront(&rec)

		if lru.LRU_Size < lru.Recent.Len() {
			delete(lru.Hash_Map, lru.Recent.Back().Value.(types.Record).Key)
			lru.Recent.Remove(lru.Recent.Back())
		}

	}

}

func (lru *LRUCache) Read(kljuc string) *types.Record {

	el, in := lru.Hash_Map[kljuc]

	if in && el.Value.(*types.Record).Tombstone == false {

		lru.Recent.MoveToFront(el)
		return el.Value.(*types.Record)

	} else {
		return nil
	}

}

func NewLRU(s int) *LRUCache {

	lru_cache := &LRUCache{LRU_Size: s, Recent: list.New(), Hash_Map: make(map[string]*list.Element)}
	return lru_cache
}
