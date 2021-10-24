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

### Example response when i using terbilang package


```json 
{
    "meta": {
        "message": "Campaign Detail",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "name": "Campaign 1",
        "short_description": "short",
        "description": "description",
        "image_url": "satu.jpg",
        "goal_amount": 100000,
        "amount_on_text": "seratus ribu rupiah",
        "current_amount": 0,
        "backer_count": 0,
        "user_id": 12,
        "slug": "campaign-1",
        "user": {
            "name": "Eko Kurniadi",
            "image_url": "images/1-laporan.png"
        },
        "perks": [
            "satu",
            "dua",
            "tiga"
        ],
        "images": [
            {
                "image_url": "satu.jpg",
                "is_primary": false
            },
            {
                "image_url": "images/1-transaksi.PNG",
                "is_primary": false
            },
            {
                "image_url": "images/1-dashboard.PNG",
                "is_primary": false
            },
            {
                "image_url": "images/1-dashboard.PNG",
                "is_primary": true
            },
            {
                "image_url": "images/1-dashboard.PNG",
                "is_primary": false
            }
        ]
    }
}
```
