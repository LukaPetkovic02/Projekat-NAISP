package hll

import (
	"math"
	"math/bits"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/utils"
)

const (
	HLL_MIN_PRECISION = 4
	HLL_MAX_PRECISION = 16
)

type HLL struct {
	m   uint64
	p   uint8
	reg []uint8
}

func NewHLL(p1 uint8) *HLL {
	p := new(HLL)
	p.p = p1
	p.m = uint64(math.Pow(2, float64(p1)))
	p.reg = make([]uint8, p.m)
	return p
}

func (hll *HLL) Add(data []byte) {
	br := &utils.HashWithSeed{}
	data1 := br.Hash(data)
	prvecifre := data1 >> (64 - hll.p)

	brnula := bits.TrailingZeros64(data1)
	if hll.reg[prvecifre] < uint8(brnula) {
		hll.reg[prvecifre] = uint8(brnula)
	}
}

func (hll *HLL) Estimate() float64 {
	sum := 0.0
	for _, val := range hll.reg {
		sum += math.Pow(math.Pow(2.0, float64(val)), -1)
	}

	alpha := 0.7213 / (1.0 + 1.079/float64(hll.m))
	estimation := alpha * math.Pow(float64(hll.m), 2.0) / sum
	emptyRegs := hll.emptyCount()
	if estimation <= 2.5*float64(hll.m) { // do small range correction
		if emptyRegs > 0 {
			estimation = float64(hll.m) * math.Log(float64(hll.m)/float64(emptyRegs))
		}
	} else if estimation > 1/30.0*math.Pow(2.0, 32.0) { // do large range correction
		estimation = -math.Pow(2.0, 32.0) * math.Log(1.0-estimation/math.Pow(2.0, 32.0))
	}
	return estimation
}

func (hll *HLL) emptyCount() int {
	sum := 0
	for _, val := range hll.reg {
		if val == 0 {
			sum++
		}
	}
	return sum
}
