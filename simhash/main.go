package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

var mapa1 = make(map[string]int)
var mapa2 = make(map[string]int)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ToBinary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}
func ProcitajFajl(mapa map[string]int, putanja string) {
	file, err := os.OpenFile(putanja, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	bufferedReader := bufio.NewReader(file)
	for {
		var lista2reci []string
		var prvarec string = ""
		var drugarec string = ""
		dataString, err1 := bufferedReader.ReadString(' ')
		if err1 != nil {
			break
		}

		rec := dataString[:len(dataString)-1]                 //brisemo razmak
		if rec[len(rec)-1] == '.' || rec[len(rec)-1] == ',' { //i tacku ili zarez
			rec = rec[:len(rec)-1]
		}
		if strings.Contains(rec, "\n") {
			lista2reci = strings.Split(rec, ".")
			prvarec = lista2reci[0]
			drugarec = lista2reci[1][2:]
			var nasaoprvu bool = false
			var nasaodrugu bool = false
			if prvarec != "" && drugarec != "" {
				for kljuc, vrednost := range mapa {
					if kljuc == prvarec {
						mapa[kljuc] = vrednost + 1
						nasaoprvu = true
					} else if kljuc == drugarec {
						mapa[kljuc] = vrednost + 1
						nasaodrugu = true
					}
				}
				if !nasaoprvu {
					mapa[prvarec] = 1
				}
				if !nasaodrugu {
					mapa[drugarec] = 1
				}
			}
		} else {
			var nasao bool = false
			for kljuc, vrednost := range mapa {
				if kljuc == rec {
					mapa[kljuc] = vrednost + 1
					nasao = true
					break
				}
			}
			if !nasao {
				mapa[rec] = 1
			}
		}
	}

	file.Close()
}
func SimHash(mapa map[string]int) [256]int {
	var lista [256]int
	for i := 0; i < 256; i++ {
		for k, v := range mapa {
			strhash := ToBinary(GetMD5Hash(k))
			itacifra := int(strhash[i]) - 48
			if itacifra == 0 {
				itacifra = -1
			}
			lista[i] += itacifra * v
		}
		if lista[i] > 0 {
			lista[i] = 1
		} else {
			lista[i] = 0
		}
	}
	return lista
}
func HemingvejovaUdaljenost(lista1 [256]int, lista2 [256]int) int {
	var suma int = 0
	for i := 0; i < len(lista1); i++ {
		if (lista1[i] == 0 && lista2[i] == 1) || (lista1[i] == 1 && lista2[i] == 0) {
			suma += 1
		}
	}
	return suma
}
func main() {
	fmt.Println(GetMD5Hash("hello"))           //svaki podatak pretvori u 32 hex cifre
	fmt.Println(ToBinary(GetMD5Hash("hello"))) //ili 32*8 binarnih
	ProcitajFajl(mapa1, "tekst1.txt")          //cita fajl i pravi mapu sa stringom i brojem njegovog ponavljanja u tekstu
	ProcitajFajl(mapa2, "tekst2.txt")
	lista1 := SimHash(mapa1)
	lista2 := SimHash(mapa2)
	//fmt.Println(lista1)
	//fmt.Println(lista2)
	fmt.Println("Hemingvejovo rastojanje:", HemingvejovaUdaljenost(lista1, lista2))
	//for k, v := range mapa1 {
	//	fmt.Print("mapa1[", k, "]=")
	//	fmt.Println(v)
	//}
	//fmt.Println("-------------------------------------------------")
	//for k, v := range mapa2 {
	//	fmt.Print("mapa2[", k, "]=")
	//	fmt.Println(v)
	//}

}
