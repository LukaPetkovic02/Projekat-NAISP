package wal

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/edsrzf/mmap-go"
)

type Podatak struct {
	key       string
	value     []byte
	tombstone byte
	timestamp int64
}

func (data *Podatak) PrintData() {
	fmt.Print("Key : " + data.key + " ; Value : " + string(data.value) + " | Timestamp : ")
	fmt.Print(data.timestamp)
	fmt.Print(" | Tombstone : ")
	fmt.Println(data.tombstone)
}

func (data *Podatak) DecodeToByte() []byte {
	upis := make([]byte, 29+(int64)(len(data.key))+(int64)(len(data.value)))

	b := make([]byte, TIMESTAMP_SIZE) //pretvaranje timestampa u niz bita
	binary.LittleEndian.PutUint64(b, (uint64)(data.timestamp))
	copy(upis[TIMESTAMP_START:], b)

	c := make([]byte, TOMBSTONE_SIZE) //tombstone
	c[0] = data.tombstone
	copy(upis[TOMBSTONE_START:], c)

	d := make([]byte, KEY_SIZE_SIZE)
	binary.LittleEndian.PutUint64(d, (uint64)(len(data.key))) //duzina kljuca
	copy(upis[KEY_SIZE_START:], d)

	e := make([]byte, VALUE_SIZE_SIZE)
	binary.LittleEndian.PutUint64(e, (uint64)(len(data.value))) //duzina vrednosti
	copy(upis[VALUE_SIZE_START:], e)

	copy(upis[KEY_START:], []byte(data.key))         //kljuc
	copy(upis[KEY_START+len(data.key):], data.value) //value

	a := make([]byte, CRC_SIZE)
	binary.LittleEndian.PutUint32(a, CRC32(upis[TIMESTAMP_START:KEY_START+len(data.key)+len(data.value)])) //crc
	copy(upis[CRC_START:], a)
	return upis
}

// brise fajl za zadati path fajla
func RemoveFile(s string) {
	e := os.Remove(s)
	if e != nil {
		log.Fatal(e)
	}
}

// brise n fajlova pocevsi od prvog zadatog
func RemoveNFilesStarting(start string, n int) {

	for i := 0; i < n; i++ {
		RemoveFile(start)
		start = NextFileString(start)
	}
}

// brise prvih n walova
func DeleteLog(n int) {
	current := startFile
	f, err := os.OpenFile(current, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Prvi fajl je prazan.") //ako ne mozemo da otvorimo prvi log jer ne postoji izlazimo iz brisanja
		return
	}

	err = f.Truncate(0) //ako nije praznimo ga
	if err != nil {
		log.Fatal(err)
	}

	next := NextFile(*f) //odredjujemo sledeci fajl
	f.Close()

	for i := 1; i < n; i++ { // prvih n fajlova praznimo
		f, err = os.OpenFile(next, os.O_RDWR, 0777)
		if err != nil { //ako neki i-ti fajl ne mozemo da otvorimo znaci da nema dovoljno fajlova za brisanje i mozemo da removujemo prvih i fajlova
			RemoveNFilesStarting(startFile, i)
			fmt.Println("Prekoracili smo fajlove za brisanje.")
			return
		}

		err = f.Truncate(0) // praznimo fajl i prelazimo na sledeci
		if err != nil {
			log.Fatal(err)
		}
		next = NextFile(*f)
		f.Close()
	}

	numOfRewritens := 0 //ovde pratimo koliko prvih fajlova smo popunili podacima iz kasnijih

	for true {
		f, err = os.OpenFile(next, os.O_RDWR, 0777)
		if err != nil { //ako nema vise fajlova za otvaranje brisemo od zadnjeg dodatog + 1 do zadnjeg fajla kojeg imamo
			fmt.Println("Prepisali smo sve fajlove")
			RemoveNFilesStarting(current, n)
			return
		}

		var ret [][]byte // ako smo ga otvorili uzimamo njegove podatke
		ret = ReadSegment(ret, *f)
		for i := 0; i < len(ret); i++ {
			WriteSegment(ret[i], current) //upisujemo ga u prvi prazan
		}

		err = f.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}
		next = NextFile(*f)
		current = NextFileString(current)
		f.Close()
		numOfRewritens += 1
	}
}

