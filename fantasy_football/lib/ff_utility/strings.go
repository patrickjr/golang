package ff_utility

import (
	"math/rand"
	"patrickjr/fantasy_football/lib/ff_random"
	"time"
)

func BytesToString(b []byte) string {
	s := string(b[:])
	return s
}

func GenerateCookie() string {
	var src = rand.NewSource(time.Now().UnixNano())
	return ff_random.RandStringBytesMaskImprSrc(49, src)
}

func ConvertToString(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	} else {
		return ""
	}
}
