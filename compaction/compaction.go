package compaction

import (
	"github.com/LukaPetkovicSV16/Projekat-NAISP/types"
)

// spaja dve liste rekorda u novu
func Merge(sstable1 []types.Record, sstable2 []types.Record) []types.Record {
	var i int = 0
	var j int = 0
	var ret []types.Record
	for i < len(sstable1) || j < len(sstable2) {

		if i == len(sstable1) {
			ret = append(ret, sstable2[j])
			j++
			continue
		} else if j == len(sstable2) {
			ret = append(ret, sstable1[i])
			i++
			continue
		}

		if sstable1[i].Key == sstable2[j].Key {
			if sstable1[i].Timestamp > sstable2[j].Timestamp {
				ret = append(ret, sstable1[i])
			} else {
				ret = append(ret, sstable2[j])
			}
			i++
			j++
		} else if sstable1[i].Key < sstable2[j].Key {
			ret = append(ret, sstable1[i])
			i++
		} else {
			ret = append(ret, sstable2[j])
			j++
		}
	}

	return ret
}
