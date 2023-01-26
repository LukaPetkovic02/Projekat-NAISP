package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	velicina  int
	korisceni *list.List
	hash_mapa map[string]*list.Element
}

type Element struct {
	kljuc    string
	vrednost []byte
}

func (lru *LRUCache) dodaj(kljuc string, vrednost []byte) {

	if vr, ima := lru.hash_mapa[kljuc]; ima {
		lru.korisceni.MoveToFront(vr)
		lru.hash_mapa[kljuc].Value.(*Element).vrednost = vrednost
	} else {

		kljuc_vr := &Element{kljuc: kljuc, vrednost: vrednost}
		lru.hash_mapa[kljuc] = lru.korisceni.PushFront(kljuc_vr)

		if lru.velicina < lru.korisceni.Len() {
			delete(lru.hash_mapa, lru.korisceni.Back().Value.(*Element).kljuc)
			lru.korisceni.Remove(lru.korisceni.Back())
		}
	}

}

func (lru *LRUCache) citaj(kljuc string) []byte {

	if vr, ima := lru.hash_mapa[kljuc]; ima {

		lru.korisceni.MoveToFront(vr)
		return lru.hash_mapa[kljuc].Value.(*Element).vrednost

	} else {
		return nil
	}

}

func main() {

	lru_cache := &LRUCache{velicina: 5, korisceni: list.New(), hash_mapa: make(map[string]*list.Element)}
	lru_cache.dodaj("1", []byte("vr1"))
	lru_cache.dodaj("2", []byte("vr2"))
	lru_cache.dodaj("3", []byte("vr3"))
	lru_cache.dodaj("4", []byte("vr4"))
	lru_cache.dodaj("5", []byte("vr5"))
	lru_cache.dodaj("6", []byte("vr6"))
	lru_cache.dodaj("7", []byte("vr7"))
	lru_cache.citaj("3")
	lru_cache.dodaj("8", []byte("vr8"))

	if lru_cache.citaj("3") == nil {
		fmt.Println("Ne postoji")
	} else {
		fmt.Println(string(lru_cache.citaj("3")))
	}

}
