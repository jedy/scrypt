package scrypt

import "testing"

func TestCrypt(t *testing.T) {
	s := "abcdefgh"
	d := Encrypt(s, "key")
	if dd := Decrypt(d, "key"); dd != s {
		t.Errorf("Source(%v) != Decrypt(Encrypt)(%v)", s, dd)
	}
}
