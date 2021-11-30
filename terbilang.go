package terbilang

import (
	"math"
	"strings"
)

type terbilang struct {
}

func Init() *terbilang {
	return &terbilang{}
}

func (terbilang *terbilang) numberToText(inputValue int64) string {
	value := math.Abs(float64(inputValue))
	huruf := []string{
		"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas",
	}
	temp := ""
	if value < 12 {
		temp = "" + huruf[int64(value)]
	} else if value < 20 {
		temp = terbilang.numberToText(int64(value)-10) + " belas"
	} else if value < 100 {
		temp = terbilang.numberToText(int64(value)/10) + " puluh " + terbilang.numberToText(int64(value)%10)
	} else if value < 200 {
		temp = " seratus " + terbilang.numberToText(int64(value)-100)
	} else if value < 1000 {
		temp = terbilang.numberToText(int64(value)/100) + " ratus " + terbilang.numberToText(int64(value)%100)
	} else if value < 2000 {
		temp = " seribu " + terbilang.numberToText(int64(value)-1000)
	} else if value < 1000000 {
		temp = terbilang.numberToText(int64(value)/1000) + " ribu " + terbilang.numberToText(int64(value)%1000)
	} else if value < 1000000000 {
		temp = terbilang.numberToText(int64(value)/1000000) + " juta " + terbilang.numberToText(int64(value)%1000000)
	} else if value < 1000000000000 {
		temp = terbilang.numberToText(int64(value)/1000000000) + " milyar " + terbilang.numberToText(int64(math.Mod(value, 1000000000)))
	} else if value < 1000000000000000 {
		temp = terbilang.numberToText(int64(value)/1000000000000) + " trilyun " + terbilang.numberToText(int64(math.Mod(value, 1000000000000)))
	}
	//return hasil
	return temp
}

func (terbilang *terbilang) Convert(inputValue int64) string {
	hasil := ""
	if inputValue < 0 {
		hasil = "minus " + strings.TrimSpace(terbilang.numberToText(inputValue))
	} else {
		hasil = strings.TrimSpace(terbilang.numberToText(inputValue))
	}

	return hasil + " rupiah"
}
