package bloomFilter

import (
	"bytes"
	"encoding/gob"
)

type BloomFilter struct {
	fns     []HashWithSeed //sve hash funkcije kojih ima k u bloom filteru
	podaci  []byte         //niz podataka koji nam govori da li element postoji ili ne
	m       uint           //ocekivani broj elemenata
	k       uint           //zeljeni broj hash funkcija
	encoder *gob.Encoder   //enkoder za hash funkcije
	decoder *gob.Decoder   //dekoder za hash funkcije
}

// kreira novi bloom filter za ocekivani br elemenata i false positive rate
func NewBloomFilter(excepted_elements int, false_positive_rate float64) *BloomFilter {

	bloom := new(BloomFilter)
	bloom.m = CalculateM(excepted_elements, false_positive_rate)
	bloom.k = CalculateK(excepted_elements, bloom.m)
	bloom.podaci = make([]byte, bloom.m)
	bloom.fns = CreateHashFunctions(bloom.k)
	var buf = &bytes.Buffer{}
	bloom.encoder = gob.NewEncoder(buf)
	bloom.decoder = gob.NewDecoder(buf)
	return bloom
}

// kreira bloom filter od vec zadatih podataka koji ce se ucitavati iz fajla za vec unapred napravljene bloom filtere
func RecreateBloomFilter(m uint, k uint, fns []HashWithSeed, podaci []byte) *BloomFilter {

	bloom := new(BloomFilter)
	bloom.m = m
	bloom.k = k
	bloom.podaci = podaci
	bloom.fns = fns
	var buf = &bytes.Buffer{}
	bloom.encoder = gob.NewEncoder(buf)
	bloom.decoder = gob.NewDecoder(buf)
	return bloom
}

// dodaje element u bloom filter
func (bloom *BloomFilter) Add(data []byte) {
	var i uint64
	for _, fn := range bloom.fns {
		err := bloom.encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = bloom.decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		i = dfn.Hash(data)
		i = i % uint64(bloom.m)
		bloom.podaci[i] = 1
	}
}

// pretrazuje bloom filter i govori da li element postoji ili ne, moze reci da postoji element koji ne postoji
func (bloom *BloomFilter) Search(data []byte) bool {
	var i uint64
	for _, fn := range bloom.fns {
		err := bloom.encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = bloom.decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		i = dfn.Hash(data)
		i = i % uint64(bloom.m)
		if bloom.podaci[i] == 0 {
			return false
		}
	}
	return true
}
