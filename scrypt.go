// http://my.oschina.net/Jacker/blog/32837

// Two way encipherment ported from PHP.
package scrypt

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func keyEd(txt []byte, encryptKey string) []byte {
	m := md5.New()
	fmt.Fprint(m, encryptKey)
	t := bytes.NewBufferString("")
	keyBytes := m.Sum(nil)
	for i, ctr := 0, 0; i < len(txt); i++ {
		if ctr == len(keyBytes) {
			ctr = 0
		}
		t.WriteByte(txt[i] ^ keyBytes[ctr])
		ctr++
	}
	return t.Bytes()
}

func Encrypt(txt, key string) string {
	rand.Seed(time.Nanosecond.Nanoseconds())
	m := md5.New()
	fmt.Fprint(m, rand.Int63())
	encryptKey := m.Sum(nil)
	txtBytes := []byte(txt)
	t := bytes.NewBufferString("")
	for i, ctr := 0, 0; i < len(txtBytes); i++ {
		if ctr == len(encryptKey) {
			ctr = 0
		}
		t.WriteByte(encryptKey[ctr])
		t.WriteByte(txtBytes[i] ^ encryptKey[ctr])
		ctr++
	}
	return string(keyEd(t.Bytes(), key))
}

func Decrypt(txt, key string) string {
	b := keyEd([]byte(txt), key)
	t := bytes.NewBufferString("")
	for i := 0; i < len(b); i += 2 {
		t.WriteByte(b[i] ^ b[i+1])
	}
	return t.String()
}

func main() {
	s := "123456789fefefef3r33你好"
	e := Encrypt(s, "salt")
	d := Decrypt(e, "salt")
	fmt.Println(d)
}
