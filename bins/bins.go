package bins

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

func AddNewBin(bin []Bin) ([]Bin, error) {
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
