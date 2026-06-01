package main

import (
	"bin_3/bins"
	"fmt"
)

func main() {
	fmt.Println("**CLI для работы с JSon-файлами**")
	var BinList []bins.Bin
	for {
		a, err := bins.AddNewBin(BinList)
		if a != nil {
			fmt.Println(a)
		} else {
			fmt.Println(err)
		}
		res := wantToRestart()
		if res != true {
			break
		}
	}
}

func wantToRestart() bool {
	var res string
	fmt.Println("Желаете добавить еще один файл? (y/n)")
	fmt.Scan(&res)
	if res == "y" {
		return true
	} else {
		return false
	}
}
