package bins

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func AddNewBin(bin []Bin) ([]Bin, error) {
	var array Bin
	var isPrivate string
	fmt.Println("Введите ID: ")
	fmt.Scanln(&array.Id)
	if array.Id == "" {
		return nil, errors.New("Ошибка! Пустая строка")
	}
	fmt.Println("Сделать приватным? (y/n)")
	fmt.Scanln(&isPrivate)
	if isPrivate == "" {
		return nil, errors.New("Ошибка! Пустая строка")
	}
	if isPrivate == "y" {
		array.Private = true
	}
	if isPrivate == "n" {
		array.Private = false
	}
	fmt.Println("Введите имя: ")
	fmt.Scanln(&array.Name)
	if array.Name == "" {
		return nil, errors.New("Ошибка! Пустая строка")
	}
	array.CreatedAt = time.Now()
	bin = append(bin, array)
	return bin, nil
}
