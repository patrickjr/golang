package ff_random

import (
	"math/rand"
	"patrickjr/fantasy_football/lib/ff_random"
	"testing"
	"time"
)

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	random_string := ff_random.RandStringBytesMaskImprSrc(49, rand.NewSource(time.Now().UnixNano()))
	if len(random_string) != 49 {
		t.Error(random_string)
	}
}
