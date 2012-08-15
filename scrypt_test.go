package scrypt

import (
	"fmt"
)

func ExampleEncrypt() {
	e := Encrypt("123456789", "salt")
	d := Decrypt(e, "salt")
	fmt.Println(d)
	// Output: 123456789
}
