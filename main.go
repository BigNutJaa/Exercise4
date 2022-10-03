package main

import (
	"github.com/BigNutJaa/sortnumASC"
	"log"
)

func main() {

	result, err := sortnumASC.SortB("52716")
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(result)
}
