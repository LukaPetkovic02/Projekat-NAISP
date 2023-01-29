package pomocnefje

type memTableInterface interface {
	Put(data Podatak) []Podatak
	SearchData() Podatak
	GetAllData() []Podatak
}
