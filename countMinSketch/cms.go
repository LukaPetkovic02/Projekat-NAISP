package countMS

import (
	"bytes"
	"encoding/gob"
	importi "projekat/utils"
)

type CMS struct {
	Fns     []importi.HashWithSeed //sve hash funkcije kojih ima k u bloom filteru
	Podaci  [][]byte               //niz podataka koji nam govori da li element postoji ili ne
	M       uint                   //ocekivani broj elemenata
	K       uint                   //zeljeni broj hash funkcija
	Encoder *gob.Encoder           //enkoder za hash funkcije
	Decoder *gob.Decoder           //dekoder za hash funkcije
}

func NewCMS(false_positive_rate float64) *CMS {

	cms := new(CMS)
	cms.M = CalculateMC(false_positive_rate)
	cms.K = CalculateKC(false_positive_rate)
	cms.Podaci = make([][]byte, cms.K)
	for i := range cms.Podaci {
		cms.Podaci[i] = make([]byte, cms.M)
	}
	cms.Fns = importi.CreateHashFunctions(cms.K)
	var buf = &bytes.Buffer{}
	cms.Encoder = gob.NewEncoder(buf)
	cms.Decoder = gob.NewDecoder(buf)
	return cms
}

// kreira bloom filter od vec zadatih podataka koji ce se ucitavati iz fajla za vec unapred napravljene bloom filtere
func RecreateCMS(m uint, k uint, fns []importi.HashWithSeed, podaci [][]byte) *CMS {

	cms := new(CMS)
	cms.M = m
	cms.K = k
	cms.Podaci = podaci
	cms.Fns = fns
	var buf = &bytes.Buffer{}
	cms.Encoder = gob.NewEncoder(buf)
	cms.Decoder = gob.NewDecoder(buf)
	return cms
}

// dodaje element u bloom filter
func (cms *CMS) Add(data []byte) {
	var j uint64
	for i, fn := range cms.Fns {
		err := cms.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &importi.HashWithSeed{}
		err = cms.Decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(cms.M)
		cms.Podaci[i][j] += 1
	}
}

// pretrazuje bloom filter i govori da li element postoji ili ne, moze reci da postoji element koji ne postoji
func (cms *CMS) Ucestalost(data []byte) byte {
	var l byte = 0
	var j uint64
	niz := make([]byte, 0)
	for i, fn := range cms.Fns {
		err := cms.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &importi.HashWithSeed{}
		err = cms.Decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(cms.M)
		niz = append(niz, cms.Podaci[i][j])
	}

	l = niz[0]
	for i, val := range niz {
		if i != 0 && val < l {
			l = val
		}
	}
	return l
}
