package cms

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"

	"math"

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

func CalculateM(epsilon float64) uint {
	return uint(math.Ceil(math.E / epsilon))
}

func CalculateK(delta float64) uint {
	return uint(math.Ceil(math.Log(math.E / delta)))
}

type CountMinSketch struct {
	Fns     []utils.HashWithSeed
	Encoder *gob.Encoder
	Decoder *gob.Decoder
	M       uint
	K       uint
	Data    [][]byte
}

func CreateCMS(epsilon float64, delta float64) *CountMinSketch {
	cms := new(CountMinSketch)
	cms.M = CalculateM(epsilon)
	cms.K = CalculateK(delta)
	var buff = &bytes.Buffer{}
	cms.Encoder = gob.NewEncoder(buff)
	cms.Decoder = gob.NewDecoder(buff)
	cms.Fns = utils.CreateHashFunctions(cms.K)
	matrica := make([][]byte, cms.K) //KxM matrix
	for i := range matrica {
		matrica[i] = make([]byte, cms.M)
	}
	cms.Data = matrica
	return cms
}

func (c *CountMinSketch) Add(data []byte) {
	var j uint64
	for i, fn := range c.Fns {
		err := c.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &utils.HashWithSeed{}
		err = c.Decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(c.M)
		c.Data[i][j] += 1
	}
}

func (c *CountMinSketch) Frequency(data []byte) byte {
	var l byte = 0
	var j uint64
	niz := make([]byte, 0)
	for i, fn := range c.Fns {
		err := c.Encoder.Encode(fn)
		if err != nil {
			panic(err)
		}
		dfn := &utils.HashWithSeed{}
		err = c.Decoder.Decode(dfn)
		if err != nil {
			panic(err)
		}
		j = dfn.Hash(data)
		j = j % uint64(c.M)
		niz = append(niz, c.Data[i][j])
	}
	l = niz[0]
	for i, val := range niz {
		if i != 0 && val < l {
			l = val
		}
	}
	return l
}

func (c *CountMinSketch) Serialize() []byte {
	var serializedCms = new(bytes.Buffer)

	binary.Write(serializedCms, binary.LittleEndian, uint64(c.M))
	binary.Write(serializedCms, binary.LittleEndian, uint64(c.K))

	for _, fn := range c.Fns {
		binary.Write(serializedCms, binary.LittleEndian, fn.Seed)
	}
	for _, f := range c.Data {
		binary.Write(serializedCms, binary.LittleEndian, f)
	}
	//binary.Write(serializedCms, binary.LittleEndian, c.Data)
	return serializedCms.Bytes()
}

func Deserialize(data []byte) *CountMinSketch {
	cms := new(CountMinSketch)
	m := data[M_START : M_START+M_SIZE]
	cms.M = uint(binary.LittleEndian.Uint64(m))
	k := data[K_START : K_START+K_SIZE]
	cms.K = uint(binary.LittleEndian.Uint64(k))
	for i := 0; i < int(cms.K); i++ { //k hash functions
		hash := new(utils.HashWithSeed)
		hash.Seed = data[HASH_START+HASH_SIZE*i : HASH_START+HASH_SIZE*(i+1)]
		cms.Fns = append(cms.Fns, *hash)
	}

	data_start := HASH_START + HASH_SIZE*int(cms.K)
	//posle hash funkcija stoji k*m bajtova koje treba ucitati u matricu
	for j := 0; j < int(cms.K); j++ { //k puta appendujem red duzine m
		cms.Data = append(cms.Data, data[data_start:data_start+int(cms.M)])
		data_start = data_start + int(cms.M)
	}
	var buff = &bytes.Buffer{}
	cms.Encoder = gob.NewEncoder(buff)
	cms.Decoder = gob.NewDecoder(buff)
	return cms
}
