package utils

import (
	"strconv"
	"strings"
)

type ParsedFileName struct {
	Level     int
	Timestamp int64
}

func ParseFileName(fileName string) ParsedFileName {
	var a = strings.Split(fileName, "_")
	var level, _ = strconv.Atoi(a[0])
	var timestamp, _ = strconv.Atoi(strings.Split(a[1], ".")[0])
	return ParsedFileName{
		Level:     level,
		Timestamp: int64(timestamp),
	}
}
