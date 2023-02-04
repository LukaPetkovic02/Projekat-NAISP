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
var m = CalculateM(0.01)
var k = CalculateK(0.01)

func dodaj(data []byte, matrica [][]byte) {
	var j uint64
	for i, fn := range fns {
		err := encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(m)
		matrica[i][j] += 1
	}
}
func ucestalost(data []byte, matrica [][]byte) byte {
	var l byte = 0
	var j uint64
	niz := make([]byte, 0)
	for i, fn := range fns {
		err := encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(m)
		niz = append(niz, matrica[i][j])
	}

	l = niz[0]
	for i, val := range niz {
		if i != 0 && val < l {
			l = val
		}
	}
	return l
}
func main() {
	fmt.Println(m, k)
	matrica := make([][]byte, k)
	for i := range matrica {
		matrica[i] = make([]byte, m)
	}
	dodaj([]byte("asdasd"), matrica)
	dodaj([]byte("asdasd"), matrica)
	dodaj([]byte("assd"), matrica)
	//fmt.Println(matrica)
	fmt.Println(ucestalost([]byte("asdasd"), matrica))
	fmt.Println(ucestalost([]byte("assd"), matrica))
	fmt.Println(ucestalost([]byte("nesto drugo"), matrica))
}
