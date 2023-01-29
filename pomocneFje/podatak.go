package pomocnefje

import (
	"encoding/binary"
	"fmt"
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
