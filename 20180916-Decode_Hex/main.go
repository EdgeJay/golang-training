package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("0026")
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hexadecimal %s is %s\n", string(src), dst[:n])
}
