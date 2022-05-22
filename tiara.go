package ekakao

import (
	"bytes"
	"math/rand"
	"regexp"
	"time"
)

var seedKey = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

var reg, _ = regexp.Compile(`[TZ\-:.]`)

func ShortenID(t int) string {
	var b bytes.Buffer
	for i := 0; i < t; i++ {
		n := rand.Int() % len(seedKey)
		b.WriteString(seedKey[n])
	}
	return b.String()
}

func RandomNumericString(t int) string {
	var b bytes.Buffer
	for i := 0; i < t; i++ {
		n := rand.Int()%10 + 48
		b.WriteString(string(rune(n)))
	}
	return b.String()
}

func CurrentTimeStamp() string {
	t := time.Now().Add(time.Hour * 9)
	s := t.Format("2006-01-02T15:04:05.999")
	return reg.ReplaceAllString(s, "")[2:]
}

func GenerateRandomUUIDWithDateNumber() string {
	var b bytes.Buffer
	b.WriteString("w-")
	b.WriteString(ShortenID(12))
	b.WriteString("_")
	b.WriteString(CurrentTimeStamp()[:6])
	b.WriteString(RandomNumericString(9))
	return b.String()
}

func GenerateRandomUUIDWithDateTime() string {
	var b bytes.Buffer
	b.WriteString("w-")
	b.WriteString(ShortenID(12))
	b.WriteString("_")
	b.WriteString(CurrentTimeStamp())
	return b.String()
}
