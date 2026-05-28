package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func main() {
	fmt.Println("**CLI для работы с Json-файлами**")
	for {
		BinList := []Bin{}
		a, err := addNewBin(BinList)
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

func addNewBin(bin []Bin) ([]Bin, error) {
	var array Bin
	var isPrivate string
	fmt.Println("Введите ID: ")
	fmt.Scanln(&array.id)
	if array.id == "" {
		return nil, errors.New("Ошибка! Пустая строка")
	}
	fmt.Println("Сделать приватным? (y/n)")
	fmt.Scanln(&isPrivate)
	if isPrivate == "" {
		return nil, errors.New("Ошибка! Пустая строка")
	}
	if isPrivate == "y" {
		array.private = true
	}
	if isPrivate == "n" {
		array.private = false
	}
	fmt.Println("Введите имя: ")
	fmt.Scanln(&array.name)
	if array.name == "" {
		return nil, errors.New("Ошибка! Пустая строка")
	}
	array.createdAt = time.Now()
	bin = append(bin, array)
	return bin, nil
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
