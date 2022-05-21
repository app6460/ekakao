package emotpreview

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

func shortenID(t int) string {
	var b bytes.Buffer
	for i := 0; i < t; i++ {
		n := rand.Int() % len(seedKey)
		b.WriteString(seedKey[n])
	}
	return b.String()
}

func randomNumericString(t int) string {
	var b bytes.Buffer
	for i := 0; i < t; i++ {
		n := rand.Int()%10 + 48
		b.WriteString(string(n))
	}
	return b.String()
}

func currentTimeStamp() string {
	t := time.Now().Add(time.Hour * 9)
	s := t.Format("2006-01-02T15:04:05.999")
	return reg.ReplaceAllString(s, "")[2:]
}

func generateRandomUUIDWithDateNumber() string {
	var b bytes.Buffer
	b.WriteString("w-")
	b.WriteString(shortenID(12))
	b.WriteString("_")
	b.WriteString(currentTimeStamp()[:6])
	b.WriteString(randomNumericString(9))
	return b.String()
}

func generateRandomUUIDWithDateTime() string {
	var b bytes.Buffer
	b.WriteString("w-")
	b.WriteString(shortenID(12))
	b.WriteString("_")
	b.WriteString(currentTimeStamp())
	return b.String()
}
