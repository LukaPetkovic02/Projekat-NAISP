package memtable

import (
	Importi "projekat/utils"
)

type MemTable interface {
	SearchData(key string) *Importi.Podatak
	AllDataSortedBegin() []Importi.Podatak
	Put(podatak Importi.Podatak) []Importi.Podatak
}
