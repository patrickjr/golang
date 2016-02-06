package ff_crypto

import (
	"patrickjr/fantasy_football/lib/ff_crypto"
	"testing"
)

func TestDigestPassword(t *testing.T) {
	pw, err := ff_crypto.DigestPassword("password")
	if err != nil {
		t.Error(err)
	}
	if len(pw) != 60 {
		t.Error(len(pw))
	}
}
