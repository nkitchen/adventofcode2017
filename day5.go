package main

import "bytes"
import "crypto/md5"
import "flag"
import "fmt"
import "log"

func main() {
	flag.Parse()
	input := flag.Arg(0)
	password := []byte{}
	i := 0
	b := make([]byte, 0, 1024)
	for len(password) < 8 {
		buf := bytes.NewBuffer(b)
		n, err := fmt.Fprintf(buf, "%s%d", input, i)
		if err != nil {
			log.Fatal(err)
		}
		s := md5.Sum(b[:n])
		h := fmt.Sprintf("%x", s[:])
		if h[:5] == "00000" {
			password = append(password, h[5])
		}
		i++
	}
	fmt.Printf("%s\n", password)
}
