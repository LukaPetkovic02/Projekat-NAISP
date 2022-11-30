package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var fns = CreateHashFunctions(5)
var buf = &bytes.Buffer{}
var encoder = gob.NewEncoder(buf)
var decoder = gob.NewDecoder(buf)
var niz []byte

func dodaj(data []byte) {
	var i uint64
	for _, fn := range fns {
		//data := []byte("hello")
		//fmt.Println(fn.Hash(data))
		err := encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		i = dfn.Hash(data)
		i = i % uint64(len(niz))
		niz[i] = 1
		//fmt.Println(dfn.Hash(data))
	}
}
func search(data []byte) bool {
	var i uint64
	for _, fn := range fns {
		err := encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		i = dfn.Hash(data)
		i = i % uint64(len(niz))
		if niz[i] == 0 {
			return false
		}
	}
	return true
}
func main() {
	m := CalculateM(5, 0.01)
	k := CalculateK(5, m)
	fmt.Println(m, k)
	niz = make([]byte, m)
	dodaj([]byte("wasd"))
	fmt.Println(niz)
	fmt.Println(search([]byte("wasd")))
	fmt.Println(search([]byte("nesto drugo")))
}