// zapisuje niz bitova u fajl
func WriteSegment(ret []byte, name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Print("Greska pri otvaranju " + name)
		log.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		fmt.Print("Nema stat :(")
		log.Fatal(err)
	}

	start := info.Size()
	errr := f.Truncate(start + (int64)(len(ret)))
	if errr != nil {
		log.Fatal(errr)
	}
	mmapFile, err := mmap.Map(f, mmap.RDWR, 0)

	copy(mmapFile[start:], ret)
	f.Close()
	mmapFile.Unmap()
}

// nalazi jedan element wala od date pozicije
func FindSegment(start int64, log []byte) []byte {
	keysize := log[start+KEY_SIZE_START : start+KEY_SIZE_START+KEY_SIZE_SIZE]
	k := binary.LittleEndian.Uint64(keysize)
	//fmt.Println("Key size: " + (string)(k))
	valuesize := log[start+VALUE_SIZE_START : start+VALUE_SIZE_START+VALUE_SIZE_SIZE]
	v := binary.LittleEndian.Uint64(valuesize)
	//fmt.Println("Value size: " + (string)(v))
	segment := log[start : start+29+(int64)(k+v)]
	return segment
}

// za zadati fajl i niz bita apendovace na niz bita svaki element wala iz tog fajla
func ReadSegment(ret [][]byte, f os.File) [][]byte {
	log, er := io.ReadAll(&f)
	if er != nil {
		fmt.Println("Greska pri citanju fajla.")
		panic(er)
	}
	var start int64 = 0
	for i := 0; start < (int64)(len(log)) && i < size_wala; i++ { //ucitavamo segmenat po segmenat
		ret = append(ret, FindSegment(start, log)) //find segment ce nam vratiti niz bita za sledeci segmenat od neke pozicije
		start += int64(len(ret[len(ret)-1]))       //povecavama start za  broj bita koje smo ucitali kao sledeci segment
	}
	return ret
}

// cita wal i vraca niz nizova bita gde svaki niz predstavlja neki element wala, oni se mogu dekodirati funkcijum encode to data da bi postali objekti Podatak
func ReadLog() [][]byte {

	var ret [][]byte

	f, err := os.OpenFile(startFile, os.O_RDWR, 0777)
	if err != nil {
		//fmt.Println("nema fajla :(")
		return ret
	}

	info, err := f.Stat()
	if err != nil {
		fmt.Println("Greska sa ucitavanjem stata.")
		log.Fatal(err)
	}

	for info.Size() != 0 {
		ret = ReadSegment(ret, *f)
		name := NextFile(*f)
		f.Close()
		f, err = os.OpenFile(name, os.O_RDWR, 0777)
		if err != nil {
			fmt.Println("Dosli smo do zadnjeg fajla.")
			break
		}
		info, err = f.Stat()
		if err != nil {
			fmt.Println("Greska sa ucitavanjem stata.")
			log.Fatal(err)
		}
		if info.Size() == 0 {
			fmt.Println("Dosli smo do prvog praznog fajla.")
			break
		}
	}
	f.Close()

	return ret
}

// vraca ime sledeceg fajla za zadato ime trenutnog fajla
func NextFileString(current string) string {

	num, er := strconv.Atoi(current[3:strings.Index(current, ".")])
	if er != nil {
		log.Fatal(er)
	}
	num += 1
	return "wal" + strconv.Itoa(num) + ".txt"
}

// vraca ime sledeceg fajla za zadat otvoreni fajl
func NextFile(f os.File) string {
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	name := info.Name()

	num, er := strconv.Atoi(name[3:strings.Index(name, ".")])
	if er != nil {
		log.Fatal(er)
	}

	num += 1
	return "wal" + strconv.Itoa(num) + ".txt"
}

// pretvara niz bajtova u element tipa Podatak
func EncodeToData(data []byte) Podatak {
	var ret Podatak
	keysize := data[KEY_SIZE_START : KEY_SIZE_START+KEY_SIZE_SIZE]
	ks := binary.LittleEndian.Uint64(keysize)

	valuesize := data[VALUE_SIZE_START : VALUE_SIZE_START+VALUE_SIZE_SIZE]
	vs := binary.LittleEndian.Uint64(valuesize)

	key := data[KEY_START : KEY_START+ks]
	ret.key = fmt.Sprintf("%s", key)

	ret.value = data[KEY_START+ks : KEY_START+ks+vs]

	timestamp := data[TIMESTAMP_START : TIMESTAMP_START+TIMESTAMP_SIZE]
	ret.timestamp = int64(binary.LittleEndian.Uint64(timestamp))

	tombstone := data[TOMBSTONE_START : TOMBSTONE_START+TOMBSTONE_SIZE]
	ret.tombstone = byte(tombstone[0])

	return ret
}

