package lru

import (
	"container/list"
	importi "projekat/utils"
)

type LRUCache struct {
	Velicina  int
	Korisceni *list.List
	Hash_mapa map[string]*list.Element
}

func (lru *LRUCache) Dodaj(podatak importi.Podatak) {

	el, ima := lru.Hash_mapa[podatak.Key]

	if ima {

		lru.Korisceni.MoveToFront(el)
		el.Value = podatak
	} else {

		lru.Hash_mapa[podatak.Key] = lru.Korisceni.PushFront(podatak)

		if lru.Velicina < lru.Korisceni.Len() {
			delete(lru.Hash_mapa, lru.Korisceni.Back().Value.(importi.Podatak).Key)
			lru.Korisceni.Remove(lru.Korisceni.Back())
		}

	}

}

func (lru *LRUCache) Citaj(kljuc string) []byte {

	el, ima := lru.Hash_mapa[kljuc]

	if ima {

		lru.Korisceni.MoveToFront(el)
		return el.Value.(importi.Podatak).Value

	} else {
		return nil
	}

}

func NoviLRU() LRUCache {

	lru_cache := &LRUCache{Velicina: 5, Korisceni: list.New(), Hash_mapa: make(map[string]*list.Element)}
	return *lru_cache
}