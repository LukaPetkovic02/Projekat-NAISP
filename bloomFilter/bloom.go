package bloomFilter

import (
	"bytes"
	"encoding/gob"
	importi "projekat/utils"
)

type BloomFilter struct {
	Fns     []importi.HashWithSeed //sve hash funkcije kojih ima k u bloom filteru
	Podaci  []byte                 //niz podataka koji nam govori da li element postoji ili ne
	M       uint                   //ocekivani broj elemenata
	K       uint                   //zeljeni broj hash funkcija
	Encoder *gob.Encoder           //enkoder za hash funkcije
	Decoder *gob.Decoder           //dekoder za hash funkcije
}

// kreira novi bloom filter za ocekivani br elemenata i false positive rate
func NewBloomFilter(excepted_elements int, false_positive_rate float64) *BloomFilter {

	bloom := new(BloomFilter)
	bloom.M = CalculateM(excepted_elements, false_positive_rate)
	bloom.K = CalculateK(excepted_elements, bloom.M)
	bloom.Podaci = make([]byte, bloom.M)
	bloom.Fns = importi.CreateHashFunctions(bloom.K)
	var buf = &bytes.Buffer{}
	bloom.Encoder = gob.NewEncoder(buf)
	bloom.Decoder = gob.NewDecoder(buf)
	return bloom
}

// kreira bloom filter od vec zadatih podataka koji ce se ucitavati iz fajla za vec unapred napravljene bloom filtere
func RecreateBloomFilter(m uint, k uint, fns []importi.HashWithSeed, podaci []byte) *BloomFilter {

	bloom := new(BloomFilter)
	bloom.M = m
	bloom.K = k
	bloom.Podaci = podaci
	bloom.Fns = fns
	var buf = &bytes.Buffer{}
	bloom.Encoder = gob.NewEncoder(buf)
	bloom.Decoder = gob.NewDecoder(buf)
	return bloom
}

// dodaje element u bloom filter
func (bloom *BloomFilter) Add(data []byte) {
	var i uint64
	for _, fn := range bloom.Fns {
		err := bloom.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &importi.HashWithSeed{}
		err = bloom.Decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		i = dfn.Hash(data)
		i = i % uint64(bloom.M)
		bloom.Podaci[i] = 1
	}
}

// pretrazuje bloom filter i govori da li element postoji ili ne, moze reci da postoji element koji ne postoji
func (bloom *BloomFilter) Search(data []byte) bool {
	var i uint64
	for _, fn := range bloom.Fns {
		err := bloom.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &importi.HashWithSeed{}
		err = bloom.Decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		i = dfn.Hash(data)
		i = i % uint64(bloom.M)
		if bloom.Podaci[i] == 0 {
			return false
		}
	}
	return true
}
