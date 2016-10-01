package main

import (
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
)

func main() {
	for _, arg1 := range os.Args[1:] {
		xlsx, err := excelize.OpenFile(arg1)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		cell := xlsx.GetCellValue("Sheet1", "A1")
		fmt.Println(cell)
		for key,val := range xlsx.XLSX {
			fmt.Printf("%s:%s\n",key,val)
			fmt.Println("------------------------------------------------------------------------")
		}
	}
}
