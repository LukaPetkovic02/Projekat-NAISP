package main

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

func removeFile(s string) {
	e := os.Remove(s)
	if e != nil {
		log.Fatal(e)
	}
}

func removeNFilesStarting(start string, n int) {

	for i := 0; i < n; i++ {
		removeFile(start)
		start = nextSegmentString(start)
	}
}

func deleteLog(n int) {
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

	next := nextSegment(*f) //odredjujemo sledeci fajl
	f.Close()

	for i := 1; i < n; i++ { // prvih n fajlova praznimo
		f, err = os.OpenFile(next, os.O_RDWR, 0777)
		if err != nil { //ako neki i-ti fajl ne mozemo da otvorimo znaci da nema dovoljno fajlova za brisanje i mozemo da removujemo prvih i fajlova
			removeNFilesStarting(startFile, i)
			fmt.Println("Prekoracili smo fajlove za brisanje.")
			return
		}

		err = f.Truncate(0) // praznimo fajl i prelazimo na sledeci
		if err != nil {
			log.Fatal(err)
		}
		next = nextSegment(*f)
		f.Close()
	}

	numOfRewritens := 0 //ovde pratimo koliko prvih fajlova smo popunili podacima iz kasnijih

	for true {
		f, err = os.OpenFile(next, os.O_RDWR, 0777)
		if err != nil { //ako nema vise fajlova za otvaranje brisemo od zadnjeg dodatog + 1 do zadnjeg fajla kojeg imamo
			fmt.Println("Prepisali smo sve fajlove")
			removeNFilesStarting(current, n)
			return
		}

		var ret [][]byte // ako smo ga otvorili uzimamo njegove podatke
		ret = readSegment(ret, *f)
		for i := 0; i < len(ret); i++ {
			writeSegment(ret[i], current) //upisujemo ga u prvi prazan
		}

		err = f.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}
		next = nextSegment(*f)
		current = nextSegmentString(current)
		f.Close()
		numOfRewritens += 1
	}
}

func writeSegment(ret []byte, name string) {
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

func findSegment(start int64, log []byte) []byte {
	keysize := log[start+KEY_SIZE_START : start+KEY_SIZE_START+8]
	k := binary.LittleEndian.Uint64(keysize)
	//fmt.Println("Key size: " + (string)(k))
	valuesize := log[start+VALUE_SIZE_START : start+VALUE_SIZE_START+8]
	v := binary.LittleEndian.Uint64(valuesize)
	//fmt.Println("Value size: " + (string)(v))
	segment := log[start : start+29+(int64)(k+v)]
	return segment
}

func readSegment(ret [][]byte, f os.File) [][]byte {
	log, er := io.ReadAll(&f)
	if er != nil {
		fmt.Println("Greska pri citanju fajla.")
		panic(er)
	}
	var start int64 = 0
	for i := 0; start < (int64)(len(log)) && i < size_wala; i++ {
		ret = append(ret, findSegment(start, log))
		start += int64(len(ret[len(ret)-1]))
		//fmt.Println(ret[len(ret)-1])
	}
	return ret
}

func readLog() [][]byte {

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
		ret = readSegment(ret, *f)
		name := nextSegment(*f)
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
	return ret
}

func nextSegmentString(current string) string {

	num, er := strconv.Atoi(current[3:strings.Index(current, ".")])
	if er != nil {
		log.Fatal(er)
	}
	num += 1
	return "wal" + strconv.Itoa(num) + ".txt"
}

func nextSegment(f os.File) string {
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

func addLog(key string, value []byte, t byte, fileName string) {

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Print("Greska pri otvaranju " + fileName)
		log.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		fmt.Print("Nema stat :(")
		log.Fatal(err)
	}

	start := info.Size()

	var ret [][]byte
	ret = readSegment(ret, *f)

	if len(ret) == size_wala {
		name := nextSegment(*f)
		f.Close()
		addLog(key, value, t, name)
		return
	}

	upis := make([]byte, 29+(int64)(len(key))+(int64)(len(value)))

	errr := f.Truncate(start + 29 + (int64)(len(key)) + (int64)(len(value)))
	if errr != nil {
		log.Fatal(errr)
	}

	mmapFile, err := mmap.Map(f, mmap.RDWR, 0)
	defer mmapFile.Unmap()

	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 8)
	now := time.Now().Unix()
	binary.LittleEndian.PutUint64(b, (uint64)(now))
	copy(upis[TIMESTAMP_START:], b)

	c := make([]byte, 1)
	c[0] = t
	copy(upis[TOMBSTONE_START:], c)

	d := make([]byte, 8)
	binary.LittleEndian.PutUint64(d, (uint64)(len(key)))
	copy(upis[KEY_SIZE_START:], d)

	e := make([]byte, 8)
	binary.LittleEndian.PutUint64(e, (uint64)(len(value)))
	copy(upis[VALUE_SIZE_START:], e)

	copy(upis[KEY_START:], []byte(key))
	copy(upis[KEY_START+len(key):], value)

	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, CRC32(upis[TIMESTAMP_START:KEY_START+len(key)+len(value)]))
	copy(upis[CRC_START:], a)

	copy(mmapFile[start:], upis)
	f.Close()
}

var startFile string = "wal1.txt"
var size_wala int = 3

func main() {

	// addLog("bababoi", []byte("xbov<<<popovic"), 1, startFile)
	// addLog("nibba", []byte("JoeBidenWakeUp"), 0, startFile)
	// addLog("ker", []byte("bajajajajaj"), 1, startFile)
	// addLog("smer", []byte("NIggleton"), 0, startFile)
	// addLog("bababoi", []byte("Pekka ridge spam :))"), 1, startFile)
	// addLog("123", []byte("Kill da hoe"), 0, startFile)
	// addLog("222", []byte("Nocturn :0"), 1, startFile)
	// addLog("s2e44r", []byte("Dantes goat"), 0, startFile)
	// addLog("t1212t", []byte("Messi << Ronaldino"), 1, startFile)
	// addLog("hhhhh", []byte("Muhamad Hanakin did 9/11"), 0, startFile)
	// addLog("asasasas", []byte("Killer Queen 3rd bomb"), 1, startFile)
	// addLog("777", []byte("ebin gejms"), 0, startFile)
	// addLog("989898", []byte("Venic Krvat"), 1, startFile)
	// addLog("1skija", []byte("Nigga Balls"), 0, startFile)
	// addLog("tum", []byte("Can I put my balls in yo jaws"), 1, startFile)
	// addLog("u don no", []byte("Balls in yo jaws"), 0, startFile)
	// x := readLog()
	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(x[i])
	// 	fmt.Println(len(x[i]))
	// }
	// deleteLog(10)

}
