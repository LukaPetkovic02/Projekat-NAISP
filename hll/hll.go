package hll

import (
	"math"
	"math/bits"
	importi "projekat/utils"
)

const (
	HLL_MIN_PRECISION = 4
	HLL_MAX_PRECISION = 16
)

type HLL struct {
	M   uint64
	P   uint8
	Reg []uint8
}

func MakeHLL(p1 uint8) *HLL {
	p := new(HLL)
	p.P = p1
	p.M = uint64(math.Pow(2, float64(p1)))
	p.Reg = make([]uint8, p.M)
	return p
}

func (hll *HLL) Add(data []byte) {
	br := &importi.HashWithSeed{}
	data1 := br.Hash(data)
	prvecifre := data1 >> (64 - hll.P)

	brnula := bits.TrailingZeros64(data1)
	if hll.Reg[prvecifre] < uint8(brnula) {
		hll.Reg[prvecifre] = uint8(brnula)
	}
}

func (hll *HLL) Estimate() float64 {
	sum := 0.0
	for _, val := range hll.Reg {
		sum += math.Pow(math.Pow(2.0, float64(val)), -1)
	}

	alpha := 0.7213 / (1.0 + 1.079/float64(hll.M))
	estimation := alpha * math.Pow(float64(hll.M), 2.0) / sum
	emptyRegs := hll.EmptyCount()
	if estimation <= 2.5*float64(hll.M) { // do small range correction
		if emptyRegs > 0 {
			estimation = float64(hll.M) * math.Log(float64(hll.M)/float64(emptyRegs))
		}
	} else if estimation > 1/30.0*math.Pow(2.0, 32.0) { // do large range correction
		estimation = -math.Pow(2.0, 32.0) * math.Log(1.0-estimation/math.Pow(2.0, 32.0))
	}
	return estimation
}

func (hll *HLL) EmptyCount() int {
	sum := 0
	for _, val := range hll.Reg {
		if val == 0 {
			sum++
		}
	}
	return sum
}
