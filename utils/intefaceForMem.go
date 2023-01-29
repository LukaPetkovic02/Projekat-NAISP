package pomocnefje

type memTableInterface interface {
	Put(data Podatak) []Podatak
	SearchData(Key string) Podatak
	GetAllData() []Podatak
}
