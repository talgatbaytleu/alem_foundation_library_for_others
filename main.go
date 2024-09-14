package main

import (
	"fmt"
	"os"

	"lib_for_others/csv"
)

func main() {
	file, _ := os.Open("./csv_lf.csv")
	defer file.Close()
	a := csv.CSVstruct{}
	fmt.Println(a.ReadLine(file))
	fmt.Println(a.ReadLine(file))
	fmt.Println(a.ReadLine(file))
}
