package main

import (
	"fmt"

	"github.com/didikprabowo/cryptography/cvc"
)

func main() {
	c, _ := cvc.New("didik prabowo ganteng")
	en, _ := c.Encode("DIDIK PRABOWO")
	de, _ := c.Decode(en)
	fmt.Println(en, de)
}
