package bloomFilter

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"os"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/utils"
)

const (
	M_START    = 0
	M_SIZE     = 8
	K_START    = M_START + M_SIZE
	K_SIZE     = 8
	HASH_START = K_START + K_SIZE
	HASH_SIZE  = 32
)

type BloomFilter struct {
	Fns     []utils.HashWithSeed //sve hash funkcije kojih ima k u bloom filteru
	Podaci  []byte               //niz podataka koji nam govori da li element postoji ili ne
	M       uint                 //ocekivani broj elemenata
	K       uint                 //zeljeni broj hash funkcija
	Encoder *gob.Encoder         //enkoder za hash funkcije
	Decoder *gob.Decoder         //dekoder za hash funkcije
}

// kreira novi bloom filter za ocekivani br elemenata i false positive rate
func CreateBloomFilter(excepted_elements int, false_positive_rate float64) *BloomFilter {

	bloom := new(BloomFilter)
	bloom.M = CalculateM(excepted_elements, false_positive_rate)
	bloom.K = CalculateK(excepted_elements, bloom.M)
	bloom.Podaci = make([]byte, bloom.M)
	bloom.Fns = utils.CreateHashFunctions(bloom.K)
	var buf = &bytes.Buffer{}
	bloom.Encoder = gob.NewEncoder(buf)
	bloom.Decoder = gob.NewDecoder(buf)
	return bloom
}

// kreira bloom filter od vec zadatih podataka koji ce se ucitavati iz fajla za vec unapred napravljene bloom filtere
func RecreateBloomFilterBloomFilter(m uint, k uint, fns []utils.HashWithSeed, podaci []byte) *BloomFilter {

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
		dfn := &utils.HashWithSeed{}
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
func (bloom *BloomFilter) Get(data []byte) bool {
	var i uint64
	for _, fn := range bloom.Fns {
		err := bloom.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &utils.HashWithSeed{}
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

func (bloom *BloomFilter) Serialize() []byte {
	var serializedBloom = new(bytes.Buffer)

	binary.Write(serializedBloom, binary.LittleEndian, uint64(bloom.M))
	binary.Write(serializedBloom, binary.LittleEndian, uint64(bloom.K))

	for _, fn := range bloom.Fns {
		binary.Write(serializedBloom, binary.LittleEndian, fn.Seed)
	}

	binary.Write(serializedBloom, binary.LittleEndian, bloom.Podaci)
	return serializedBloom.Bytes()
}

// Racunica za kraj bita ucitanog bloom filtera 16 + int(bloom.K)*32 + int(bloom.M) kako bi smo znali gde se dalje treba pozicionirati u fajlu
func Deserialize(data []byte) *BloomFilter {
	bloom := new(BloomFilter)

	m := data[M_START : M_START+M_SIZE]
	bloom.M = uint(binary.LittleEndian.Uint64(m))

	k := data[K_START : K_START+K_SIZE]
	bloom.K = uint(binary.LittleEndian.Uint64(k))

	for i := 0; i < int(bloom.K); i++ {
		hash := new(utils.HashWithSeed)
		hash.Seed = data[HASH_START+HASH_SIZE*i : HASH_START+HASH_SIZE*(i+1)]
		bloom.Fns = append(bloom.Fns, *hash)
	}

	podaci_start := HASH_START + HASH_SIZE*int(bloom.K)
	bloom.Podaci = data[podaci_start : podaci_start+int(bloom.M)]

	return bloom
}

func ReadFromFile(file *os.File) *BloomFilter {
	var bloom = new(BloomFilter)
	var m uint64
	var k uint64
	var b = make([]byte, M_SIZE)
	file.Read(b)
	m = binary.LittleEndian.Uint64(b)
	bloom.M = uint(m)
	bloom.M = uint(m)
	b = make([]byte, K_SIZE)
	file.Read(b)
	k = binary.LittleEndian.Uint64(b)
	bloom.K = uint(k)
	bloom.Fns = make([]utils.HashWithSeed, bloom.K)
	for i := 0; i < int(bloom.K); i++ {
		b = make([]byte, HASH_SIZE)
		file.Read(b)
		bloom.Fns[i].Seed = b
	}
	bloom.Podaci = make([]byte, bloom.M)
	var buf = &bytes.Buffer{}
	bloom.Encoder = gob.NewEncoder(buf)
	bloom.Decoder = gob.NewDecoder(buf)
	file.Read(bloom.Podaci)
	return bloom
}
