package cms

import (
	"bytes"
	"encoding/gob"
)

type CMS struct {
	fns     []HashWithSeed //sve hash funkcije kojih ima k u bloom filteru
	podaci  [][]byte       //niz podataka koji nam govori da li element postoji ili ne
	m       uint           //ocekivani broj elemenata
	k       uint           //zeljeni broj hash funkcija
	encoder *gob.Encoder   //enkoder za hash funkcije
	decoder *gob.Decoder   //dekoder za hash funkcije
}

func NewCMS(false_positive_rate float64) *CMS {

	cms := new(CMS)
	cms.m = CalculateM(false_positive_rate)
	cms.k = CalculateK(false_positive_rate)
	cms.podaci = make([][]byte, cms.k)
	cms.fns = CreateHashFunctions(cms.k)
	var buf = &bytes.Buffer{}
	cms.encoder = gob.NewEncoder(buf)
	cms.decoder = gob.NewDecoder(buf)
	return cms
}

// kreira bloom filter od vec zadatih podataka koji ce se ucitavati iz fajla za vec unapred napravljene bloom filtere
func RecreateCMS(m uint, k uint, fns []HashWithSeed, podaci [][]byte) *CMS {

	cms := new(CMS)
	cms.m = m
	cms.k = k
	cms.podaci = podaci
	cms.fns = fns
	var buf = &bytes.Buffer{}
	cms.encoder = gob.NewEncoder(buf)
	cms.decoder = gob.NewDecoder(buf)
	return cms
}

// dodaje element u bloom filter
func (cms *CMS) Add(data []byte) {
	var j uint64
	for i, fn := range cms.fns {
		err := cms.encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = cms.decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(cms.m)
		cms.podaci[i][j] += 1
	}
}

// pretrazuje bloom filter i govori da li element postoji ili ne, moze reci da postoji element koji ne postoji
func (cms *CMS) Ucestalost(data []byte) byte {
	var l byte = 0
	var j uint64
	niz := make([]byte, 0)
	for i, fn := range cms.fns {
		err := cms.encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &HashWithSeed{}
		err = cms.decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(cms.m)
		niz = append(niz, cms.podaci[i][j])
	}

	l = niz[0]
	for i, val := range niz {
		if i != 0 && val < l {
			l = val
		}
	}
	return l
}