// pretvara kljuc vrednost i tombstone u niz bajtova i pravi im timestamp za momenat pravljenja
func DecodeToByte(key string, value []byte, t byte) []byte {
	upis := make([]byte, 29+(int64)(len(key))+(int64)(len(value)))

	b := make([]byte, TIMESTAMP_SIZE) //pretvaranje timestampa u niz bita
	now := time.Now().Unix()
	binary.LittleEndian.PutUint64(b, (uint64)(now))
	copy(upis[TIMESTAMP_START:], b)

	c := make([]byte, TOMBSTONE_SIZE) //tombstone
	c[0] = t
	copy(upis[TOMBSTONE_START:], c)

	d := make([]byte, KEY_SIZE_SIZE)
	binary.LittleEndian.PutUint64(d, (uint64)(len(key))) //duzina kljuca
	copy(upis[KEY_SIZE_START:], d)

	e := make([]byte, VALUE_SIZE_SIZE)
	binary.LittleEndian.PutUint64(e, (uint64)(len(value))) //duzina vrednosti
	copy(upis[VALUE_SIZE_START:], e)

	copy(upis[KEY_START:], []byte(key))    //kljuc
	copy(upis[KEY_START+len(key):], value) //value

	a := make([]byte, CRC_SIZE)
	binary.LittleEndian.PutUint32(a, CRC32(upis[TIMESTAMP_START:KEY_START+len(key)+len(value)])) //crc
	copy(upis[CRC_START:], a)
	return upis
}

// dodaje log
func AddLog(key string, value []byte, t byte, fileName string) {

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777) //otvaramo fajl
	if err != nil {
		fmt.Print("Greska pri otvaranju " + fileName)
		log.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		fmt.Print("Nema stat :(")
		log.Fatal(err)
	}

	start := info.Size() //pamtimo kraj fajla kako bi smo upisali na njega

	var ret [][]byte
	ret = ReadSegment(ret, *f)

	if len(ret) == size_wala { //ako je trenutni fajl popunjen
		name := NextFile(*f) //next segment nam vraca sledeci fajl u listi valova
		f.Close()
		AddLog(key, value, t, name) //pokusavamo da dodamo element u sledeci wal
		return
	}

	errr := f.Truncate(start + 29 + (int64)(len(key)) + (int64)(len(value))) //postavljamo velicinu na trenutnu + dovoljno mesta za sledeci segment
	if errr != nil {
		log.Fatal(errr)
	}

	mmapFile, err := mmap.Map(f, mmap.RDWR, 0)

	if err != nil {
		log.Fatal(err)
	}

	upis := DecodeToByte(key, value, t)

	copy(mmapFile[start:], upis)
	mmapFile.Unmap()
	f.Close()
}

var startFile string = "wal1.txt"
var size_wala int = 16

// func main() {

// 	addLog("bababoi", []byte("xbov<<<popovic"), 1, startFile)
// 	addLog("nibba", []byte("JoeBidenWakeUp"), 0, startFile)
// 	addLog("ker", []byte("bajajajajaj"), 1, startFile)
// 	addLog("smer", []byte("NIggleton"), 0, startFile)
// 	addLog("bababoi", []byte("Pekka ridge spam :))"), 1, startFile)
// 	addLog("123", []byte("Kill da hoe"), 0, startFile)
// 	addLog("222", []byte("Nocturn :0"), 1, startFile)
// 	addLog("s2e44r", []byte("Dantes goat"), 0, startFile)
// 	addLog("t1212t", []byte("Messi << Ronaldino"), 1, startFile)
// 	addLog("hhhhh", []byte("Muhamad Hanakin did 9/11"), 0, startFile)
// 	addLog("asasasas", []byte("Killer Queen 3rd bomb"), 1, startFile)
// 	addLog("777", []byte("ebin gejms"), 0, startFile)
// 	addLog("989898", []byte("Venic Krvat"), 1, startFile)
// 	addLog("1skija", []byte("Nigga Balls"), 0, startFile)
// 	addLog("tum", []byte("Can I put my balls in yo jaws"), 1, startFile)
// 	addLog("u don no", []byte("Balls in yo jaws"), 0, startFile)
// 	var m Podatak
// 	x := readLog()
// 	fmt.Println("*************************************************************")
// 	for i := 0; i < len(x); i++ {
// 		m = encodeToData(x[i])
// 		m.printData()
// 		fmt.Println()
// 	}
// 	deleteLog(10)
// }
