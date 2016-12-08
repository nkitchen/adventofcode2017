package main

import "bytes"
import "crypto/md5"
import "flag"
import "fmt"
import "log"

func main() {
	flag.Parse()
	input := flag.Arg(0)
	password := []byte("________")
	remaining := 8
	i := 0
	b := make([]byte, 0, 1024)
	for remaining > 0 {
		buf := bytes.NewBuffer(b)
		n, err := fmt.Fprintf(buf, "%s%d", input, i)
		if err != nil {
			log.Fatal(err)
		}
		s := md5.Sum(b[:n])
		h := fmt.Sprintf("%x", s[:])
		if h[:5] == "00000" {
			k := int(h[5] - '0')
			if k < len(password) && password[k] == '_' {
				password[k] = h[6]
				remaining--
				fmt.Printf("%s\n", password)
			}
		}
		i++
	}
}
