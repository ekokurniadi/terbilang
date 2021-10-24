# 👋 Welcome

# TERBILANG
Terbilang adalah package untuk mengubah nominal angka rupiah ke dalam nominal angka rupiah dalam bentuk teks 

### How to install
```sh
 go get github.com/ekokurniadi/terbilang
```

### Then import the indodate package

```go
import "github.com/ekokurniadi/terbilang"
```
### Example Project

```go
package main

import (
	"fmt"
	"time"

	"github.com/ekokurniadi/terbilang"
)

func main() {
	angka := 50000
    terbilang := Terbilang{}
    fmt.Println(terbilang.Generate(angka))
    
    //output => lima puluh ribu rupiah
}

```
### Run example project and you will see the output
```sh
 go run main.go
lima puluh ribu rupiah
```

### Have Fun :D
