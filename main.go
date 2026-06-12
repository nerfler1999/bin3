package main

import (
	"bin_3/bins"
	"bin_3/storage"
	"fmt"
)

func main() {
	fmt.Println("**CLI для работы с JSon-файлами**")
	newBinList := storage.NewBinList()
	for {
		a, err := bins.AddNewBin()
		if err != nil {
			fmt.Println(err)
		} else {
			newBinList.AddBinToBinList(a)
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
