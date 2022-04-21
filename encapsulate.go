package main

import (
	"fmt"

	"github.com/mungkiice/kelas-online/objek"
)

func main() {
	dinda := objek.Constructor("dinda", 10)
	fmt.Println(dinda.Name)
	fmt.Println(dinda.KasihTauSaldo())
}
