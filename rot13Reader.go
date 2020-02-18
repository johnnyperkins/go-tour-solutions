package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	cypher := make(map[byte]byte)
	decrypted := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	encrypted := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

	for i := range decrypted {
		cypher[decrypted[i]] = encrypted[i]
	}

	n, err := r13.r.Read(b)

	for i := 0; i < n; i++ {
		if val, ok := cypher[b[i]]; ok {
			b[i] = val
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
