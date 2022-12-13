package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/edsrzf/mmap-go"
)

// func deleteLog(f os.File) {

// }

func findSegment(start int64, log []byte) []byte {
	keysize := log[start+KEY_SIZE_START : start+KEY_SIZE_START+8]
	k := binary.LittleEndian.Uint64(keysize)
	valuesize := log[start+VALUE_SIZE_START : start+VALUE_SIZE_START+8]
	v := binary.LittleEndian.Uint64(valuesize)
	segment := log[start : start+37+(int64)(k+v)]
	return segment
}

func readSegment(ret [][]byte, f os.File) [][]byte {
	log, er := io.ReadAll(&f)
	if er != nil {
		panic(er)
	}
	var start int64 = 0
	for i := 0; start < (int64)(len(log)) && i < 3; i++ {
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
		return ret
	}

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	for info.Size() != 0 {
		ret = readSegment(ret, *f)
		name := nextSegment(*f)
		f.Close()
		f, err = os.OpenFile(name, os.O_RDWR, 0777)
		if err != nil {
			break
		}
		info, err = f.Stat()
		if err != nil {
			log.Fatal(err)
		}
		if info.Size() == 0 {
			break
		}
	}
	return ret
}

func nextSegment(f os.File) string {
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	name := info.Name()
	num := 0
	for i := 3; i < len(name) && name[i] != '.'; i++ {
		num += num*10 + (int)(name[i])
	}
	return "wal" + (string)(num) + ".txt"
}

func addLog(key string, value []byte, t byte, startFile string) {

	f, err := os.OpenFile(startFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Print(startFile)
		log.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	start := info.Size()
	var ret [][]byte
	fmt.Println(len(readSegment(ret, *f)))
	if len(readSegment(ret, *f)) >= 3 {
		name := nextSegment(*f)
		f.Close()
		addLog(key, value, t, name)
		return

	}

	upis := make([]byte, 37+(int64)(len(key))+(int64)(len(value)))

	errr := f.Truncate(start + 37 + (int64)(len(key)) + (int64)(len(value)))
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

func main() {
	f, err := os.OpenFile(startFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Print("dada1")
		log.Fatal(err)
	}
	defer f.Close()

	addLog("bababoi", []byte("xbov<<<popovic"), 1, startFile)
	addLog("nibba", []byte("JoeBidenWakeUp"), 0, startFile)
	addLog("bababoi", []byte("xbov<<<popovic"), 1, startFile)
	addLog("nibba", []byte("JoeBidenWakeUp"), 0, startFile)
	readLog()
}
