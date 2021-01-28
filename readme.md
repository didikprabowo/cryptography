# Tools cryptography[On develop]



## How to install
```bash
go get -u github.com/didikprabowo/cryptography
```

## Algorithm
- Caesar Cipher(`cooming soon`)
- [Cryptography Vigenere Cipher](#Cryptography-Vigenere)


### Cryptography Vigenere
Cryptography Vigenere is one of many alogirihm cryptography, this algorithm in a way do shift with key. this key be able to suitable you want.

`example : ` 
```go
package main

import (
	"fmt"

	"github.com/didikprabowo/cryptography/cvc"
)

func main() {
	c := cvc.New("RAHASIABKSDUSDSDSKDSISMSIISS")
	en := c.Encode("My name is Didik Prabowo")
	de := c.Decode(en)
	fmt.Println(en, de)
}

```



