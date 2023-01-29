package pomocnefje

import (
	"encoding/binary"
	"fmt"
	"time"
)

type Podatak struct {
	Key       string
	Value     []byte
	Tombstone byte
	Timestamp int64
}

func (data *Podatak) PrintData() {
	fmt.Print("Key : " + data.Key + " ; Value : " + string(data.Value) + " | Timestamp : ")
	fmt.Print(data.Timestamp)
	fmt.Print(" | Tombstone : ")
	fmt.Println(data.Tombstone)
}

func (data *Podatak) DecodeToByte() []byte {
	upis := make([]byte, 29+(int64)(len(data.Key))+(int64)(len(data.Value)))

	b := make([]byte, TIMESTAMP_SIZE) //pretvaranje timestampa u niz bita
	binary.LittleEndian.PutUint64(b, (uint64)(data.Timestamp))
	copy(upis[TIMESTAMP_START:], b)

	c := make([]byte, TOMBSTONE_SIZE) //tombstone
	c[0] = data.Tombstone
	copy(upis[TOMBSTONE_START:], c)

	d := make([]byte, KEY_SIZE_SIZE)
	binary.LittleEndian.PutUint64(d, (uint64)(len(data.Key))) //duzina kljuca
	copy(upis[KEY_SIZE_START:], d)

	e := make([]byte, VALUE_SIZE_SIZE)
	binary.LittleEndian.PutUint64(e, (uint64)(len(data.Value))) //duzina vrednosti
	copy(upis[VALUE_SIZE_START:], e)

	copy(upis[KEY_START:], []byte(data.Key))         //kljuc
	copy(upis[KEY_START+len(data.Key):], data.Value) //value

	a := make([]byte, CRC_SIZE)
	binary.LittleEndian.PutUint32(a, CRC32(upis[TIMESTAMP_START:KEY_START+len(data.Key)+len(data.Value)])) //crc
	copy(upis[CRC_START:], a)
	return upis
}

// pretvara niz bajtova u element tipa Podatak
func EncodeToData(data []byte) Podatak {
	var ret Podatak
	keysize := data[KEY_SIZE_START : KEY_SIZE_START+KEY_SIZE_SIZE]
	ks := binary.LittleEndian.Uint64(keysize)

	valuesize := data[VALUE_SIZE_START : VALUE_SIZE_START+VALUE_SIZE_SIZE]
	vs := binary.LittleEndian.Uint64(valuesize)

	key := data[KEY_START : KEY_START+ks]
	ret.Key = fmt.Sprintf("%s", key)

	ret.Value = data[KEY_START+ks : KEY_START+ks+vs]

	timestamp := data[TIMESTAMP_START : TIMESTAMP_START+TIMESTAMP_SIZE]
	ret.Timestamp = int64(binary.LittleEndian.Uint64(timestamp))

	tombstone := data[TOMBSTONE_START : TOMBSTONE_START+TOMBSTONE_SIZE]
	ret.Tombstone = byte(tombstone[0])

	return ret
}

func NewPodatak(Key string, Value []byte, Tombstone byte) Podatak {
	var ret Podatak

	ret.Key = Key

	ret.Value = Value

	now := time.Now().Unix()
	ret.Timestamp = int64(now)

	ret.Tombstone = Tombstone

	return ret
}
